package services

import (
	"context"
	"fmt"
	"log"

	DBHelper "github.com/adhitamafikri/sociozone-api/utils/helpers/db"
	"github.com/adhitamafikri/sociozone-api/objects/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Login compares data with the user data in the DB, then logs user in
func Login(username, password string) {
	fmt.Println("Running Login() on login_service!")
	client := DBHelper.ConnectDB()
	defer DBHelper.DisconnectDB(client)

	collection := client.Database("sociozone").Collection("users")

	var result objects.User
	filter := bson.D{
		primitive.E{Key: "username", Value: username},
		primitive.E{Key: "password", Value: password},
	}

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)
}
