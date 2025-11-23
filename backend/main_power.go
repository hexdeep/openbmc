package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/labstack/echo/v4"
)

func (h *Handler) MainPowerOn(c echo.Context) error {

	id := c.Param("id")

	if !IsPowerIDValid(id) {
		return c.JSON(400, Res("非法的电源标识", nil))
	}

	if err := MainPowerOn(id); err != nil {
		return err
	}

	return c.JSON(200, Res("操作成功", true))
}

func (h *Handler) MainPowerOff(c echo.Context) error {

	id := c.Param("id")

	if !IsPowerIDValid(id) {
		return c.JSON(400, Res("非法的电源标识", nil))
	}

	if err := MainPowerOff(id); err != nil {
		return err
	}

	return c.JSON(200, Res("操作成功", true))
}

type MainPowerStatus struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

var MainPowerStatusMap = map[string][]MainPowerStatus{
	"0x0022\r\n": []MainPowerStatus{
		{"3", "上方主电源", false},
		{"7", "下方主电源", false},
	},
	"0x002b\r\n": []MainPowerStatus{
		{"3", "上方主电源", true},
		{"7", "下方主电源", false},
	},
	"0x00b2\r\n": []MainPowerStatus{
		{"3", "上方主电源", false},
		{"7", "下方主电源", true},
	},
	"0x00bb\r\n": []MainPowerStatus{
		{"3", "上方主电源", true},
		{"7", "下方主电源", true},
	},
}

func (h *Handler) ListMainPower(c echo.Context) error {

	data, err := os.ReadFile(fmt.Sprintf("/proc/hexdeep_main_pwr/pwr_status"))
	if err != nil {
		return err
	}

	return c.JSON(200, Res("", MainPowerStatusMap[string(data)]))
}

func IsPowerIDValid(id string) bool {
	return slices.Contains([]string{"3", "7"}, id)
}

func MainPowerOn(id string) error {
	return os.WriteFile("/proc/hexdeep_main_pwr/pwr_on", []byte(id), 0)
}

func MainPowerOff(id string) error {
	return os.WriteFile("/proc/hexdeep_main_pwr/pwr_off", []byte(id), 0)
}
