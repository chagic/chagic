package main

import (
	"chagic/conf"
	"chagic/log"
	"chagic/router"
	"chagic/server"

	docs "chagic/docs"
	"github.com/joho/godotenv"
)

func main() {
	log.InitLogger(conf.GetConfig().Log.Path, conf.GetConfig().Log.Level)
	log.Logger.Info("config", log.Any("config", conf.GetConfig()))

	docs.SwaggerInfo.BasePath = "/v1"

	// load env
	err := godotenv.Load("./.env")
	if err != nil {
		log.Logger.Info("load env", log.Any("err", err))
	}

	go server.MyServer.Start()
	r := router.NewRouter()
	r.Run(conf.GetConfig().Server.Port)
}
