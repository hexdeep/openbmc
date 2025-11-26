package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) SlotPowerOn(c echo.Context) error {

	if err := h.Proc.SlotPwrOn.Do(c.Param("id")); err != nil {
		return err
	}

	return c.JSON(200, Res("执行成功", true))
}

func (h *Handler) SlotPowerOff(c echo.Context) error {

	if err := h.Proc.SlotPwrOff.Do(c.Param("id")); err != nil {
		return err
	}

	return c.JSON(200, Res("执行成功", true))
}
