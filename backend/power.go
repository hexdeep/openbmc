package main

import "github.com/labstack/echo/v4"

func (h *Handler) ListPower(c echo.Context) error {

	type Power struct {
		ID      uint   `json:"id"`
		Name    string `json:"name"`
		Powered bool   `json:"powered"`
		Running bool   `json:"running"`
	}

	return c.JSON(200, Res("", []Power{
		{
			ID:      1,
			Name:    "上方主电源",
			Powered: false,
			Running: false,
		},
		{
			ID:      2,
			Name:    "下方主电源",
			Powered: false,
			Running: false,
		},
	}))
}
