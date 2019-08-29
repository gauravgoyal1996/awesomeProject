
package client

import (
	"cloud.google.com/go/storage"
	"context"
	"google.golang.org/api/iterator"
	"log"
	)

func GetCloudStorageInfo(projectID string,ctx context.Context) []string{
// Creates a client.
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	var buckets []string
	it := client.Buckets(ctx, projectID)
	for {
		battrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			panic(err)
		}
		buckets = append(buckets, battrs.Name)
	}
	return buckets
}