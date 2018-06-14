package main

import (
	"log"
	"net/http"
	"time"

	// "github.com/julienschmidt/httprouter"

	"github.com/RangelReale/osin"
	// "github.com/facebookgo/grace/gracehttp"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/manumura/golang-app-device/config"
	"github.com/manumura/golang-app-device/controller/channel"
	"github.com/manumura/golang-app-device/controller/device"
	"github.com/manumura/golang-app-device/controller/device-type"
	"github.com/manumura/golang-app-device/controller/user"
	"github.com/manumura/golang-app-device/dao/channel"
	"github.com/manumura/golang-app-device/dao/device"
	"github.com/manumura/golang-app-device/dao/device-type"
	"github.com/manumura/golang-app-device/dao/user"
	"github.com/manumura/golang-app-device/oauth"
	"github.com/manumura/golang-app-device/security"
	"github.com/manumura/golang-app-device/service/channel"
	"github.com/manumura/golang-app-device/service/device"
	"github.com/manumura/golang-app-device/service/device-type"
	"github.com/manumura/golang-app-device/service/user"
	"github.com/ory/osin-storage/storage/postgres"
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

	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	// TODO
	// sess := sessions.NewCookieStore([]byte("secret"))
	// sess.Options = &sessions.Options{
	// 	Path:     "/",
	// 	MaxAge:   86400 * 7,
	// 	HttpOnly: true,
	// }
	// e.Use(session.Middleware(sess))

	apiGroup := e.Group("/api")
	// this logs the server interaction
	apiGroup.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}]  ${status}  ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	e.GET("/", index)

	// Get a ChannelController instance
	channelDao := channeldao.NewChannelDao()
	channelService := channelservice.NewChannelService(channelDao)
	channelController := channel.NewChannelController(channelService)
	e.GET("/api/channels", channelController.FindChannels)
	e.GET("/api/channels/:id", channelController.GetChannel)

	// Get a DeviceTypeController instance
	deviceTypeDao := devicetypedao.NewDeviceTypeDao()
	deviceTypeService := devicetypeservice.NewDeviceTypeService(deviceTypeDao)
	deviceTypeController := devicetypecontroller.NewDeviceTypeController(deviceTypeService)
	e.GET("/api/deviceTypes", deviceTypeController.FindDeviceTypes)
	e.GET("/api/deviceTypes/:id", deviceTypeController.GetDeviceType)

	// Get a DeviceController instance
	deviceDao := devicedao.NewDeviceDao()
	deviceService := deviceservice.NewDeviceService(deviceDao)
	deviceController := device.NewDeviceController(deviceService)
	apiGroup.GET("/devices", deviceController.FindDevices)
	apiGroup.GET("/devices/:id", deviceController.GetDevice)
	apiGroup.GET("/devices/statuses", deviceController.FindDeviceStatuses)
	apiGroup.PUT("/devices", deviceController.UpdateDevice)
	apiGroup.DELETE("/devices/:id", deviceController.DeleteDevice)

	// Get a UserController instance
	userDao := userdao.NewUserDao()
	userService := userservice.NewUserService(userDao)
	userController := user.NewUserController(userService)
	e.POST("/api/users", userController.CreateUser)

	// TODO : oauth + refresh token + protect endpoints
	// Access token endpoint
	e.POST("/oauth/token", login)
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

	// https://github.com/ory/fosite#example
	log.Println("login request")

	// TODO : password policy
	// https://github.com/go-validator/validator
	// TODO : scopes ?
	// TODO : externalize config + DB storage
	sconfig := osin.NewServerConfig()
	sconfig.AllowedAuthorizeTypes = osin.AllowedAuthorizeType{osin.TOKEN}
	sconfig.AllowedAccessTypes = osin.AllowedAccessType{osin.REFRESH_TOKEN, osin.PASSWORD}
	sconfig.AllowGetAccessRequest = false
	sconfig.AllowClientSecretInParams = false

	// server := osin.NewServer(sconfig, oauth.NewInMemoryStorage())
	store := postgres.New(config.Database)

	client, err := store.GetClient("appdevicemgmtclientid")
	if err != nil {
		cl := osin.DefaultClient{
			Id:          "appdevicemgmtclientid",
			Secret:      "appdevicemgmt2018",
			RedirectUri: "http://localhost:8091/",
		}
		err = store.CreateClient(&cl)

		if err != nil {
			log.Fatalln("Cannot create oauthclient")
		}
		log.Println(cl)
	}
	log.Println(client)

	server := osin.NewServer(sconfig, store)

	resp := server.NewResponse()
	defer resp.Close()

	// get the http.Request
	r := c.Request()

	if ar := server.HandleAccessRequest(resp, r); ar != nil {
		switch ar.Type {
		case osin.REFRESH_TOKEN:
			ar.Authorized = true
		case osin.PASSWORD:
			userDao := userdao.NewUserDao()
			userService := userservice.NewUserService(userDao)
			user, err := userService.GetUserByUsername(ar.Username)

			if err != nil {
				log.Println("Cannot get user")
			}
			log.Println("user: ", user)

			valid, err := security.VerifyPassword(user.Password, ar.Password)

			if err != nil {
				log.Println("Cannot validate password")
			}

			if valid {
				ar.Authorized = true
			}
		}

		server.FinishAccessRequest(resp, r, ar)
	}

	log.Println("response: ", resp)

	if resp.IsError && resp.InternalError != nil {
		log.Println("ERROR: ", resp.InternalError)
	}
	// if !resp.IsError {
	// 	resp.Output["custom_parameter"] = 19923
	// }

	// TODO : save token+username

	// Session cookie
	setSessionCookie(c)

	// sess, _ := session.Get("secret", c)
	// sess.Values["foo"] = "bar"
	// sess.Save(c.Request(), c.Response())
	// log.Println(sess)

	// writeCookie(c)
	// cookie := &http.Cookie{}
	// cookie.Name = "username"
	// cookie.Value = "jon"
	// cookie.Expires = time.Now().Add(24 * time.Hour)
	// // cookie.HttpOnly = true
	// c.SetCookie(cookie)
	// log.Println("cookie: ", cookie)
	// log.Println(c)

	return oauth.OutputJSON(resp, c)
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome!")
}

func setSessionCookie(c echo.Context) {
	cookie := &http.Cookie{}
	cookie.Name = "SESSIONID"
	cookie.Value = "jon"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	log.Println("cookie: ", cookie)
	c.SetCookie(cookie)
}

func test(c echo.Context) error {
	log.Println("test")

	user, err := c.Cookie("username")
	if err != nil {
		log.Println(err)
	}
	log.Println("user: ", user)

	cookie := &http.Cookie{}
	// cookie.Domain = "localhost"
	cookie.Path = "/"
	cookie.Name = "username"
	cookie.Value = "jon"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	// cookie.HttpOnly = true
	// cookie.Secure = true
	c.SetCookie(cookie)
	log.Println("cookie: ", cookie)
	// log.Println(c)

	return c.JSON(http.StatusOK, `{"hello": "world"}`)
}
