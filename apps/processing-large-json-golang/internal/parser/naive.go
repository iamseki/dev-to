package parser

import (
	"encoding/json"
	"log"
	"os"

	"github.com/iamseki/dev-to/apps/processing-large-json-golang/internal"
)

func NaiveParseFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error on open file: %s, error: %v", filename, err)
	}
	defer file.Close()

	users := []internal.User{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&users); err != nil {
		log.Fatalf("error on decode users from json file: %v", err)
	}

	for _, user := range users {
		log.Println(user)
	}
}
