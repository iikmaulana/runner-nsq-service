package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/iikmaulana/runner-nsq-service/service"
)

type gatewayHandler struct {
	service    *gin.Engine
	nsqUsecase service.NsqUsecase
}

func NewGatewayHandler(svc *gin.Engine,
	nsqUsecase service.NsqUsecase,
) {
	h := gatewayHandler{
		service:    svc,
		nsqUsecase: nsqUsecase,
	}

	h.initNsqUsage()

}
