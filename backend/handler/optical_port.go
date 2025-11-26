package handler

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

func (h *Handler) ListOpticalPort(c echo.Context) error {

	type OpticalPort struct {
		Port   string `json:"port"`
		Active bool   `json:"active"`
	}

	interfaces, err := h.Proc.SwitchSerial.ShowInterface(time.Second)
	if err != nil {
		return err
	}

	return c.JSON(200, Res("", lo.Map([]string{"49", "50", "51", "52"}, func(item string, index int) *OpticalPort {
		return &OpticalPort{Port: item, Active: interfaces[item]}
	})))
}
