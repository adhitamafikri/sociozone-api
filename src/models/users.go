package models

import "time"

// User is DB schema for users
type User struct {
	ID             string     `gorm:"primary_key" json: "id"`
	Name           string     `gorm:"column: name" json:"name"`
	Email          string     `gorm:"column: email" json:"email"`
	Username       string     `gorm:"column: username" json:"username"`
	Password       string     `gorm:"column: password" json:"password"`
	ProfilePicture string     `gorm:"column: profile_picture" json:"profile_picture"`
	Bio            string     `gorm:"column: bio" json:"bio"`
	Website        string     `gorm:"column: website" json:"website"`
	CreatedAt      *time.Time `gorm:"column: created_at" json:"created_at"`
	UpdatedAt      *time.Time `gorm:"column: updated_at" json:"updated_at"`
	DeletedAt      *time.Time `gorm:"column: deleted_at" json:"deleted_at"`
}

func (User) TableName() string {
	return "users"
}
