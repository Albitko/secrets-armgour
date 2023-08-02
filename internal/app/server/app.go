package server

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/Albitko/secrets-armgour/internal/adapter/repository"
	"github.com/Albitko/secrets-armgour/internal/config/server"
	"github.com/Albitko/secrets-armgour/internal/controller/handler"
	"github.com/Albitko/secrets-armgour/internal/secrets/processor"
	"github.com/Albitko/secrets-armgour/internal/utils/logger"
)

func Run() {
	appCfg, err := server.Config()
	logger.Init()
	if err != nil {
		panic(fmt.Errorf("can't configure application: %w", err))
	}

	repo, err := repository.New(appCfg.DatabaseDsn)
	defer repo.Close()
	if err != nil {
		panic(fmt.Errorf("can't connecto to DB: %w", err))
	}

	secretsProcessor := processor.New(repo)
	h := handler.New(secretsProcessor)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/v1/user/login", h.Login)
	router.GET("/v1/user/logout", h.Logout)
	router.GET("/v1/user/register", h.Register)
	router.GET("/v1/secrets/list", h.List)
	router.GET("/v1/secrets/get", h.Get)
	router.POST("/v1/secrets/credentials/create", h.CredentialsCreate)
	router.PUT("/v1/secrets/edit", h.Edit)
	router.DELETE("/v1/secrets/del", h.Delete)
	router.GET("/v1/secrets/generate_password", h.GeneratePassword)

	err = router.Run(appCfg.ServerAddr)
	if err != nil {
		panic(fmt.Errorf("start server failed: %w", err))
	}
}
