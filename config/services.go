package config

import (
	"github.com/iikmaulana/runner-nsq-service/controller"
	handler2 "github.com/iikmaulana/runner-nsq-service/service/handler"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
)

func (cfg Config) InitService() serror.SError {

	nsqUsecase := controller.NewNsqUsecase()

	handler2.NewGatewayHandler(cfg.Gin, nsqUsecase)
	return nil
}
