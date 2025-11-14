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
	verbose := flag.Bool("verbose", false, "Print verification status messages")

	flag.Parse()

	if *file == required {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *verbose {
		fmt.Printf("Verifying %s...\n", *file)
	}

	reader, err := maxminddb.Open(*file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening %s: %v\n", *file, err)
		os.Exit(1)
	}

	if err := reader.Verify(); err != nil {
		reader.Close()
		fmt.Fprintf(os.Stderr, "Error verifying %s: %v\n", *file, err)
		os.Exit(1)
	}
	reader.Close()

	if *verbose {
		fmt.Printf("%s is valid\n", *file)
	}
}
