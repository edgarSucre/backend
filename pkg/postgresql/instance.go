package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func (c *Client) GetInstanceSqlDB(opts UrlOpts) (*sql.DB, error) {
	url := opts.URL()

	db, err := sql.Open("pgx", url)
	if err != nil {
		return nil, fmt.Errorf("postgresql - can't get sql instance - sql.open: %w", err)
	}

	err = db.Ping()

	if err != nil {
		return nil, fmt.Errorf("postgresql - can't get sql connection - sql.ping: %w", err)
	}

	return db, nil
}

func (c *Client) GetInstancePool(ctx context.Context, urlOpts UrlOpts) (*pgxpool.Pool, error) {
	url := urlOpts.URL()

	// if pool is active return it
	if c.pool != nil && c.pool.Ping(ctx) != nil {
		return c.pool, nil
	}

	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("postgresql - invalid pool configuration - pgxpool.ParseConfig: %w", err)
	}

	poolConfig.MaxConns = int32(c.maxPoolSize)

	attempts := c.connAttempts
	for attempts > 0 {
		c.pool, err = pgxpool.NewWithConfig(ctx, poolConfig)
		if c.pool != nil && c.pool.Ping(ctx) == nil {
			return c.pool, nil
		}

		//TODO: replace this with logrus
		log.Printf("Postgres is trying to connect, attempts left: %d", attempts)
		time.Sleep(c.connTimeout)

		attempts--
	}

	return nil, fmt.Errorf("postgres - can't get connection from pool - pool.Ping")
}
