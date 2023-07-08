package postgresql

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	DEFAULT_MAX_POOL_SIZE          = 1
	DEFAULT_POOL_CREATION_ATTEMPTS = 10
	DEFAULT_POOL_CREATION_TIMEOUT  = time.Millisecond * 200
)

type client struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration
	pool         *pgxpool.Pool
}

func New(opts ...ClientOption) (*client, error) {
	c := &client{
		maxPoolSize:  DEFAULT_MAX_POOL_SIZE,
		connAttempts: DEFAULT_POOL_CREATION_ATTEMPTS,
		connTimeout:  DEFAULT_POOL_CREATION_TIMEOUT,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c, nil
}

func (c *client) Close() {
	if c.pool != nil {
		c.pool.Close()
	}
}
