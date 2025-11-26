package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

func (h *Handler) ListFanSpeed(c echo.Context) error {

	type FanSpeed struct {
		ID    int `json:"id"`
		Speed int `json:"speed"`
	}

	return c.JSON(200, Res("", lo.Map([]int{30, 20, 40}, func(item int, index int) FanSpeed {
		return FanSpeed{
			ID:    index + 1,
			Speed: item,
		}
	})))
}
