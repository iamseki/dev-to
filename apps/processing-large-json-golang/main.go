package main

import "flag"

func main() {

	var targetSizeMB int
	flag.IntVar(&targetSizeMB, "size", 10, "Target size of the JSON file in MB")
	flag.Parse()

	generateFile(targetSizeMB)
}
