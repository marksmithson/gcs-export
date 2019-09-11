/*
Fetch a file from Google Cloud Storage

Usage:
	gcs-fetch gs://bucket/object output-file
 */
package main

import (
	"fmt"
	"github.com/marksmithson/gcs-export/internal/pkg/gcsexport"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		printUsage()
		os.Exit(1)
	}

	inputFilename := os.Args[1]
	gsObject := os.Args[2]


	if gsObject == "" || inputFilename == "" {
		printUsage()
		os.Exit(1)
	}
	inputReader, err := os.Open(inputFilename)
	if err != nil { log.Fatal(err) }
	defer inputReader.Close()

	bytes, err := gcsexport.Export(inputReader, gsObject)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Copied %d Bytes\n", bytes)
}

func printUsage() {
	fmt.Println("Usage: gcs-export local-file gs://bucket/object")
}