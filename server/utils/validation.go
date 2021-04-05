package utils

import (
	"bytes"
	"lireddit/models"
	"strings"
)

// validate all the fields in the users registration
func ValidateRegister(options models.UserInput) *models.UserResponse {
	if strings.Contains(options.Username, "@") {
		return &models.UserResponse{
			Error: &models.FieldError{
				Field:   "username",
				Message: "@ not allowed in username",
			},
		}
	}

	if len(options.Username) < 3 {
		return &models.UserResponse{
			Error: &models.FieldError{
				Field:   "username",
				Message: "username must be at least 3 characters",
			},
		}
	}

	if !strings.Contains(options.Email, "@") || !strings.Contains(options.Email, ".") {
		return &models.UserResponse{
			Error: &models.FieldError{
				Field:   "email",
				Message: "Invalid email",
			},
		}
	}

	if len(options.Password) <= 3 {
		return &models.UserResponse{
			Error: &models.FieldError{
				Field:   "password",
				Message: "Password must be at least 4 characters",
			},
		}
	}

	return nil
}

// Generate user response with error if there is any error
func GenUserResponseError(field, message string) *models.UserResponse {
	return &models.UserResponse{
		Error: &models.FieldError{
			Field:   field,
			Message: message,
		},
		User: nil,
	}
}

// Generate a user reponse with the user
// there was no errors in the request
func GenUserResponse(user models.User) *models.UserResponse {
	return &models.UserResponse{
		Error: nil,
		User:  &user,
	}
}

// split string by length
func SplitSubN(s string, n int) []string {
	sub := ""
	subs := []string{}

	runes := bytes.Runes([]byte(s))
	l := len(runes)
	for i, r := range runes {
		sub = sub + string(r)
		if (i+1)%n == 0 {
			subs = append(subs, sub)
			sub = ""
		} else if (i + 1) == l {
			subs = append(subs, sub)
		}
	}

	return subs
}
