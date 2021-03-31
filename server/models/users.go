package models

import "time"

type User struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Username  string    `json:"username" gorm:"unique"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string
}

type UserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type UserResponse struct {
	Error *FieldError `json:"error"`
	User  *User       `json:"user"`
}
