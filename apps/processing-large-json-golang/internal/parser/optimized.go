package parser

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	"github.com/iamseki/dev-to/apps/processing-large-json-golang/internal"
)

func OptimizedParseFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("error opening file: ", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)

	token, err := decoder.Token()
	if err != nil {
		log.Fatalln("Error reading opening token: ", err)
	}

	if delim, ok := token.(json.Delim); !ok || delim != '[' {
		log.Fatalln("Expeceted start of JSON array")
	}

	for decoder.More() {
		user := &internal.User{}
		err := decoder.Decode(user)
		if err != nil {
			log.Fatalln("Error decoding JSON user: ", err)
		}

		// DO STUFF
	}

	// Read the closing bracket of the JSON array
	token, err = decoder.Token()
	if err != nil {
		log.Fatalln("Error reading closing token:", err)
	}

	// Check if the closing token is the end of the array
	if delim, ok := token.(json.Delim); !ok || delim != ']' {
		log.Fatalln("Expected end of JSON array")
	}

}
