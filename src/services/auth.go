package services

import (
	"fmt"

	"github.com/adhitamafikri/sociozone-api/src/repositories"
)

type AuthService struct {
	repository repositories.AuthRepository
}

func AuthServiceHandler() AuthService {
	service := AuthService{repositories.AuthRepositoryHandler()}
	return service
}

type IAuthService interface {
	GetLogin()
	PostLogin(username, password string)
	GetRegister()
	PostRegister(name, username, password string)
}

func (service *AuthService) GetLogin() {}

func (service *AuthService) PostLogin(username, password string) {
	fmt.Println("Running Posts() on AuthService!")
	service.repository.PostLogin(username, password)
}

func (service *AuthService) PostRegister(name, username, password string) {
	fmt.Println("Running RegisterUserData() on AuthService!")
	service.repository.PostRegister(name, username, password)
}
