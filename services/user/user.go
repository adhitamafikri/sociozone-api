package services

import (
	"context"
	"fmt"
	"log"

	DBHelper "github.com/adhitamafikri/sociozone-api/helpers/db"
	"github.com/adhitamafikri/sociozone-api/objects/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Index will retrieve all users
func Index() {
	fmt.Println("Getting all users")

	client := DBHelper.ConnectDB()
	defer DBHelper.DisconnectDB(client)

	collection := client.Database("sociozone").Collection("users")

	var result objects.User
	filter := bson.D{primitive.E{Key: "name", Value: "Myree Bowden"}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
