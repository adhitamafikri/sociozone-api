package services

import (
	"context"
	"fmt"
	"log"

	DBHelper "github.com/adhitamafikri/sociozone-api/helpers/db"
	objects "github.com/adhitamafikri/sociozone-api/objects/user"
)

// RegisterUserData will save all registration data to MongoDB
func RegisterUserData(name *string, username *string, password *string) {
	fmt.Println("Running RegisterUserData()")
	client := DBHelper.ConnectDB()
	collection := client.Database("sociozone").Collection("users")

	userData := objects.User{Name: *name, Username: *username, Password: *password}
	insertResult, err := collection.InsertOne(context.TODO(), userData)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	DBHelper.DisconnectDB(client)
}
