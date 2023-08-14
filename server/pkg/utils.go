package pkg

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	// ex, err := os.Executable()

	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// load .env file

	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatalln(err)
	}

	return os.Getenv(key)
}

// BUKmz7qnLzGFWS@
