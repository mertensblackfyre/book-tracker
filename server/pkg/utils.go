package pkg

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln(err)
	}

	return os.Getenv(key)
}

// BUKmz7qnLzGFWS@
