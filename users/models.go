package users

import "time"

// User is a person
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// UserWonder is the wonder that a user independently did
type UserWonder struct {
	ID          int       `json:"id"`
	UserID      int       `json:"userId"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
}
