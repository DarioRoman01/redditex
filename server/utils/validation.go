package utils

import (
	"lireddit/models"
	"strings"
)

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

func GenUserResponseError(field, message string) *models.UserResponse {
	return &models.UserResponse{
		Error: &models.FieldError{
			Field:   field,
			Message: message,
		},
		User: nil,
	}
}

func GenUserResponse(user models.User) *models.UserResponse {
	return &models.UserResponse{
		Error: nil,
		User:  &user,
	}
}
