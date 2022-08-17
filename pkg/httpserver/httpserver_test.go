package httpserver

import (
	"errors"
	"net/http"
	"testing"
	"time"
)

type handler struct{}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte{})
}

func TestServer(t *testing.T) {
	readTO := 100 * time.Millisecond
	readHTO := 200 * time.Millisecond
	writeTO := 300 * time.Millisecond
	exitTO := 400 * time.Millisecond
	port := ":9090"

	h := handler{}
	server := New(
		h,
		WithPort(port),
		WithReadHeaderTimeOout(readHTO),
		WithReadTimeOut(readTO),
		WithWriteTimeOut(writeTO),
		WithExitTimeOut(exitTO),
	)

	cases := []struct {
		expected any
		got      any
		n        string
	}{
		{port, server.instance.Addr, "WithPort"},
		{readTO, server.instance.ReadTimeout, "WithReadHeaderTimeOout"},
		{readHTO, server.instance.ReadHeaderTimeout, "WithReadTimeOut"},
		{writeTO, server.instance.WriteTimeout, "WithWriteTimeOut"},
		{exitTO, server.exitTimeOut, "WithExitTimeOut"},
	}

	for _, tc := range cases {
		t.Run(tc.n, func(t *testing.T) {
			if tc.expected != tc.got {
				t.Errorf("Failed, Expected: %v, Got: %v", tc.expected, tc.got)
			}
		})
	}
}

func TestRun(t *testing.T) {
	h := handler{}
	server := New(h)
	go func() {
		time.Sleep(300 * time.Millisecond)
		server.Exit()
	}()

	err := server.Start()
	if !errors.Is(err, http.ErrServerClosed) {
		t.Errorf("Failed, Expected: %v, Got: %v", http.ErrServerClosed, err)
	}
}
