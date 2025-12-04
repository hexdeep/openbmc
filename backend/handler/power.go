package handler

import (
	"github.com/labstack/echo/v4"
)

type Powerer interface {
	PowerOn(id string) error
	PowerOff(id string) error
	PowerStatus() (map[string]bool, error)
}

type PowerHandler struct {
	Powerer Powerer
}

func NewPowerHandler(powerer Powerer) *PowerHandler {
	return &PowerHandler{Powerer: powerer}
}

func (ph *PowerHandler) PowerOn(c echo.Context) error {

	if err := ph.Powerer.PowerOn(c.Param("id")); err != nil {
		return err
	}

	return c.JSON(200, Res("操作成功", true))
}

func (ph *PowerHandler) PowerOff(c echo.Context) error {

	if err := ph.Powerer.PowerOff(c.Param("id")); err != nil {
		return err
	}

	return c.JSON(200, Res("操作成功", true))
}

func (ph *PowerHandler) PowerStatus(c echo.Context) error {

	status, err := ph.Powerer.PowerStatus()
	if err != nil {
		return err
	}

	return c.JSON(200, Res("", status))
}
