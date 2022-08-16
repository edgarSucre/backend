package httpserver

import "time"

type Option func(*Server)

func WithPort(port string) Option {
	return func(s *Server) {
		s.instance.Addr = port
	}
}

func WithReadTimeOut(t time.Duration) Option {
	return func(s *Server) {
		s.instance.ReadTimeout = t
	}
}

func WithReadHeaderTimeOout(t time.Duration) Option {
	return func(s *Server) {
		s.instance.ReadHeaderTimeout = t
	}
}

func WithWriteTimeOut(t time.Duration) Option {
	return func(s *Server) {
		s.instance.WriteTimeout = t
	}
}

func WithExitTimeOut(t time.Duration) Option {
	return func(s *Server) {
		s.exitTimeOut = t
	}
}
