package postgresql_test

import (
	"context"
	"path/filepath"
	"strings"
	"testing"

	"github.com/edgarSucre/backend/pkg/postgresql"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func TestGetInstancePool(t *testing.T) {
	p := filepath.Join("..", "..", ".env")
	env, err := godotenv.Read(p)
	if err != nil {
		t.Fatal("error: can't load env file")
	}

	workingUser := env["DB_USER"]
	workingPass := env["DB_PASS"]

	cases := []struct {
		name    string
		urlOpts postgresql.UrlOpts
		check   func(t *testing.T, pool *pgxpool.Pool, err error)
	}{
		{
			"success",
			postgresql.UrlOpts{
				User: workingUser,
				Pass: workingPass,
			},
			func(t *testing.T, pool *pgxpool.Pool, err error) {
				if err != nil {
					t.Errorf("expected err to be nil, got: %s", err)
				}

				if pool == nil {
					t.Error("expected pool to not be nil")
				}

				err = pool.Ping(context.Background())
				if err != nil {
					t.Errorf("expected pool.Ping to be nil, got: %s", err)
				}

				pool.Close()
			},
		},
		{
			"parseError",
			postgresql.UrlOpts{User: "\\@"},
			func(t *testing.T, pool *pgxpool.Pool, err error) {
				if err == nil {
					t.Errorf("expected err to not be nil")
				}

				msg := "invalid pool config"
				if !strings.Contains(err.Error(), msg) {
					t.Errorf("expected err to contain: %s, got: %s", msg, err)
				}

				if pool != nil {
					t.Error("expected pool to be nil")
				}

			},
		},
		{
			"invalidCreds",
			postgresql.UrlOpts{
				User: "vito",
				Pass: workingPass,
			},
			func(t *testing.T, pool *pgxpool.Pool, err error) {
				if err == nil {
					t.Errorf("expected err to not be nil")
				}

				msg := "can't get connection from pool"
				if !strings.Contains(err.Error(), msg) {
					t.Errorf("expected err to contain: %s, got: %s", msg, err)
				}

				if pool != nil {
					t.Error("expected pool to be nil")
				}

			},
		},
	}

	for _, tc := range cases {
		client, err := postgresql.New(postgresql.WithConnAttempts(1))
		if err != nil {
			t.Fatal("error: can't create postgres client")
		}

		t.Run(tc.name, func(t *testing.T) {
			instance, err := client.GetInstancePool(context.Background(), tc.urlOpts)
			tc.check(t, instance, err)
		})
	}
}
