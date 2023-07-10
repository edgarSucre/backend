package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

type connectFn func(context.Context) error

func (c *Client) openSqlDB(config *pgx.ConnConfig) connectFn {
	return func(ctx context.Context) error {
		db := stdlib.OpenDB(*config)
		if err := db.PingContext(ctx); err != nil {
			return fmt.Errorf("[postgresql] can't connect to db {sql.ping: %w}", err)
		}

		c.db = db

		return nil
	}
}

func (c *Client) createPool(config *pgxpool.Config) connectFn {
	return func(ctx context.Context) error {
		pool, err := pgxpool.NewWithConfig(ctx, config)
		if err != nil {
			return fmt.Errorf("[postgresql] can't create pool {pgxpool.NewWithConfig: %w}", err)
		}

		if err := pool.Ping(ctx); err != nil {
			return fmt.Errorf("[postgresql] can't get connectionfrom pool {pool.ping: %w}", err)
		}

		c.pool = pool

		return nil
	}
}
