package main

import "github.com/labstack/echo/v4"

func (h *Handler) ListOpticalPort(c echo.Context) error {

	type OpticalPort struct {
		Name      string `json:"name"`
		Connected bool   `json:"connected"`
	}

	return c.JSON(200, Res("", []OpticalPort{
		{
			Name:      "光口 1",
			Connected: true,
		},
		{
			Name:      "光口 2",
			Connected: true,
		},
		{
			Name:      "光口 3",
			Connected: false,
		},
		{
			Name:      "光口 4",
			Connected: false,
		},
	}))
}
