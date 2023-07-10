package postgresql

import (
	"database/sql"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	DEFAULT_MAX_POOL_CONN          = 1
	DEFAULT_POOL_CREATION_ATTEMPTS = 10
	DEFAULT_POOL_CREATION_TIMEOUT  = time.Millisecond * 200
)

type Client struct {
	maxPoolConn  int
	connAttempts int
	connTimeout  time.Duration
	pool         *pgxpool.Pool
	db           *sql.DB
}

func New(opts ...ClientOption) (*Client, error) {
	c := &Client{
		maxPoolConn:  DEFAULT_MAX_POOL_CONN,
		connAttempts: DEFAULT_POOL_CREATION_ATTEMPTS,
		connTimeout:  DEFAULT_POOL_CREATION_TIMEOUT,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c, nil
}

func (c *Client) Close() {
	if c.pool != nil {
		c.pool.Close()
	}
}
