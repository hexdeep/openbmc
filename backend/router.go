package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GetRouter(h *Handler) *echo.Echo {

	r := echo.New()
	r.HTTPErrorHandler = HTTPErrorHandler
	r.Use(middleware.CORS())
	r.Use(h.LoggerMiddleWare)

	pub := r.Group("")
	pub.POST("/login", h.Login)

	pro := r.Group("")
	// pro.Use(h.WithAuthToken)

	pro.GET("/som-statuses", h.ListSOMStatus)
	pro.GET("/soms", h.ListSOM)
	pro.GET("/fan-speeds", h.ListFanSpeed)
	pro.GET("/powered-interfaces", h.ListPoweredInterface)
	pro.GET("/powers", h.ListPower)
	pro.GET("/optical-ports", h.ListOpticalPort)
	pro.GET("/folder", h.GetFolder)
	pro.POST("/folder", WithBind(h, CreateFolder))
	pro.POST("/file", h.UploadFile)
	pro.Static("/file", "/data/file")
	pro.POST("/delete-file", WithBind(h, DeleteFile))
	pro.GET("/logs", WithBind(h, ListLog))

	return r
}
