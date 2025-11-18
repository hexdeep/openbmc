package main

import "github.com/labstack/echo/v4"

func GetRouter(h *Handler) *echo.Echo {

	r := echo.New()
	r.Use(h.LoggerMiddleWare)

	api := r.Group("/api")

	pub := api.Group("")
	pub.POST("/login", h.Login)

	pro := api.Group("")
	pro.Use(h.WithAuthToken)

	soms := pro.Group("/soms")
	soms.GET("", h.ListSOM)
	soms.GET("/:id", h.GetSOM)

	files := pro.Group("/files")
	files.GET("", h.ListFile)
	files.POST("", h.UploadFile)
	files.DELETE("", h.DeleteFile)

	return r
}
