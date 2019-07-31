package services

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

// Index will retrieve all users
func Index() {
	fmt.Println("Getting all users")
	data := bson.M{"name": "Admin"}
	fmt.Println(data)
}
