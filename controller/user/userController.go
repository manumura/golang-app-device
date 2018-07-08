package user

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/manumura/golang-app-device/model/user"
	"github.com/manumura/golang-app-device/service/user"
	"strconv"
	"database/sql"
)

// UserController : Operations on user
type UserController struct {
	userService userservice.UserService
}

// NewUserController : Create a new instance of UserController
func NewUserController(userService userservice.UserService) *UserController {
	return &UserController{userService}
}

// GetUser : Get user by id as json
func (uc UserController) GetUser(c echo.Context) error {

	idAsString := c.Param("id")
	if idAsString == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	id, err := strconv.Atoi(idAsString)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	user, err := uc.userService.GetUser(id)
	// TODO : throw more appropriate error
	switch {
	case err == sql.ErrNoRows:
		return echo.NewHTTPError(http.StatusNotFound, "Page not found")
	case err != nil:
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, user)
}

// CreateUser : Create new user
func (uc UserController) CreateUser(c echo.Context) error {

	log.Println("CreateUser")

	u := usermodel.User{}
	if err := c.Bind(&u); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}
	log.Println(u)

	if u.ID != 0 || u.Username == "" || u.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	// Check username unique
	unique, err := uc.userService.CheckUsernameUnique(u.Username)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	if !unique {
		return echo.NewHTTPError(http.StatusConflict, "User already exists")
	}

	// Create user
	u, err = uc.userService.Create(u)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, u)
}
