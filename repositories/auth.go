package repositories

import (
	"context"
	"fmt"
	"log"

	DBHelper "github.com/adhitamafikri/sociozone-api/helpers/db"
	"github.com/adhitamafikri/sociozone-api/objects"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthRepository struct{}

func AuthRepositoryHandler() AuthRepository {
	repository := AuthRepository{}
	return repository
}

type IAuthRepository interface {
	GetLogin()
	PostLogin(username, password string)
	GetRegister()
	PostRegister(name, username, password string)
}

func (repository *AuthRepository) GetLogin() {}

func (repository *AuthRepository) PostLogin(username, password string) {
	client := DBHelper.ConnectDB()
	defer DBHelper.DisconnectDB(client)

	collection := client.Database("sociozone").Collection("users")

	var result objects.UserResponseObject
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

func (repository *AuthRepository) GetRegister() {}

func (repository *AuthRepository) PostRegister(name, username, password string) {
	client := DBHelper.ConnectDB()
	defer DBHelper.DisconnectDB(client)

	collection := client.Database("sociozone").Collection("users")

	userData := objects.UserRequestObject{Name: name, Username: username, Password: password}
	insertResult, err := collection.InsertOne(context.TODO(), userData)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult)
}
