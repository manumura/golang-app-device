package userservice

import (
	"github.com/manumura/golang-app-device/dao/user"
	"github.com/manumura/golang-app-device/model/user"
)

// UserServiceImpl : implementation for services on device type
type UserServiceImpl struct {
	userDao userdao.UserDao
}

// NewUserService : Create a new instance of UserService implemenation
func NewUserService(userdao userdao.UserDao) UserService {
	return UserServiceImpl{userdao}
}

// GetUser : retrieve user by id
func (us UserServiceImpl) GetUser(id int) (usermodel.User, error) {

	user, err := us.userDao.GetUser(id)
	return user, err
}

// GetUserByUsername : retrieve user by username
func (us UserServiceImpl) GetUserByUsername(username string) (usermodel.User, error) {

	user, err := us.userDao.GetUserByUsername(username)
	return user, err
}

// CheckUsernameUnique : check username is unique
func (us UserServiceImpl) CheckUsernameUnique(username string) (bool, error) {

	unique, err := us.userDao.CheckUsernameUnique(username)
	return unique, err
}

// Create : create one user
func (us UserServiceImpl) Create(user usermodel.User) (usermodel.User, error) {

	userCreated, err := us.userDao.Create(user)
	return userCreated, err
}
