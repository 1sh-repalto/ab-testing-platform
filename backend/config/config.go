package config 

import (
	"log"
	"os"

	"github.com/joho/dotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}