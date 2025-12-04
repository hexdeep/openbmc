package handler

import (
	"github.com/hexdeep/openbmc/backend/file"
	"github.com/hexdeep/openbmc/backend/log"
	"github.com/hexdeep/openbmc/backend/proc"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GetRouter(h *Handler) *echo.Echo {

	r := echo.New()
	r.HTTPErrorHandler = HTTPErrorHandler
	r.Use(middleware.CORS())
	r.Use(h.LoggerMiddleWare)

	api := r.Group("/api")

	api.POST("/login", h.Login)

	pro := api.Group("")
	// pro.Use(h.WithAuthToken)
	pro.GET("/usage", WithSSE(h, Usage))
	pro.GET("/fan-speeds", h.ListFanSpeed)
	pro.GET("/optical-port", h.ListOpticalPort)
	pro.Static("/file", "/data/file")

	pro.GET("/powered-slot", WithSSE(h, ListPoweredSlot))
	pro.POST("/bmc/opentty", h.BMCOpenTTY)
	pro.POST("/bmc/closetty", h.BMCCloseTTY)
	pro.POST("/switch/opentty", h.SwitchOpenTTY)
	pro.POST("/switch/closetty", h.SwitchCloseTTY)
	pro.POST("/slot/:id/opentty", h.SlotOpenTTY)
	pro.POST("/slot/:id/closetty", h.SlotCloseTTY)
	pro.POST("/slot/:id/flash", h.Flash)

	fileRouter := pro.Group("/file")
	fileHandler := file.NewHandler(file.NewRepository("/data"))
	fileRouter.GET("", WithBind(fileHandler.ListFolder))
	fileRouter.POST("", WithBind(fileHandler.CreateFolder))
	fileRouter.POST("", WithBind(fileHandler.UploadFile))
	fileRouter.DELETE("", WithBind(fileHandler.Delete))

	logRouter := pro.Group("/log")
	logHandler := log.NewHandler(log.NewMySQL(nil, nil))
	logRouter.GET("", WithBind(logHandler.List))
	logRouter.DELETE("", WithBind(logHandler.Delete))

	slotRouter := pro.Group("/slot")
	slotHandler := NewSlotHandler(proc.NewSlot())
	slotRouter.GET("/power-status", slotHandler.PowerStatus)
	slotRouter.POST("/:id/power-on", slotHandler.PowerOn)
	slotRouter.POST("/:id/power-off", slotHandler.PowerOff)

	powerRouter := pro.Group("/power")
	powerHandler := NewPowerHandler(proc.NewPower())
	powerRouter.GET("/power-status", powerHandler.PowerStatus)
	powerRouter.POST("/:id/power-on", powerHandler.PowerOn)
	powerRouter.POST("/:id/power-off", powerHandler.PowerOff)

	r.GET("/*", HandleFrontend)

	return r
}
