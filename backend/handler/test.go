package handler

import "github.com/labstack/echo/v4"

type Powerer interface {
	PowerOn(id string) error
	PowerOff(id string) error
	PowerStatus() (map[string]bool, error)
}

type Switcher interface {
	OpticsStatus() (map[string]bool, error)
}

// --------------

func (sh *SlotHandler) PowerOn(c echo.Context) error {
	return nil
}

// ------------

type PowerHandler struct {
	Powerer
}

func (ph *PowerHandler) PowerStatus(c echo.Context) error {

	return nil
}

func (ph *PowerHandler) PowerOn(c echo.Context) error {

	return nil
}

func (ph *PowerHandler) PowerOff(c echo.Context) error {

	return nil
}
