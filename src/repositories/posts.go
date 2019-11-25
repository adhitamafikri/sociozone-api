package repositories

import (
	"context"
	"fmt"
	"log"

	DBHelper "github.com/adhitamafikri/sociozone-api/src/helpers/db"
	"github.com/adhitamafikri/sociozone-api/src/objects"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PostsRepository struct{}

func PostsRepositoryHandler() PostsRepository {
	repository := PostsRepository{}
	return repository
}

type IPostsRepository interface {
	GetAllPosts()
}

func (repository *PostsRepository) GetAllPosts() {
	client := DBHelper.ConnectDB()
	defer DBHelper.DisconnectDB(client)

	collection := client.Database("sociozone").Collection("posts")

	findOptions := options.Find()
	findOptions.SetLimit(5)

	var results []*objects.PostsResponseObject

	// Find() returns a cursor
	cursor, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate the cursor
	for cursor.Next(context.TODO()) {
		var item objects.PostsResponseObject
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
