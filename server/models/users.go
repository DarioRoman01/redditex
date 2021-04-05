package models

import "time"

// User model contains all the data related to the user
// the password is not retrieved in any circunstance
type User struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Username  string    `json:"username" gorm:"unique"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string
	Posts     []Post   `gorm:"foreignKey:CreatorId;references:Id"`
	Updoots   []Updoot `gorm:"foreignKey:UserID;references:Id"`
}

// input type for users signup
type UserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// error field if any error in users login or signup
type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Return a response with user and error
// if no error only return user
// if error return the error and return nil user
type UserResponse struct {
	Error *FieldError `json:"error"`
	User  *User       `json:"user"`
}
