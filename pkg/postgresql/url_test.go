package postgresql

import "testing"

func TestUrlOpts(t *testing.T) {
	cases := []struct {
		name     string
		opt      UrlOpts
		expected string
	}{
		{
			"noDBName",
			UrlOpts{
				Host: "host",
				Pass: "pass",
				Port: 36,
				User: "user",
			},
			"postgresql://user:pass@host:36/user?sslmode=disable",
		},
		{
			"noPort",
			UrlOpts{
				DBName: "db",
				Host:   "host",
				Pass:   "pass",
				User:   "user",
			},
			"postgresql://user:pass@host:5432/db?sslmode=disable",
		},
		{
			"noHost",
			UrlOpts{
				DBName: "db",
				Pass:   "pass",
				Port:   36,
				User:   "user",
			},
			"postgresql://user:pass@localhost:36/db?sslmode=disable",
		},
		{
			"withAll",
			UrlOpts{
				DBName: "db",
				Host:   "host",
				Pass:   "pass",
				Port:   36,
				User:   "user",
			},
			"postgresql://user:pass@host:36/db?sslmode=disable",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.opt.URL()
			if tc.expected != actual {
				template := `
					value missmatch: {
					  expected: %s
					  actual: %s
					}
				`
				t.Errorf(template, tc.expected, actual)
			}
		})
	}
}
