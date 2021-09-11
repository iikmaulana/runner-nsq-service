package main

import (
	"github.com/iikmaulana/runner-nsq-service/config"
	"github.com/joho/godotenv"
	"github.com/uzzeet/uzzeet-gateway/libs"
	"github.com/uzzeet/uzzeet-gateway/libs/helper"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.Init()

	err = cfg.Gin.Run(":" + helper.Env(libs.AppPort, "6001"))
	if err != nil {
		log.Fatal(err)
	}
}
