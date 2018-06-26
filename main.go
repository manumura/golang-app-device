package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	// "github.com/julienschmidt/httprouter"

	jwt "github.com/dgrijalva/jwt-go"
	// "github.com/facebookgo/grace/gracehttp"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/manumura/golang-app-device/controller/channel"
	"github.com/manumura/golang-app-device/controller/device"
	"github.com/manumura/golang-app-device/controller/device-type"
	"github.com/manumura/golang-app-device/controller/user"
	"github.com/manumura/golang-app-device/dao/channel"
	"github.com/manumura/golang-app-device/dao/device"
	"github.com/manumura/golang-app-device/dao/device-type"
	"github.com/manumura/golang-app-device/dao/user"
	"github.com/manumura/golang-app-device/security"
	"github.com/manumura/golang-app-device/service/channel"
	"github.com/manumura/golang-app-device/service/device"
	"github.com/manumura/golang-app-device/service/device-type"
	"github.com/manumura/golang-app-device/service/user"
	"github.com/tylerb/graceful"
	// "github.com/rs/cors"
)

// TODO : DI for database
// Application starts here.
func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// r := httprouter.New()
	// r.GET("/", index)
	e := echo.New()

	corsConfig := middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:8091"},
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowCredentials: true,
	}
	e.Use(middleware.CORSWithConfig(corsConfig))

	// TODO
	// sess := sessions.NewCookieStore([]byte("secret"))
	// sess.Options = &sessions.Options{
	// 	Path:     "/",
	// 	MaxAge:   86400 * 7,
	// 	HttpOnly: true,
	// }
	// e.Use(session.Middleware(sess))

	apiV1Group := e.Group("/api/v1")
	// this logs the server interaction
	apiV1Group.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}]  ${status}  ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	apiV1Group.GET("/", index)

	// Get a ChannelController instance
	channelDao := channeldao.NewChannelDao()
	channelService := channelservice.NewChannelService(channelDao)
	channelController := channel.NewChannelController(channelService)
	apiV1Group.GET("/channels", channelController.FindChannels)
	apiV1Group.GET("/channels/:id", channelController.GetChannel)

	// Get a DeviceTypeController instance
	deviceTypeDao := devicetypedao.NewDeviceTypeDao()
	deviceTypeService := devicetypeservice.NewDeviceTypeService(deviceTypeDao)
	deviceTypeController := devicetypecontroller.NewDeviceTypeController(deviceTypeService)
	apiV1Group.GET("/deviceTypes", deviceTypeController.FindDeviceTypes)
	apiV1Group.GET("/deviceTypes/:id", deviceTypeController.GetDeviceType)

	// Get a DeviceController instance
	deviceDao := devicedao.NewDeviceDao()
	deviceService := deviceservice.NewDeviceService(deviceDao)
	deviceController := device.NewDeviceController(deviceService)
	apiV1Group.GET("/devices", deviceController.FindDevices)
	apiV1Group.GET("/devices/:id", deviceController.GetDevice)
	apiV1Group.GET("/devices/statuses", deviceController.FindDeviceStatuses)
	apiV1Group.PUT("/devices", deviceController.UpdateDevice)
	apiV1Group.DELETE("/devices/:id", deviceController.DeleteDevice)

	// Get a UserController instance
	userDao := userdao.NewUserDao()
	userService := userservice.NewUserService(userDao)
	userController := user.NewUserController(userService)
	apiV1Group.POST("/users", userController.CreateUser)

	// TODO : protect endpoints
	// Access token endpoint
	apiV1Group.POST("/login", login)
	// Refresh token endpoint
	// oauth/refresh

	// TODO : remove
	e.POST("/test", test)
	e.GET("/test", test)

	// e.Logger.Fatal(e.Start(":17172"))
	e.Server.Addr = ":17172"
	graceful.ListenAndServe(e.Server, 5*time.Second)
	// e.Logger.Fatal(gracehttp.Serve(e.Server))
}

func login(c echo.Context) error {

	log.Println("login request")

	username := c.FormValue("username")
	password := c.FormValue("password")
	log.Println(username, password)

	if username == "" || password == "" {
		log.Println("Cannot get request parameters")
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	// TODO : password policy
	// https://github.com/go-validator/validator
	userDao := userdao.NewUserDao()
	userService := userservice.NewUserService(userDao)
	user, err := userService.GetUserByUsername(username)

	if err != nil {
		log.Println("Cannot get user")
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}
	log.Println("user: ", user)

	valid, err := security.VerifyPassword(user.Password, password)

	if err != nil {
		log.Println("Cannot validate password")
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Something went wrong"})
	}

	if !valid {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid username/password"})
	}

	// create jwt token
	token, err := createJwtToken(user.ID)
	if err != nil {
		log.Println("Error Creating JWT token", err)
		return c.JSON(http.StatusInternalServerError, "something went wrong")
	}

	// Session cookie
	cookie := getSessionCookie(token)
	c.SetCookie(cookie)

	// sess, _ := session.Get("secret", c)
	// sess.Values["foo"] = "bar"
	// sess.Save(c.Request(), c.Response())
	// log.Println(sess)

	return c.JSON(http.StatusOK, map[string]string{
		"message":      "You were logged in!",
		"access_token": token,
	})
}

// TODO : RS512
func createJwtToken(userID int) (string, error) {
	claims := jwt.StandardClaims{
		Subject: strconv.Itoa(userID),
		// Id:        "main_user_id",
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err := rawToken.SignedString([]byte("mySecret"))
	if err != nil {
		return "", err
	}
	log.Println(token)

	return token, nil
}

func getSessionCookie(token string) *http.Cookie {
	cookie := &http.Cookie{}
	cookie.Path = "/"
	cookie.Name = "SESSIONID"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	// TODO
	// cookie.Secure = true
	log.Println("cookie: ", cookie)
	return cookie
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome!")
}

func test(c echo.Context) error {
	log.Println("test")

	// cookie, err := c.Cookie("SESSIONID")
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println("cookie: ", cookie)

	cookie := &http.Cookie{}
	// cookie := new(http.Cookie)
	// cookie.Domain = "localhost"
	cookie.Path = "/"
	cookie.Name = "test"
	// cookie.Value = token
	cookie.Value = "test"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	// cookie.HttpOnly = true
	c.SetCookie(cookie)
	// http.SetCookie(c.Response().Writer, cookie)
	// log.Println(cookie)

	return c.JSON(http.StatusOK, map[string]string{
		"message": "OK!",
	})
	// return c.String(http.StatusOK, "write a cookie")
}
