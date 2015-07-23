package main

import (
	"flag"
	"fmt"

	"github.com/oschwald/maxminddb-golang"
)

func main() {
	required := "<REQUIRED>"
	file := flag.String("file", required, "The path to the MaxMind DB to verify")

	flag.Parse()

	if *file == required {
		flag.PrintDefaults()
		return
	}

	reader, err := maxminddb.Open(*file)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	if err := reader.Verify(); err != nil {
		fmt.Printf("Error verifying file: %v\n", err)
	}
}
