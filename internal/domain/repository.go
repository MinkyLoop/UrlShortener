package domain

import "context"

type UrlRepository interface {
	Set(ctx context.Context, url *Url) error
	Get(ctx context.Context, shortCode string) (*Url, error)
}

type CacheRepository interface {
	Set(ctx context.Context, shortCode string, originalUrl string) error
	Get(ctx context.Context, shortCode string) (string, error)
}
