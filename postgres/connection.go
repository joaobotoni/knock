package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/url"
	"strconv"
)

func Open(ctx context.Context, c *Config) (*pgxpool.Pool, error) {
	if err := validate(c); err != nil {
		return nil, fmt.Errorf("configuração inválida: %w", err)
	}

	pool, err := pgxpool.New(ctx, dsn(c))
	if err != nil {
		return nil, fmt.Errorf("erro ao criar pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("erro ao conectar: %w", err)
	}
	return pool, nil
}

func dsn(d *Config) string {
	return (&url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(d.Username, d.Password),
		Host:   d.Host + ":" + strconv.Itoa(d.Port),
		Path:   "/" + d.Name,
	}).String()
}
