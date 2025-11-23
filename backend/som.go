package main

import (
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func (h *Handler) SerialShowInterface() (map[string]string, error) {

	rawResult, err := h.SerialCommand("show interface\n")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(rawResult, "\n")
	result := make(map[string]string)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "Switch_config") ||
			strings.HasPrefix(line, "port ") ||
			strings.HasPrefix(line, "HexDeep") {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) < 5 {
			continue
		}

		result[fields[0]] = fields[2]
	}

	return result, nil
}

func (h *Handler) ListSOM(c echo.Context) error {

	return nil
}

func (h *Handler) ListPoweredInterface(c echo.Context) error {

	type Interface struct {
		ID          uint    `json:"id"`
		Active      bool    `json:"active"`
		Drawer      uint    `json:"drawer"`
		DiskUsed    uint    `json:"diskUsed"`
		DiskTotal   uint    `json:"diskTotal"`
		CPUUsage    uint    `json:"cpuUsage"`
		Temperature float64 `json:"temperature"`
		IP          string  `json:"ip"`
	}

	return c.JSON(200, Res("", nil))
}

func isBetween1And48(s string) bool {
	n, err := strconv.Atoi(s)
	if err != nil {
		return false // not a number
	}
	return n >= 1 && n <= 48
}

func (h *Handler) SlotPowerOn(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, Res("插槽编号不合法", nil))
	}

	slot, err := NewSlot(id)
	if err != nil {
		return c.JSON(400, Res("插槽编号超出范围", nil))
	}

	if err := slot.PowerOn(); err != nil {
		return err
	}

	return c.JSON(200, Res("执行成功", true))
}

func (h *Handler) SlotPowerOff(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, Res("插槽编号不合法", nil))
	}

	slot, err := NewSlot(id)
	if err != nil {
		return c.JSON(400, Res("插槽编号超出范围", nil))
	}

	if err := slot.PowerOff(); err != nil {
		return err
	}

	return c.JSON(200, Res("执行成功", true))
}
