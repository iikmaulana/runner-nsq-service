package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/iikmaulana/runner-nsq-service/lib"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
	"io/ioutil"
	"net/http"
)

var appid = "uzzeet"
var svcid = "nsq"
var controller = "nsq-usage"

func (ox gatewayHandler) SenderNsq(ctx *gin.Context) {

	res, errx := lib.ClaimToken(ctx.Request.Header["Authorization"])
	if errx != nil {
		serr := serror.NewFromError(errx)
		lib.Response(http.StatusBadRequest, serr.Error(), appid, svcid, serr.File(), ctx.Request.Method, "", ctx)
		return
	}
	if res.Username != "nsq" {
		serr := serror.NewFromError(errx)
		lib.Response(http.StatusUnauthorized, serr.Error(), appid, svcid, serr.File(), ctx.Request.Method, "", ctx)
		return
	}

	jsonData, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		serr := serror.NewFromError(err)
		lib.Response(http.StatusBadRequest, serr.Error(), appid, svcid, serr.File(), ctx.Request.Method, "", ctx)
		return
	}

	err = ox.nsqUsecase.SenderNSQUsecase(jsonData)
	if err != nil {
		serr := serror.NewFromError(err)
		lib.Response(http.StatusNotImplemented, serr.Error(), appid, svcid, serr.File(), ctx.Request.Method, "", ctx)
		return
	}

	lib.Response(http.StatusOK, "", appid, svcid, controller, ctx.Request.Method, "success sender nsq", ctx)
	return
}
