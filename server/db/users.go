package db

import (
	"lireddit/models"
	"lireddit/utils"

	"gorm.io/gorm"
)

type UsersTable struct {
	Table gorm.DB
}

var passwdCfg *utils.PasswordConfig

func init() {
	passwdCfg = &utils.PasswordConfig{
		Time:    1,
		Memory:  64 * 1024,
		Threads: 4,
		KeyLen:  32,
	}
}

// Handle users creation and validation for unique fields
func (u *UsersTable) SingUp(user models.User) *models.UserResponse {
	usernameTaken := u.GetUserByUsername(user.Username)
	if usernameTaken != nil {
		return utils.GenUserResponseError("username", "username already taken")
	}

	emailTaken := u.GetUserByEmail(user.Email)
	if emailTaken != nil {
		return utils.GenUserResponseError("email", "email already in use")
	}

	hashPasswd, err := utils.GeneratePassword(passwdCfg, user.Password)
	if err != nil {
		return utils.GenUserResponseError("server", "unable to hash password")
	}

	user.Password = hashPasswd
	result := u.Table.Create(&user)
	if result.RowsAffected == 0 || result.Error != nil {
		return utils.GenUserResponseError("user", "unable to create user")
	}

	return utils.GenUserResponse(user)
}

// Login users by Email and check credentials
func (u *UsersTable) LoginByEmail(email, password string) *models.UserResponse {
	storeUser := u.GetUserByEmail(email)
	if storeUser == nil {
		return utils.GenUserResponseError("usernameOrEmail", "that email does not exist")
	}

	ok, err := utils.ComparePasswords(password, storeUser.Password)
	if err != nil {
		return utils.GenUserResponseError("server", "internal server error")
	}

	if !ok {
		return utils.GenUserResponseError("password", "invalid credentials")
	}
	return utils.GenUserResponse(*storeUser)
}

// Login users by username and check credentials
func (u *UsersTable) LoginByUsername(username, password string) *models.UserResponse {
	storeUser := u.GetUserByUsername(username)
	if storeUser == nil {
		return utils.GenUserResponseError("usernameOrEmail", "that username does not exist")
	}

	ok, err := utils.ComparePasswords(password, storeUser.Password)
	if err != nil {
		return utils.GenUserResponseError("server", "internal server error")
	}

	if !ok {
		return utils.GenUserResponseError("password", "invalid credentials")
	}

	return utils.GenUserResponse(*storeUser)
}

// retrieve users by username. if not exist only return null
func (u *UsersTable) GetUserByUsername(username interface{}) *models.User {
	var user models.User

	u.Table.Find(&models.User{}).Where("username = ?", username).Find(&user)

	if user.Id == 0 {
		return nil
	}

	return &user
}

// retrieve users by email. if not exist only return null
func (u *UsersTable) GetUserByEmail(email string) *models.User {
	var user models.User

	u.Table.Model(&models.User{}).Where("email = ?", email).Find(&user)

	if user.Id == 0 {
		return nil
	}

	return &user
}
