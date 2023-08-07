package handlers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {

	err := godotenv.Load(".env")
	// load .env file

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
