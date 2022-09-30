package main

import (
	"log"
	"os"

	// Blank-import the function package so the init() runs
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	_ "github.com/leafy-ai/backend_b2"
)

func main() {
	// Use PORT environment variable, or default to 8080.
	port := "8080"
	if envPort := os.Getenv("port"); envPort != "" {
		port = envPort
	}
	log.Println("Listening on port", port)
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
