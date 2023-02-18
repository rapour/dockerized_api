package http

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port int `envconfig:"http_port"`
}

func NewHttpConfig(prefix string) (Config, error) {

	var conf Config
	err := envconfig.Process(prefix, &conf)

	return conf, err
}
