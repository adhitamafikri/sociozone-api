package services

import (
	"context"
	"fmt"
	"log"

	DBHelper "github.com/adhitamafikri/sociozone-api/helpers/db"
	"github.com/adhitamafikri/sociozone-api/objects/posts"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// RetrievePosts retrieves all posts to be shown in the homepage
func RetrievePosts() {
	fmt.Println("Retrieving all posts")

	client := DBHelper.ConnectDB()
	defer DBHelper.DisconnectDB(client)

	collection := client.Database("sociozone").Collection("posts")

	findOptions := options.Find()
	findOptions.SetLimit(5)

	var results []*objects.Post

	// Find() returns a cursor
	cursor, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate the cursor
	for cursor.Next(context.TODO()) {
		var item objects.Post
		err := cursor.Decode(&item)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &item)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cursor.Close(context.TODO())

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
}

// UploadPost stores user's post into DB
func UploadPost() {
	fmt.Println("Uploading Post!")
}

func uploadPhotos() {
	fmt.Println("upload photos")
}
