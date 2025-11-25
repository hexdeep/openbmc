package main

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) MainPowerOn(c echo.Context) error {

	if err := h.Proc.MainPwrOn.Do(c.Param("id")); err != nil {
		return err
	}

	return c.JSON(200, Res("操作成功", true))
}

func (h *Handler) MainPowerOff(c echo.Context) error {

	if err := h.Proc.MainPwrOff.Do(c.Param("id")); err != nil {
		return err
	}

	return c.JSON(200, Res("操作成功", true))
}

func (h *Handler) ListMainPower(c echo.Context) error {

	status, err := h.Proc.MainPwrStatus.Get()
	if err != nil {
		return err
	}

	return c.JSON(200, Res("", status))
}
