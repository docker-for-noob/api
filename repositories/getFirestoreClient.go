package repositories

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
)

func CreateClient(ctx context.Context) *firestore.Client {
	// Sets your Google Cloud Platform project ID.
	projectID := "projet-full-stack"

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	// Close client when done with
	// defer client.Close()
	return client
}
