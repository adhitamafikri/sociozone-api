package services

import (
	"fmt"
	"log"

	"github.com/adhitamafikri/sociozone-api/objects/user"

	"github.com/adhitamafikri/sociozone-api/repositories/users"
)

type UserService struct {
	repository repositories.UsersRepository
}

type IUserService interface {
	Index() objects.UserResponseObject
}

// Index will retrieve all users
func (service *UserService) Index() objects.UserResponseObject {
	fmt.Println("Getting all users")

	result, err := service.repository.GetUsers()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	return result
}
