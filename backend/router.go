package main

import "github.com/labstack/echo/v4"

func GetRouter(h *Handler) *echo.Echo {

	r := echo.New()

	pub := r.Group("")
	pub.POST("/login", h.Login)

	pro := r.Group("")
	pro.GET("/api/files", h.ListFile)

	return r
}
