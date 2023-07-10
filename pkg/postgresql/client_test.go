package postgresql

import (
	"testing"
	"time"
)

func TestClient(t *testing.T) {

	// t.Run("Invalid url", func(t *testing.T) {
	// 	c, err := New(WithConnAttempts(1), WithConnTimeout(time.Millisecond*100))

	// 	if c != nil {
	// 		t.Errorf("Failed, Expected Client to be nil, Got: %v", c)
	// 	}

	// 	if !strings.Contains(err.Error(), "dial error") {
	// 		t.Errorf("Failed, Expected connection error, Got: %v", err)
	// 	}
	// })

	//to test the actual conn a valid connection string is required
	connAttempts := 5
	connTimeOut := 500 * time.Millisecond
	poolSize := 2

	c, err := New(
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
		{poolSize, c.maxPoolConn, "WithMaxPoolSize"},
	}

	// checks if the parameters were set correctly
	for _, tc := range cases {
		t.Run(tc.n, func(t *testing.T) {
			if tc.expected != tc.got {
				t.Errorf("Failed, Expected: %v, Got: %v", tc.expected, tc.got)
			}
		})
	}

	// // finally test if the db responds
	// err = c.pool.Ping(context.Background())
	// if err != nil {
	// 	t.Errorf("Failed, Expected Err to be nil, Got: %v", err)
	// }
}
