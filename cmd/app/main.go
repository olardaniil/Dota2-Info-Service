package main

import (
	"dota2_info_service/confing"
	"dota2_info_service/internal/handler"
	"dota2_info_service/internal/repository"
	"dota2_info_service/internal/service"
	"dota2_info_service/pkg/database"
	"github.com/joho/godotenv"
	"log"
)

// @title Dota2-Info-Service
// @version 1.0
// @description API Server for Dota2-Info-Service

func main() {
	cfg := confing.GetConfig()

	redis, err := database.NewRedis(cfg.RedisHost+":"+cfg.RedisPort, cfg.RedisPass, cfg.RedisDB)
	if err != nil {
		log.Fatalln(err)
	}

	repo := repository.NewRepository(redis)

	serv := service.NewService(repo)

	handlers := handler.NewHandler(serv)

	err = handlers.Run(cfg.AppPort)
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Lshortfile)

	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
}
