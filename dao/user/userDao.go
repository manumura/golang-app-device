package userdao

import (
	"github.com/manumura/golang-app-device/model/user"
	"github.com/manumura/golang-app-device/config"
)

// UserDao : interface defining DB operations on user
type UserDao interface {
	GetUser(id int) (usermodel.User, error)
	GetUserByUsername(username string) (usermodel.User, error)
	CheckUsernameUnique(username string) (bool, error)
	Create(user usermodel.User) (usermodel.User, error)
}

// NewUserDao : Create a new instance of UserDao implemenation
func NewUserDao(db *config.DB) UserDao {
	return UserDaoImpl{db}
}
