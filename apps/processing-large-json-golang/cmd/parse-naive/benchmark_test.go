package main

import (
	"log"
	"testing"

	"github.com/iamseki/dev-to/apps/processing-large-json-golang/internal/parser"
)

func BenchmarkParseJSONNaive(b *testing.B) {
	filename := "../../largefile-20240718-083247-100mb.json"

	log.Println("Starting BenchmarkParseJSONNaive")
	for i := 0; i < b.N; i++ {
		parser.NaiveParseFile(filename)
	}
	log.Println("Finished BenchmarkParseJSONNaive")
}
