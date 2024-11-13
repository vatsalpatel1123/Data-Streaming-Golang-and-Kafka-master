package main

import (
	"log"
	"net/http"
	"your_project_name/config"  // Replace with your actual project/module name
	"your_project_name/router"  // Replace with your actual project/module name
)

func main() {
	// Initialize Kafka Writer
	writer := config.NewKafkaWriter("localhost:9092", "default-topic")

	// Initialize router
	r := router.NewRouter()

	// Start the HTTP server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
