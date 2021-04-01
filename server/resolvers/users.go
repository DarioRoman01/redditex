package resolvers

import (
	"context"
	"fmt"

	"lireddit/cache"
	"lireddit/db"
	"lireddit/models"
	"lireddit/utils"
	"strings"
)

var userTable db.UsersTable

func init() {
	psql, _ := db.Connect()
	userTable = db.UsersTable{
		Table: *psql,
	}
}

// handle users signup and validate all the fields required for signup
func (r *mutationResolver) Register(ctx context.Context, options models.UserInput) (*models.UserResponse, error) {
	isInvalid := utils.ValidateRegister(options)
	if isInvalid != nil {
		return isInvalid, nil
	}

	userResponse := userTable.SingUp(models.User{
		Username: options.Username,
		Email:    options.Email,
		Password: options.Password,
	},
	)

	return userResponse, nil
}

// Handle users login by email or username
func (r *mutationResolver) Login(ctx context.Context, usernameOrEmail string, password string) (*models.UserResponse, error) {
	if strings.Contains(usernameOrEmail, "@") {

		response := userTable.LoginByEmail(usernameOrEmail, password)
		if response.Error != nil {
			return response, nil
		}

		ec, err := utils.EchoContextFromContext(ctx)
		if err != nil {
			return utils.GenUserResponseError("server", "internal server error"), nil
		}

		session := cache.Default(ec)
		session.Set("userId", response.User.Username)
		session.Save()
		return response, nil

	} else {
		response := userTable.LoginByUsername(usernameOrEmail, password)
		if response.Error != nil {
			return response, nil
		}
		ec, err := utils.EchoContextFromContext(ctx)
		if err != nil {
			fmt.Println(err)
			return utils.GenUserResponseError("server", "internal server error"), nil
		}
		session := cache.Default(ec)
		session.Set("userId", response.User.Username)
		session.Save()

		return response, nil
	}
}

// Logout remove session cookie from the client
func (m *mutationResolver) Logout(ctx context.Context) (bool, error) {
	ec, err := utils.EchoContextFromContext(ctx)
	if err != nil {
		fmt.Println(err)
		return false, fmt.Errorf("unable to logout")
	}

	session := cache.Default(ec)
	session.Clear()
	session.Save()

	return true, nil
}

// return user data based on the session
func (r *queryResolver) Me(ctx context.Context) (*models.User, error) {
	ec, err := utils.EchoContextFromContext(ctx)
	if err != nil {
		return nil, nil
	}

	session := cache.Default(ec)
	val := session.Get("userId")
	user := userTable.GetUserByUsername(val)
	return user, nil
}

// handle forgot password validate email exist in the db and send email for change password
func (m *mutationResolver) ForgotPassword(ctx context.Context, email string) (bool, error) {
	user := userTable.GetUserByEmail(email)

	if user == nil {
		return true, nil
	}

	success := utils.SendEmail(email)

	if !success {
		return false, fmt.Errorf("unable to send email")
	}

	return true, nil
}
