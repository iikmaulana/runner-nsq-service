package service

import (
	"github.com/iikmaulana/runner-nsq-service/models"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
)

type NsqUsageRepo interface {
	SenderNSQRepo(form models.ViewAllTruckRequest) (serr serror.SError)
}
