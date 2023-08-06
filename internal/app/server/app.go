package server

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/Albitko/secrets-armgour/internal/adapter/repository"
	"github.com/Albitko/secrets-armgour/internal/config"
	"github.com/Albitko/secrets-armgour/internal/controller/handler"
	"github.com/Albitko/secrets-armgour/internal/secrets/processor"
	"github.com/Albitko/secrets-armgour/internal/utils/logger"
)

// Run - server side application
func Run() {
	log := logger.Init()
	appCfg, err := config.NewServerFromEnv(log)

	if err != nil {
		panic(fmt.Errorf("can't configure application: %w", err))
	}

	repo, err := repository.New(appCfg.DatabaseDsn, log)
	if err != nil {
		panic(fmt.Errorf("can't connecto to DB: %w", err))
	}

	defer repo.Close()

	secretsProcessor := processor.New(repo)
	h := handler.New(secretsProcessor)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.POST("/v1/user/login", h.Login)
	router.POST("/v1/user/register", h.Register)
	router.GET("/v1/secrets/get/:data/:id/:user", h.Get)
	router.GET("/v1/secrets/list/:data/:user", h.List)
	router.DELETE("/v1/secrets/:data/:id", h.Delete)
	router.POST("/v1/secrets/credentials/:user", h.CredentialsCreate)
	router.POST("/v1/secrets/text/:user", h.TextCreate)
	router.POST("/v1/secrets/binary/:user", h.BinaryCreate)
	router.POST("/v1/secrets/card/:user", h.CardCreate)
	router.PUT("/v1/secrets/credentials/:id", h.CredentialsEdit)
	router.PUT("/v1/secrets/text/:id", h.TextEdit)
	router.PUT("/v1/secrets/binary/:id", h.BinaryEdit)
	router.PUT("/v1/secrets/card/:id", h.CardEdit)

	err = router.Run(appCfg.ServerAddr)
	if err != nil {
		panic(fmt.Errorf("start server failed: %w", err))
	}
}
