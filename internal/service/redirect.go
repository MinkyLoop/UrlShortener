package service

import (
	"context"
	"urlshortener/internal/repository"
)

type Redirect struct {
	cache *repository.CacheStrategy
}

func NewRedirect(cache *repository.CacheStrategy) *Redirect {
	return &Redirect{cache: cache}
}

func (r *Redirect) Redirect(ctx context.Context, shortCode string) (string, error) {
	url, err := r.cache.GetUrl(ctx, shortCode)
	if err != nil {
		return "", err
	}
	
	return url, nil
}
