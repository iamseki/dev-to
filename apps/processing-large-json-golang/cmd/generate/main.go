package main

import (
	"flag"

	"github.com/iamseki/dev-to/apps/processing-large-json-golang/internal/generator"
)

func main() {

	var targetSizeMB int
	flag.IntVar(&targetSizeMB, "size", 10, "Target size of the JSON file in MB")
	flag.Parse()

	generator.GenerateFile(targetSizeMB)
}
