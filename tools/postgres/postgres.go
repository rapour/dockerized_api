package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Client struct {
	Config Config
	DB     *sql.DB
}

func NewPostgresClient(c Config) (*Client, error) {

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", c.Username, c.Password, c.Hostname, c.Database)

	db, err := sql.Open("postgres", connStr)

	return &Client{
		Config: c,
		DB:     db,
	}, err
}
