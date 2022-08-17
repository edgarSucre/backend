package httpserver

import "time"

type HttpServerOption func(*Server)

func WithPort(port string) HttpServerOption {
	return func(s *Server) {
		s.instance.Addr = port
	}
}

func WithReadTimeOut(t time.Duration) HttpServerOption {
	return func(s *Server) {
		s.instance.ReadTimeout = t
	}
}

func WithReadHeaderTimeOout(t time.Duration) HttpServerOption {
	return func(s *Server) {
		s.instance.ReadHeaderTimeout = t
	}
}

func WithWriteTimeOut(t time.Duration) HttpServerOption {
	return func(s *Server) {
		s.instance.WriteTimeout = t
	}
}

func WithExitTimeOut(t time.Duration) HttpServerOption {
	return func(s *Server) {
		s.exitTimeOut = t
	}
}
