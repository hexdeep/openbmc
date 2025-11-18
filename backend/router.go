package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GetRouter(h *Handler) *echo.Echo {

	r := echo.New()
	r.Use(middleware.CORS())
	r.Use(h.LoggerMiddleWare)

	pub := r.Group("")
	pub.POST("/login", h.Login)

	pro := r.Group("")
	pro.Use(h.WithAuthToken)

	soms := pro.Group("/soms")
	soms.GET("", h.ListSOM)
	soms.GET("/:id", h.GetSOM)

	files := pro.Group("/files")
	files.GET("", h.ListFile)
	files.POST("", h.UploadFile)
	files.DELETE("/:id", h.DeleteFile)

	return r
}
