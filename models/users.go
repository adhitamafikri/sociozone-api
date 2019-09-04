package models

import "time"

// User is DB schema for users
type User struct {
	Name           string     `json:"name"`
	Email          string     `json:"email"`
	Username       string     `json:"username"`
	Password       string     `json:"password"`
	ProfilePicture string     `json:"profile_picture"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
}

func (User) TableName() string {
	return "users"
}
