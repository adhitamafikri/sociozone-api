package services

import (
	"fmt"
	"log"

	"github.com/adhitamafikri/sociozone-api/objects"
	"github.com/adhitamafikri/sociozone-api/repositories"
)

type UserService struct {
	repository repositories.UserRepository
}

func UserServiceHandler() UserService {
	service := UserService{repository: repositories.UserRepositoryHandler()}
	return service
}

type IUserService interface {
	Index() objects.UserResponseObject
}

// Index will retrieve all users
func (service *UserService) GetAllUsers() objects.UserResponseObject {
	fmt.Println("Getting all users")

	result, err := service.repository.GetAllUsers()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	return result
}
