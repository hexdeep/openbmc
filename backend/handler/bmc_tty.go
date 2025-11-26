package handler

import "github.com/labstack/echo/v4"

func (h *Handler) BMCOpenTTY(c echo.Context) error {

	if err := h.Proc.BMCSerial.OpenTTY(); err != nil {
		return err
	}

	return c.JSON(200, Res("串口开启成功", true))
}

func (h *Handler) BMCCloseTTY(c echo.Context) error {

	if err := h.Proc.BMCSerial.CloseTTY(); err != nil {
		return err
	}

	return c.JSON(200, Res("串口关闭成功", true))
}
