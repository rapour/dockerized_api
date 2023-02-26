package postgres

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Database string `envconfig:"postgres_database"`
	Username string `envconfig:"postgres_username"`
	Password string `envconfig:"postgres_password"`
	Hostname string `envconfig:"postgres_hostname"`
}

func NewPostgresConfig(prefix string) (Config, error) {

	var conf Config
	err := envconfig.Process(prefix, &conf)

	return conf, err
}
