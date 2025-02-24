package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	port := "8080"

	serveMux := http.ServeMux{}

	server := &http.Server{
		Handler: &serveMux,
		Addr:    ":" + port,
	}

	workingDirectory, err := os.Getwd()
	if err != nil {
		fmt.Printf("error getting the working directory: %v", err)
	}

	filepathRoot, err := filepath.Abs(workingDirectory + "/files")
	if err != nil {
		fmt.Printf("error getting the filepath: %v", err)
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(server.ListenAndServe())
}
