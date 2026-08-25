package postgres

import (
	"context"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joaobotoni/knock"
)

func dsn(d *knock.Database) string {
	return (&url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(d.Username, d.Password),
		Host:   fmt.Sprintf("%s:%d", d.Host, d.Port),
		Path:   d.Name,
	}).String()
}

func Open(ctx *context.Context, d *knock.Database) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(*ctx, dsn(d))
	if err != nil {
		return nil, fmt.Errorf("erro ao criar pool: %w", err)
	}

	if err := pool.Ping(*ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("erro ao conectar: %w", err)
	}
	return pool, nil
}