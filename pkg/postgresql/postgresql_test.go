package postgresql

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

func TestClient(t *testing.T) {

	t.Run("Invalid url", func(t *testing.T) {
		url := ""
		c, err := New(url, WithConnAttempts(1), WithConnTimeout(time.Millisecond*100))

		if c != nil {
			t.Errorf("Failed, Expected Client to be nil, Got: %v", c)
		}

		if !strings.Contains(err.Error(), "dial error") {
			t.Errorf("Failed, Expected connection error, Got: %v", err)
		}
	})

	//to test the actual conn a valid connection string is required
	useDB := os.Getenv("USE_DB")
	if strings.ToLower(useDB) == "true" {

		url := fmt.Sprintf(
			"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)

		connAttempts := 5
		connTimeOut := 500 * time.Millisecond
		poolSize := 2

		c, err := New(
			url,
			WithConnAttempts(connAttempts),
			WithConnTimeout(connTimeOut),
			WithMaxPoolSize(poolSize),
		)

		// checks if the connection failed
		if err != nil {
			t.Errorf("Failed, Expected Err to be nil, Got: %v", err)
		}

		cases := []struct {
			expected any
			got      any
			n        string
		}{
			{connAttempts, c.connAttempts, "WithConnAttempts"},
			{connTimeOut, c.connTimeout, "WithConnTimeout"},
			{poolSize, c.maxPoolSize, "WithMaxPoolSize"},
		}

		// checks if the parameters were set correctly
		for _, tc := range cases {
			t.Run(tc.n, func(t *testing.T) {
				if tc.expected != tc.got {
					t.Errorf("Failed, Expected: %v, Got: %v", tc.expected, tc.got)
				}
			})
		}

		// finally test if the db responds
		err = c.Pool.Ping(context.Background())
		if err != nil {
			t.Errorf("Failed, Expected Err to be nil, Got: %v", err)
		}

	}
}
