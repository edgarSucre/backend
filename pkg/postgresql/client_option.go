package postgresql

import "time"

type ClientOption func(*Client)

func WithMaxPoolSize(pz int) ClientOption {
	return func(p *Client) {
		p.maxPoolConn = pz
	}
}

func WithConnAttempts(att int) ClientOption {
	return func(p *Client) {
		p.connAttempts = att
	}
}

func WithConnTimeout(t time.Duration) ClientOption {
	return func(p *Client) {
		p.connTimeout = t
	}
}
