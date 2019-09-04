package objects

// UserResponseObject struct represents user object in API Response
type UserResponseObject struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Username       string `json:"username"`
	ProfilePicture string `json:"profile_picture"`
}

// UserRequestObject struct represents user object in API Request
type UserRequestObject struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
