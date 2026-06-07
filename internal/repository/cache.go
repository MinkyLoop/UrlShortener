package repository

import (
	"context"

	"golang.org/x/sync/singleflight"
)

type CacheStrategy struct {
	redis *RedisRepo
	pg    *URLRepo
	sf    singleflight.Group
}

func NewCacheStrategy(redis *RedisRepo, pg *URLRepo) *CacheStrategy {
	return &CacheStrategy{
		redis: redis,
		pg:    pg,
	}
}

func (c *CacheStrategy) GetUrl(ctx context.Context, shortCode string) (string, error) {
	OriginalUrl, err := c.redis.Get(ctx, shortCode)
	if err == nil {
		return OriginalUrl, nil
	}

	v, err, _ := c.sf.Do(shortCode, func() (interface{}, error) {
		Url, err := c.pg.Get(ctx, shortCode)
		if err != nil {
			return nil, err
		}

		if Url == nil {
			return nil, nil
		}

		c.redis.Set(ctx, shortCode, Url.OriginalUrl)

		return Url.OriginalUrl, nil
	})
	if err != nil {
		return "", err
	}

	return v.(string), err
}
