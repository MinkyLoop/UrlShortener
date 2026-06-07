package repository

import (
	"context"
	"urlshortener/internal/domain"

	"github.com/jackc/pgx/v5"
)

type URLRepo struct {
	conn *pgx.Conn
}

func NewURLRepo(dburl string) (*URLRepo, error) {
	conn, err := pgx.Connect(context.Background(), dburl)
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	return &URLRepo{conn: conn}, nil
}

func (r *URLRepo) Set(ctx context.Context, url *domain.Url) error {
	_, err := r.conn.Exec(ctx, `
		INSERT INTO urlshortener.urls (short_code, url)
		VALUES ($1, $2)`,
		url.ShortCode, url.OriginalUrl)

	return err
}

func (r *URLRepo) Get(ctx context.Context, shortCode string) (*domain.Url, error) {
	var url domain.Url

	err := r.conn.QueryRow(ctx, `
		SELECT short_code, url
		FROM urlshortener.urls
		WHERE short_code = $1;`,
		shortCode).Scan(&url.ShortCode, &url.OriginalUrl)

	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &url, nil
}

func (r *URLRepo) Close() error {
	return r.conn.Close(context.Background())
}

//:TODO func (r *URLRepo) Update() {} Можно добавить время жизни ссылки
//:TODO func (r *URLRepo) Delete() {} Можно добавить время жизни ссылки
