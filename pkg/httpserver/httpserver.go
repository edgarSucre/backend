package httpserver

import (
	"context"
	"net/http"
	"time"
)

//TODO: Add logger
type Server struct {
	instance    *http.Server
	onErr       chan error
	exitTimeOut time.Duration
}

const (
	DEFAULT_WRITE_TIMEOUT       = 4 * time.Second
	DEFAULT_READ_TIMEOUT        = 4 * time.Second
	DEFAULT_HEADER_READ_TIMEOUT = 2 * time.Second
	DEFAULT_ADDRES              = ":8080"
	DEFAULT_EXIT_TIMEOUT        = 3 * time.Second
)

func New(handler http.Handler, opts ...HttpServerOption) *Server {
	s := &http.Server{
		Handler:           handler,
		Addr:              DEFAULT_ADDRES,
		ReadTimeout:       DEFAULT_READ_TIMEOUT,
		ReadHeaderTimeout: DEFAULT_HEADER_READ_TIMEOUT,
		WriteTimeout:      DEFAULT_WRITE_TIMEOUT,
	}

	server := &Server{
		instance:    s,
		onErr:       make(chan error),
		exitTimeOut: DEFAULT_EXIT_TIMEOUT,
	}

	for _, option := range opts {
		option(server)
	}

	return server
}

func (s *Server) Start() error {
	return s.instance.ListenAndServe()
}

func (s *Server) Exit() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.exitTimeOut)
	defer cancel()

	return s.instance.Shutdown(ctx)
}
