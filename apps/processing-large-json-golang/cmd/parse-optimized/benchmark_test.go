package main

import (
	"log"
	"testing"

	"github.com/iamseki/dev-to/apps/processing-large-json-golang/internal/parser"
)

func BenchmarkParseJSONOptimized(b *testing.B) {
	filename := "../../largefile-20240718-083247-100mb.json"

	log.Println("Starting BenchmarkParseJSONOptimized")
	for i := 0; i < b.N; i++ {
		parser.OptimizedParseFile(filename)
	}
	log.Println("Finished BenchmarkParseJSONOptimized")
}
