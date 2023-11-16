package main

import (
	"chagic/conf"
	"chagic/log"
	"chagic/router"
	"chagic/server"

	"github.com/joho/godotenv"
)

func main() {
	log.InitLogger(conf.GetConfig().Log.Path, conf.GetConfig().Log.Level)
	log.Logger.Info("config", log.Any("config", conf.GetConfig()))

	// load env
	err := godotenv.Load("../.env")
	if err != nil {
		log.Logger.Info("load env", log.Any("err", err))
	}

	go server.MyServer.Start()
	r := router.NewRouter()
	r.Run(conf.GetConfig().Server.Port)
}
