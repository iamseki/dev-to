package generator

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/iamseki/dev-to/apps/processing-large-json-golang/internal"
)

// This function will try to generate the file given the limit of mb with random users
// and will panic if some error occurs
func GenerateFile(targetSizeMB int) {
	rand.Seed(time.Now().UnixNano())

	file := createFile(targetSizeMB)
	defer file.Close()

	targetSizeBytes := targetSizeMB * 1024 * 1024

	file.WriteString("[")

	size := 1
	id := 1

	for size < targetSizeBytes {
		user := internal.User{
			ID:   id,
			Name: randomString(10),
			Age:  rand.Intn(100),
		}

		bytes, err := json.Marshal(user)
		if err != nil {
			log.Fatalf("error Marshal user: %v", err)
		}

		if size+len(bytes)+1 >= targetSizeBytes {
			break
		}

		if id > 1 {
			file.WriteString(",")
			size += 1
		}

		file.Write(bytes)
		size += len(bytes)
		id += 1
	}

	file.WriteString("]") // End of JSON array
}

func createFile(targetSizeMB int) *os.File {
	timestamp := time.Now().Format("20060102-150405")
	filename := fmt.Sprintf("largefile-%s-%vmb.json", timestamp, targetSizeMB)

	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("error trying to create the file: %v", err)
	}

	return file
}

func randomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
