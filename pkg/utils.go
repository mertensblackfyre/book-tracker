package pkg

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	ex, err := os.Executable()

	if err != nil {
		log.Fatalln(err)
	}
	// load .env file

	err = godotenv.Load(filepath.Dir(ex) + "/.env")

	if err != nil {
		log.Fatalln(err)
	}

	return os.Getenv(key)
}

// BUKmz7qnLzGFWS@
