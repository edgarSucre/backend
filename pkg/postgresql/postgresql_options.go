package postgresql

import "time"

type PostgresOption func(*Client)

func WithMaxPoolSize(pz int) PostgresOption {
	return func(p *Client) {
		p.maxPoolSize = pz
	}
}

func WithConnAttempts(att int) PostgresOption {
	return func(p *Client) {
		p.connAttempts = att
	}
}

func WithConnTimeout(t time.Duration) PostgresOption {
	return func(p *Client) {
		p.connTimeout = t
	}
}
