package handler

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Flash(c echo.Context) error {

	id := c.Param("id")

	if err := h.Proc.SlotPwrOff.Do(id); err != nil {
		return fmt.Errorf("failed to power off: %w\n", err)
	}
	time.Sleep(300 * time.Millisecond)

	if err := h.Proc.SlotBootOn.Do(id); err != nil {
		return fmt.Errorf("failed to boot on: %w\n", err)
	}
	time.Sleep(300 * time.Millisecond)

	if err := h.Proc.SlotPwrOn.Do(id); err != nil {
		return fmt.Errorf("failed to power on: %w\n", err)
	}
	time.Sleep(300 * time.Millisecond)

	if err := h.Proc.SlotSerial.Flash(c.Request().Context()); err != nil {
		return fmt.Errorf("failed to flash: %w\n", err)
	}
	time.Sleep(300 * time.Millisecond)

	if err := h.Proc.SlotBootOff.Do(id); err != nil {
		return fmt.Errorf("failed to boot off: %w\n", err)
	}
	time.Sleep(300 * time.Millisecond)

	if err := h.Proc.SlotPwrOff.Do(id); err != nil {
		return fmt.Errorf("failed to power off: %w\n", err)
	}
	time.Sleep(300 * time.Millisecond)

	if err := h.Proc.SlotPwrOn.Do(id); err != nil {
		return fmt.Errorf("failed to power on: %w\n", err)
	}

	return c.JSON(200, Res("finished", nil))
}
