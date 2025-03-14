package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APIToken string
}

var AppConfig Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	AppConfig.APIToken = os.Getenv("API_TOKEN")
}
