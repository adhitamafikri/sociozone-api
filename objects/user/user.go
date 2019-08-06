package objects

// User struct represents user object in MongoDB
type User struct {
	Name     string `json:name`
	Username string `json:username`
	Password string `json:password`
}
