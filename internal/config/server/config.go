package server

import (
	"github.com/caarlos0/env/v9"

	"github.com/Albitko/secrets-armgour/internal/entity"
	"github.com/Albitko/secrets-armgour/internal/utils/logger"
)

func Config() (cfg entity.ServerConfig, err error) {
	if err = env.Parse(&cfg); err != nil {
		logger.Zap.Errorf("error: %s reading config from env", err.Error())
		return cfg, err
	}
	return cfg, nil
}
