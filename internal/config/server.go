package config

import (
	"github.com/caarlos0/env/v9"
)

type logger interface {
	Errorf(template string, args ...interface{})
}

type serverConfig struct {
	DatabaseDsn string `env:"DB_DSN" envDefault:"postgres://postgres@localhost:5432/postgres"`
	ServerAddr  string `env:"SRV_ADDR" envDefault:"0.0.0.0:8080"`
}

func NewServerFromEnv(l logger) (cfg serverConfig, err error) {
	if err = env.Parse(&cfg); err != nil {
		l.Errorf("error: %s reading config from env", err.Error())
		return cfg, err
	}
	return cfg, nil
}
