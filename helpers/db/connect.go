package helpers

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB will do connection to MongoDB
func ConnectDB() *mongo.Client {
	// set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to DB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	// Check connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// // Test Connection
	// collection := client.Database("sociozone").Collection("users")
	// fmt.Println(collection)

	return client
}

// DisconnectDB will disconnect program from MongoDB
func DisconnectDB(client *mongo.Client) {
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
