package confing

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	AppPort   string
	RedisHost string
	RedisPort string
	RedisPass string
	RedisDB   int
}

func GetConfig() Config {
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8080"
		fmt.Println("APP_PORT is not set in environment variables. \nUsing default value: 8080")
	}

	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "127.0.0.1"
		fmt.Println("REDIS_HOST is not set in environment variables. \nUsing default value: 127.0.0.1")
	}

	redisPort := os.Getenv("REDIS_PORT")
	if redisPort == "" {
		redisPort = "6379"
		fmt.Println("REDIS_PORT is not set in environment variables. \nUsing default value: 6379")
	}

	redisPass := os.Getenv("REDIS_PASSWORD")
	if redisPass == "" {
		redisPass = "password"
		fmt.Println("REDIS_PASSWORD is not set in environment variables. \nUsing default value: password")
	}

	redisDB := os.Getenv("REDIS_DB")
	if redisDB == "" {
		redisDB = "0"
		fmt.Println("REDIS_DB is not set in environment variables. \nUsing default value: 0")
	}

	redisDBInt, err := strconv.Atoi(redisDB)
	if err != nil {
		redisDBInt = 0
	}

	return Config{
		AppPort:   appPort,
		RedisHost: redisHost,
		RedisPort: redisPort,
		RedisPass: redisPass,
		RedisDB:   redisDBInt,
	}
}
