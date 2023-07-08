package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func (c *client) GetInstanceSqlDB(opts UrlOpts) (*sql.DB, error) {
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

func (c *client) GetInstancePool(ctx context.Context, url string) (*pgxpool.Pool, error) {
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
		c.pool, err = pgxpool.ConnectConfig(ctx, poolConfig)
		if err == nil {
			break
		}

		//TODO: replace this with logrus
		log.Printf("Postgres is trying to connect, attempts left: %d", attempts)

		time.Sleep(c.connTimeout)

		attempts--
	}

	if err != nil {
		return nil, fmt.Errorf("postgres - can't get Pool instance - pgxpool.ConnectConfig: %w", err)
	}

	return c.pool, nil
}
