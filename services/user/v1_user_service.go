package services

import (
	"context"
	"fmt"
	"log"

	DBHelper "github.com/adhitamafikri/sociozone-api/helpers/db"
	"github.com/adhitamafikri/sociozone-api/objects/user"
	"go.mongodb.org/mongo-driver/bson"
)

// Index will retrieve all users
func Index() {
	fmt.Println("Getting all users")

	client := DBHelper.ConnectDB()
	collection := client.Database("sociozone").Collection("users")

	var result objects.User
	filter := bson.D{{"name", "Admin"}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
	DBHelper.DisconnectDB(client)
}
