package service

import (
	"context"
	"crypto/rand"
	"math/big"
	"time"

	"urlshortener/internal/domain"
)

type Shortener struct {
	repo    domain.UrlRepository
	baseUrl string
}

func NewShortener(repo domain.UrlRepository, baseUrl string) *Shortener {
	return &Shortener{
		repo:    repo,
		baseUrl: baseUrl,
	}
}

func generateShortCode(length int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	result := make([]byte, length)

	for i := 0; i < length; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(alphabet))))
		result[i] = alphabet[n.Int64()]
	}

	return string(result)
}

func (s *Shortener) Create(ctx context.Context, originalUrl string) (*domain.Url, error) {
	var shortCode string

	for {
		shortCode = generateShortCode(6)
		if v, _ := s.repo.Get(ctx, shortCode); v == nil {
			break
		}

		time.Sleep(10 * time.Millisecond)
	}

	url := &domain.Url{
		OriginalUrl: originalUrl,
		ShortCode:   shortCode,
	}

	if err := s.repo.Set(ctx, url); err != nil {
		return nil, err
	}

	return url, nil
}
