package postgresql

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	DEFAULT_MAX_POOL_SIZE = 1
	DEFAULT_CONN_ATTEMPTS = 10
	DEFAULT_CONN_TIMEOUT  = time.Second
)

type Client struct {
	maxPoolSize  int
	close        <-chan struct{}
	connAttempts int
	connTimeout  time.Duration
	Pool         *pgxpool.Pool
}

func New(url string, opts ...PostgresOption) (*Client, error) {
	p := &Client{
		maxPoolSize:  DEFAULT_MAX_POOL_SIZE,
		connAttempts: DEFAULT_CONN_ATTEMPTS,
		connTimeout:  DEFAULT_CONN_TIMEOUT,
	}

	for _, opt := range opts {
		opt(p)
	}

	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("postgres - can't start client - pgxpool.ParseConfig: %w", err)
	}

	poolConfig.MaxConns = int32(p.maxPoolSize)

	attempts := p.connAttempts
	for attempts > 0 {
		p.Pool, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
		if err == nil {
			break
		}

		//TODO: replace this with logrus
		log.Printf("Postgres is trying to connect, attempts left: %d", attempts)

		time.Sleep(p.connTimeout)

		attempts--
	}

	if err != nil {
		return nil, fmt.Errorf("postgres - can't start client - connAttempts == 0: %w", err)
	}

	if p.close != nil {
		go func() {
			<-p.close
			p.Close()
		}()
	}

	return p, nil
}

func (c *Client) Close() {
	if c.Pool != nil {
		c.Pool.Close()
	}
}
