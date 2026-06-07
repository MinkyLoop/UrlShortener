package repository

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	client *redis.Client
	ttl    time.Duration
}

func NewRedisRepo(redisUrl string) (*RedisRepo, error) {
	client := redis.NewClient(&redis.Options{Addr: redisUrl})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return &RedisRepo{
		client: client,
		ttl:    24 * time.Hour,
	}, nil
}

func (r *RedisRepo) Set(ctx context.Context, shorCode string, originalUrl string) error {
	return r.client.Set(ctx, shorCode, originalUrl, r.ttl).Err()
}

func (r *RedisRepo) Get(ctx context.Context, shortCode string) (string, error) {
	return r.client.Get(ctx, shortCode).Result()
}

func (r *RedisRepo) Close() error {
	return r.client.Close()
}
