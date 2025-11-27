package handler

import (
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
	pro.POST("/slot/:id/power-on", h.SlotPowerOn)
	pro.POST("/slot/:id/power-off", h.SlotPowerOff)
	pro.POST("/clear-logs", WithBind(h, ClearLog))
	pro.GET("/fan-speeds", h.ListFanSpeed)
	pro.GET("/optical-port", h.ListOpticalPort)
	pro.GET("/folder", h.GetFolder)
	pro.POST("/folder", WithBind(h, CreateFolder))
	pro.POST("/file", h.UploadFile)
	pro.Static("/file", "/data/file")
	pro.POST("/delete-file", WithBind(h, DeleteFile))
	pro.GET("/logs", WithBind(h, ListLog))

	pro.GET("/main-power", h.ListMainPower)
	pro.POST("/main-power/:id/on", h.MainPowerOn)
	pro.POST("/main-power/:id/off", h.MainPowerOff)
	pro.GET("/sub-power", h.ListSubPower)

	pro.GET("/powered-slot", WithSSE(h, ListPoweredSlot))
	pro.POST("/bmc/opentty", h.BMCOpenTTY)
	pro.POST("/bmc/closetty", h.BMCCloseTTY)
	pro.POST("/switch/opentty", h.SwitchOpenTTY)
	pro.POST("/switch/closetty", h.SwitchCloseTTY)
	pro.POST("/slot/:id/opentty", h.SlotOpenTTY)
	pro.POST("/slot/:id/closetty", h.SlotCloseTTY)

	r.GET("/*", HandleFrontend)

	return r
}
