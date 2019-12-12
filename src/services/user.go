package services

type UserService struct {
}

func UserServiceHandler() UserService {
	service := UserService{}
	return service
}

// type IUserService interface {
// 	Index() objects.UserResponseObject
// }

// Index will retrieve all users
// func (service *UserService) GetAllUsers() objects.UserResponseObject {
// 	fmt.Println("Getting all users")

// 	result, err := service.repository.GetAllUsers()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(result)

// 	return result
// }
