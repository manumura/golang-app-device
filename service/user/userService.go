package userservice

import (
	"github.com/manumura/golang-app-device/model/user"
)

// UserService : interface defining services on user
type UserService interface {
	GetUser(id int) (usermodel.User, error)
	GetUserByUsername(username string) (usermodel.User, error)
	CheckUsernameUnique(username string) (bool, error)
	Create(user usermodel.User) (usermodel.User, error)
}
