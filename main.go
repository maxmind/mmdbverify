// mmdbverify verifies a MaxMind DB file to ensure that is valid.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/oschwald/maxminddb-golang/v2"
)

func main() {
	required := "<REQUIRED>"
	file := flag.String("file", required, "The path to the MaxMind DB to verify")

	flag.Parse()

	if *file == required {
		flag.PrintDefaults()
		os.Exit(1)
	}

	reader, err := maxminddb.Open(*file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}

	if err := reader.Verify(); err != nil {
		reader.Close()
		fmt.Fprintf(os.Stderr, "Error verifying file: %v\n", err)
		os.Exit(1)
	}
	reader.Close()
}
