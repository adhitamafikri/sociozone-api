package repositories

import (
	"context"

	DBHelper "github.com/adhitamafikri/sociozone-api/src/helpers/db"
	"github.com/adhitamafikri/sociozone-api/src/objects"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository struct{}

func UserRepositoryHandler() UserRepository {
	repository := UserRepository{}
	return repository
}

func (repository *UserRepository) GetAllUsers() (objects.UserResponseObject, error) {
	client := DBHelper.ConnectDB()
	defer DBHelper.DisconnectDB(client)

	collection := client.Database("sociozone").Collection("users")

	var result objects.UserResponseObject
	filter := bson.D{primitive.E{Key: "name", Value: "Stokeley"}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	return result, err
}
