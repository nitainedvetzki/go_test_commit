package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dapr/dapr/pkg/client"
)

func main() {
	// Initialize Dapr client
	daprClient, err := client.NewClient()
	if err != nil {
		log.Fatalf("Error initializing Dapr client: %v", err)
	}
	defer daprClient.Close()

	// Publish a message to a Dapr topic
	topic := "my-topic"
	data := []byte("Hello, Dapr!")

	err = daprClient.PublishEvent(context.Background(), topic, "myOperation", data)
	if err != nil {
		log.Fatalf("Error publishing message: %v", err)
	}

	fmt.Println("Message published successfully.")
}

