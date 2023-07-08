package postgresql

import "fmt"

type UrlOpts struct {
	DBName string
	Host   string
	Pass   string
	Port   int
	User   string
}

// TODO: add validation fro mandatory fields
func (opt UrlOpts) URL() string {
	if opt.DBName == "" {
		opt.DBName = opt.User
	}

	if opt.Port == 0 {
		opt.Port = 5432
	}

	if opt.Host == "" {
		opt.Host = "localhost"
	}

	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%v/%s?sslmode=disable",
		opt.User,
		opt.Pass,
		opt.Host,
		opt.Port,
		opt.DBName,
	)
}
