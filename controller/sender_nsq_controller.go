package controller

import (
	"github.com/iikmaulana/runner-nsq-service/service"
	"github.com/nsqio/go-nsq"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
	"log"
)

type nsqUsecase struct {
}

func NewNsqUsecase() service.NsqUsecase {
	return nsqUsecase{}
}

func (f nsqUsecase) SenderNSQUsecase(form []byte) (serr serror.SError) {

	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("192.168.9.171:4150", config)
	if err != nil {
		log.Fatal(err)
	}

	topic := "test_satu"
	err = producer.Publish(topic, form)
	if err != nil {
		return serror.NewFromError(err)
	}

	return nil
}
