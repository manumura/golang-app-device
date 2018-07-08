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
	"io"
	"crypto/rand"
	"encoding/base64"
)

// TODO : DI for database
// Application starts here.
func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// r := httprouter.New()
	// r.GET("/", index)
	e := echo.New()

	apiV1Group := e.Group("/api/v1")

	// Recover middleware
	e.Use(middleware.Recover())

	// Logs middleware : this logs the server interaction
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}]  ${status}  ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	// CORS middleware
	corsConfig := middleware.CORSConfig{
		AllowOrigins:     	[]string{"http://localhost:8091"},
		AllowMethods:     	[]string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
		AllowHeaders:		[]string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderCookie,
			echo.HeaderAuthorization,
			echo.HeaderAccessControlAllowCredentials,
			echo.HeaderSetCookie,
			echo.HeaderXCSRFToken,
		},
		AllowCredentials: 	true,
	}
	e.Use(middleware.CORSWithConfig(corsConfig))

	// CSRF middleware
	csrfConfig := middleware.CSRFConfig{
		TokenLookup: 	"header:" + echo.HeaderXCSRFToken,
		CookieName:   	"_csrf",
		CookiePath: 	"/",
		CookieHTTPOnly: false,
		CookieSecure:   false,

	}
	apiV1Group.Use(middleware.CSRFWithConfig(csrfConfig))

	// TODO : constants
	// JWT middleware
	jwtConfig := middleware.JWTConfig{
		SigningKey: []byte("mySecret"),
		SigningMethod: "HS512",
		TokenLookup: "cookie:SESSIONID",
	}
	apiV1Group.Use(middleware.JWTWithConfig(jwtConfig))

	// Index
	apiV1Group.GET("/", index)

	// Login endpoint
	e.POST("/api/v1/login", login)

	// Logout endpoint : delete cookies session + csrf
	e.POST("/api/v1/logout", logout)

	// TODO get session cookie + refresh tokens
	//e.GET("/api/v1/login", getUserSession)

	// Get a UserController instance
	userDao := userdao.NewUserDao()
	userService := userservice.NewUserService(userDao)
	userController := user.NewUserController(userService)
	e.POST("/users", userController.CreateUser)

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

	// TODO : remove test
	e.POST("/test", test)
	e.GET("/test", test)

	e.Server.Addr = ":17172"
	graceful.ListenAndServe(e.Server, 5*time.Second)
	// e.Logger.Fatal(gracehttp.Serve(e.Server))
	// e.Logger.Fatal(e.Start(":17172"))
}

func login(c echo.Context) error {

	log.Println("login request")

	type login struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	l := &login{}
	if err := c.Bind(l); err != nil {
		log.Println(err)
		return err
	}

	//username := c.FormValue("username")
	//password := c.FormValue("password")
	username := l.Username
	password := l.Password
	log.Println(username, password)

	if username == "" || password == "" {
		log.Println("Cannot get request parameters")
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	// TODO : password policy
	// https://github.com/go-validator/validator
	userDao := userdao.NewUserDao()
	userService := userservice.NewUserService(userDao)
	u, err := userService.GetUserByUsername(username)

	if err != nil {
		log.Println("Cannot get user")
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}
	log.Println("user: ", u)

	valid, err := security.VerifyPassword(u.Password, password)

	if err != nil {
		log.Println("Cannot validate password")
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Something went wrong"})
	}

	if !valid {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid username/password"})
	}

	// create csrf token
	csrfToken, err := createCsrfToken(32)
	if err != nil {
		log.Println("Error Creating CSRF token", err)
		return c.JSON(http.StatusInternalServerError, "Something went wrong")
	}

	// CSRF cookie
	csrfCookie := setCsrfCookie(csrfToken, time.Now().Add(24 * time.Hour))
	c.SetCookie(csrfCookie)

	// create jwt token
	jwtToken, err := createJwtToken(u.ID)
	if err != nil {
		log.Println("Error Creating JWT token", err)
		return c.JSON(http.StatusInternalServerError, "Something went wrong")
	}

	// Session cookie
	jwtCookie := setSessionCookie(jwtToken, time.Now().Add(24 * time.Hour))
	c.SetCookie(jwtCookie)

	//return c.JSON(http.StatusOK, map[string]string{
	//	"message":      "You were logged in!",
	//	"access_token": token,
	//})
	return c.JSON(http.StatusOK, u)
}

func createCsrfToken(tokenLength int) (string, error) {
	buffer := make([]byte, tokenLength)

	if _, err := io.ReadFull(rand.Reader, buffer); err != nil {
		return "", err
	}

	token := base64.StdEncoding.EncodeToString(buffer)
	//log.Println("CSRFF= ", token)

	return token[:tokenLength], nil
}

func setCsrfCookie(token string, expiresAt time.Time) *http.Cookie {
	cookie := &http.Cookie{}
	cookie.Path = "/"
	cookie.Name = "_csrf"
	cookie.Value = token
	// time.Now().Add(24 * time.Hour)
	cookie.Expires = expiresAt
	cookie.HttpOnly = false
	log.Println("csrf cookie: ", cookie)
	return cookie
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

func setSessionCookie(token string, expiresAt time.Time) *http.Cookie {
	cookie := &http.Cookie{}
	cookie.Path = "/"
	cookie.Name = "SESSIONID"
	cookie.Value = token
	// time.Now().Add(24 * time.Hour)
	cookie.Expires = expiresAt
	cookie.HttpOnly = true
	// cookie.Secure = true
	log.Println("jwt cookie: ", cookie)
	return cookie
}

func logout(c echo.Context) error {

	log.Println("logout request")

	// CSRF cookie
	csrfCookie := setCsrfCookie("", time.Unix(0, 0))
	c.SetCookie(csrfCookie)

	// Session cookie
	jwtCookie := setSessionCookie("", time.Unix(0, 0))
	c.SetCookie(jwtCookie)

	return c.JSON(http.StatusOK, map[string]string{
		"message":      "You were logged out!",
	})
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
