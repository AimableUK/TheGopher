package services

import (
	"context"
	"fmt"
	"go-mongo-api/configs"

	"go.mongodb.org/mongo-driver/mongo"
)

func WatchBooksChanges() {
	collection := configs.GetCollection(configs.DB, "books")
	stream, err := collection.Watch(context.TODO(), mongo.Pipeline{})
	if err != nil {
		fmt.Println("Error starting change stream:", err)
		return
	}
	defer stream.Close(context.Background())

	fmt.Println("Watching changes on books collection...")

	for stream.Next(context.Background()) {
		var changeEvent map[string]interface{}
		if err := stream.Decode(&changeEvent); err != nil {
			fmt.Println("Error decoding change event:", err)
			continue
		}

		fmt.Println("Change detected:", changeEvent)
	}
}
