package pkg

import (
	"fmt"
	"log"
	"os"
	"encoding/json"
	"io/ioutil"


	"github.com/joho/godotenv"
)

func GetEnv(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln(err)
	}

	return os.Getenv(key)
}

func JSONStruct(file string) []Book {
	// Open JSON file
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	// Read opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var b []Book
	err = json.Unmarshal(byteValue, &b)
	if err != nil {
		fmt.Println(err)
	}
	return b
}

// BUKmz7qnLzGFWS@
