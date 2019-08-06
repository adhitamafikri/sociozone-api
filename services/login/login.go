package services

import (
	"context"
	"fmt"
	"log"

	DBHelper "github.com/adhitamafikri/sociozone-api/helpers/db"
	"github.com/adhitamafikri/sociozone-api/objects/user"
	"go.mongodb.org/mongo-driver/bson"
)

// Login compares data with the user data in the DB, then logs user in
func Login(username, password *string) {
	fmt.Println("Running Login() on login_service!")
	client := DBHelper.ConnectDB()
	collection := client.Database("sociozone").Collection("users")

	var result objects.User
	filter := bson.D{
		bson.E{Key: "Username", Value: *username},
		bson.E{Key: "Password", Value: *password},
	}
	// filter := bson.D{{"Username", *username}, {"Password", *password}}
	// filter := bson.M{"Username": *username, "Password": *password}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)

	DBHelper.DisconnectDB(client)
}
