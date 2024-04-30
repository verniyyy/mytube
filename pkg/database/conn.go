package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// NewConn creates a new connection pool to the database from given configuration.
func NewConn(ctx context.Context, cfg *Config) (*pgxpool.Pool, error) {
	pgxCfg, err := pgxpool.ParseConfig(cfg.ConnectionURL())
	if err != nil {
		return nil, fmt.Errorf("failed to parse config : %w", err)
	}

	pgxCfg.BeforeAcquire = func(ctx context.Context, c *pgx.Conn) bool {
		return c.Ping(ctx) == nil
	}

	pgxPool, err := pgxpool.NewWithConfig(ctx, pgxCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool : %w", err)
	}

	return pgxPool, nil
}
