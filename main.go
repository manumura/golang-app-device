package main

import (
	"log"
	"net/http"

	// "github.com/julienschmidt/httprouter"

	"github.com/RangelReale/osin"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/manumura/golang-app-device/controller/channel"
	"github.com/manumura/golang-app-device/controller/device"
	"github.com/manumura/golang-app-device/controller/device-type"
	"github.com/manumura/golang-app-device/dao/channel"
	"github.com/manumura/golang-app-device/dao/device"
	"github.com/manumura/golang-app-device/dao/device-type"
	"github.com/manumura/golang-app-device/oauth"
	"github.com/manumura/golang-app-device/service/channel"
	"github.com/manumura/golang-app-device/service/device"
	"github.com/manumura/golang-app-device/service/device-type"
	// "github.com/rs/cors"
)

// TODO : DI for database
// Application starts here.
func main() {
	// r := httprouter.New()
	// r.GET("/", index)
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	e.GET("/", index)

	// Get a ChannelController instance
	channelDao := channeldao.NewChannelDao()
	channelService := channelservice.NewChannelService(channelDao)
	channelController := channel.NewChannelController(channelService)
	e.GET("/dm/api/v1/channels", channelController.FindChannels)
	e.GET("/dm/api/v1/channels/:id", channelController.GetChannel)

	// Get a DeviceTypeController instance
	deviceTypeDao := devicetypedao.NewDeviceTypeDao()
	deviceTypeService := devicetypeservice.NewDeviceTypeService(deviceTypeDao)
	deviceTypeController := devicetypecontroller.NewDeviceTypeController(deviceTypeService)
	e.GET("/dm/api/v1/deviceTypes", deviceTypeController.FindDeviceTypes)
	e.GET("/dm/api/v1/deviceTypes/:id", deviceTypeController.GetDeviceType)

	// Get a DeviceController instance
	deviceDao := devicedao.NewDeviceDao()
	deviceService := deviceservice.NewDeviceService(deviceDao)
	deviceController := device.NewDeviceController(deviceService)
	e.GET("/dm/api/v1/devices", deviceController.FindDevices)
	e.GET("/dm/api/v1/devices/:id", deviceController.GetDevice)
	e.GET("/dm/api/v1/devices/statuses", deviceController.FindDeviceStatuses)

	// TODO : oauth + protect endpoints
	// Access token endpoint
	e.POST("/oauth/token", login)

	e.POST("/test", test)
	e.GET("/test", test)

	e.Logger.Fatal(e.Start(":17172"))
}

func login(c echo.Context) error {

	log.Println("login request")

	// username := c.FormValue("username")
	// password := c.FormValue("password")
	// grantType := c.FormValue("grant_type")
	// log.Println(username, password, grantType)

	// type userData struct {
	// 	Username  string `json:"username"`
	// 	Password  string `json:"password"`
	// 	GrantType string `json:"grant_type"`
	// }

	// input := &userData{}
	// if err := c.Bind(input); err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	// }
	// log.Println(input)

	// m := echo.Map{}
	// if err := c.Bind(&m); err != nil {
	// 	return err
	// }
	// log.Println(m)

	// TODO : externalize config + DB storage
	sconfig := osin.NewServerConfig()
	sconfig.AllowedAuthorizeTypes = osin.AllowedAuthorizeType{osin.TOKEN}
	sconfig.AllowedAccessTypes = osin.AllowedAccessType{osin.REFRESH_TOKEN, osin.PASSWORD}
	sconfig.AllowGetAccessRequest = false
	sconfig.AllowClientSecretInParams = false

	server := osin.NewServer(sconfig, oauth.NewInMemoryStorage())
	resp := server.NewResponse()
	defer resp.Close()

	// get the http.Request
	r := c.Request()

	if ar := server.HandleAccessRequest(resp, r); ar != nil {
		switch ar.Type {
		case osin.REFRESH_TOKEN:
			ar.Authorized = true
		case osin.PASSWORD:
			// TODO : check for DB
			if ar.Username == "admin" && ar.Password == "thepass" {
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

	return oauth.OutputJSON(resp, c)

	// return c.JSON(http.StatusOK, `{"hello": "world"}`)
}

func test(c echo.Context) error {
	log.Println("test")
	return c.JSON(http.StatusOK, `{"hello": "world"}`)
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome!")
}
