package postgresql

import "time"

type ClientOption func(*client)

func WithMaxPoolSize(pz int) ClientOption {
	return func(p *client) {
		p.maxPoolSize = pz
	}
}

func WithConnAttempts(att int) ClientOption {
	return func(p *client) {
		p.connAttempts = att
	}
}

func WithConnTimeout(t time.Duration) ClientOption {
	return func(p *client) {
		p.connTimeout = t
	}
}
