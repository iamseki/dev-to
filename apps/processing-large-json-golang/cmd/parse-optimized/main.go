package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/iamseki/dev-to/apps/processing-large-json-golang/internal/parser"
)

func main() {
	var filename string
	flag.StringVar(&filename, "filename", "defaultfile-1mb.json", "Filename to parse")
	flag.Parse()

	// profiling CPU
	cpu_prof, err := os.Create("cpu-opt.prof")
	if err != nil {
		log.Fatalf("error create cpu.prof: %v", err)
	}
	pprof.StartCPUProfile(cpu_prof)
	defer pprof.StopCPUProfile()

	parser.OptimizedParseFile(filename)

	// profiling MEM
	mem_prof, err := os.Create("mem-opt.prof")
	if err != nil {
		log.Fatalf("error create mem.prof: %v", err)
	}
	defer mem_prof.Close()

	runtime.GC() // get up-to-date statistics
	if err := pprof.WriteHeapProfile(mem_prof); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
}
