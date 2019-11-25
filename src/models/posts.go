package models

import "time"

// Posts is DB schema for posts
type Posts struct {
	UserID    string     `json:"user_id"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	Photos    []string   `json:"photos"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (Posts) TableName() string {
	return "posts"
}
