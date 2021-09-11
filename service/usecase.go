package service

import (
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
)

type NsqUsecase interface {
	SenderNSQUsecase(form []byte) (serr serror.SError)
}
