package services

import (
	"fmt"

	"github.com/adhitamafikri/sociozone-api/src/repositories"
)

type PostsService struct {
	repository repositories.PostsRepository
}

func PostsServiceHandler() PostsService {
	service := PostsService{repository: repositories.PostsRepositoryHandler()}
	return service
}

type IPostsService interface {
	GetAllPosts()
}

func (service *PostsService) GetAllPosts() {
	fmt.Print("Getting all posts!")
	service.repository.GetAllPosts()
}
