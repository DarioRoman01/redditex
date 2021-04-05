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

// validate email field in users resolvers
func (u *UserResolver) Email(ctx context.Context, obj *models.User) (string, error) {
	val, err := utils.GetUserSession(ctx)

	// current user is not logged in
	if err != nil {
		return "", nil
	}

	// current user is logged in and wants to se someone elses email
	if val != obj.Id {
		return "", nil
	}

	// this is the current user and its ok to show them their email
	return obj.Email, nil
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

		if err := utils.GenerateSession(ctx, response.User.Id); err != nil {
			return utils.GenUserResponseError("server", "unable to create session"), nil
		}

		return response, nil

	} else {
		response := userTable.LoginByUsername(usernameOrEmail, password)
		if response.Error != nil {
			return response, nil
		}

		if err := utils.GenerateSession(ctx, response.User.Id); err != nil {
			return utils.GenUserResponseError("server", "unable to create session"), nil
		}

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
	val, err := utils.GetUserSession(ctx)

	if err != nil {
		return nil, nil
	}

	user := userTable.GetUserByid(val)
	return user, nil
}

// handle forgot password validate email exist in the db and send email for change password
func (m *mutationResolver) ForgotPassword(ctx context.Context, email string) (bool, error) {
	user := userTable.GetUserByEmail(email)

	if user == nil {
		return true, nil
	}

	success := utils.SendEmail(user)

	if !success {
		return false, fmt.Errorf("unable to send email")
	}

	return true, nil
}

// Handle change password request validate token and password. also if the request go well create a session for the user
func (m *mutationResolver) ChangePassword(ctx context.Context, token string, newPassword string) (*models.UserResponse, error) {
	if len(newPassword) < 3 {
		return utils.GenUserResponseError("newPassword", "password must at least 3 characters"), nil
	}

	key := fmt.Sprintf("forgot-password:%s", token)
	redis := cache.ConnectRedis()
	userId := redis.Get(ctx, key)

	if userId.Val() == "" {
		return utils.GenUserResponseError("token", "Token expired"), nil
	}

	response := userTable.ChangeUserPassword((userId.Val()), newPassword)
	if response.Error != nil {
		return response, nil
	}

	id, _ := userId.Int()
	if err := utils.GenerateSession(ctx, id); err != nil {
		return utils.GenUserResponseError("server", err.Error()), nil
	}

	redis.Del(context.Background(), key)
	return response, nil
}
