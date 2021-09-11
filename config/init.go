package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/uzzeet/uzzeet-gateway/controller"
	"github.com/uzzeet/uzzeet-gateway/service"
)

type AppConfig interface {
	Init() Config
}

type Config struct {
	Version  string
	DB       *sqlx.DB
	Registry *controller.Registry
	Server   *service.Server
	Gin      *gin.Engine
	Gateway  *service.Service
	CK       *sqlx.DB
}

func Init() Config {
	var cfg Config

	errx := cfg.initGateway()
	if errx != nil {
		errx.Panic()
	}

	errx = cfg.InitService()
	if errx != nil {
		errx.Panic()
	}

	fmt.Println("Server is running ..")
	return cfg
}
