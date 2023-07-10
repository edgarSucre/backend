package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (c *Client) GetInstanceSqlDB(ctx context.Context, opts UrlOpts) (*sql.DB, error) {
	if c.db != nil && c.db.Ping() != nil {
		return c.db, nil
	}

	url := opts.URL()
	config, err := c.getConnConfig(url)
	if err != nil {
		return nil, fmt.Errorf("c.getConfig(url), %w", err)
	}

	if err := c.attemptDbConnection(ctx, c.openSqlDB(config)); err != nil {
		return nil, fmt.Errorf("[postgresql] can't get sql connection {sql.ping: %w}", err)
	}

	return c.db, nil
}

func (c *Client) getConnConfig(url string) (*pgx.ConnConfig, error) {
	config, err := pgx.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("[postgresql] invalid conn config {pgx.ParseConfig(url): %w}", err)
	}

	config.ConnectTimeout = c.connTimeout

	return config, nil
}

func (c *Client) GetInstancePool(ctx context.Context, urlOpts UrlOpts) (*pgxpool.Pool, error) {
	if c.pool != nil && c.pool.Ping(ctx) != nil {
		return c.pool, nil
	}

	url := urlOpts.URL()
	config, err := c.getPoolConfig(url)
	if err != nil {
		return nil, fmt.Errorf("c.getPoolConfig(url), %w", err)
	}

	if err := c.attemptDbConnection(ctx, c.createPool(config)); err != nil {
		return nil, fmt.Errorf("[postgres] can't get connection from pool {pool.Ping: %w}", err)
	}

	return c.pool, nil
}

func (c *Client) getPoolConfig(url string) (*pgxpool.Config, error) {
	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("[postgresql] invalid pool config {pgxpool.ParseConfig: %w}", err)
	}

	poolConfig.MaxConns = int32(c.maxPoolConn)
	poolConfig.ConnConfig.ConnectTimeout = c.connTimeout

	return poolConfig, nil
}

func (c *Client) attemptDbConnection(ctx context.Context, connFn connectFn) error {
	attempts := c.connAttempts
	var err error

	for attempts > 0 {
		if err := connFn(ctx); err == nil {
			return nil
		}

		time.Sleep(time.Millisecond * 300)
		attempts--

		log.Printf("trying to connect to postgres, attempts left: %d", attempts)
	}

	return fmt.Errorf("connFn, %w", err)
}
