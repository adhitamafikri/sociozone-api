package repositories

import (
	"context"

	"github.com/adhitamafikri/sociozone-api/objects/user"
	DBHelper "github.com/adhitamafikri/sociozone-api/utils/helpers/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UsersRepository struct {
}

func (repository *UsersRepository) GetUsers() (objects.UserResponseObject, error) {
	client := DBHelper.ConnectDB()
	defer DBHelper.DisconnectDB(client)

	collection := client.Database("sociozone").Collection("users")

	var result objects.UserResponseObject
	filter := bson.D{primitive.E{Key: "name", Value: "Stokeley"}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	return result, err
}
