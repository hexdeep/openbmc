package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type InterfaceInfo struct {
	Port   int    `json:"port"`
	Status string `json:"status"`
}

func (h *Handler) ShowInterface(c echo.Context) error {

	type SOMStatus struct {
		ID     uint `json:"id"`
		Status bool `json:"status"`
	}

	result, err := h.SerialCommand("show interface\n")
	if err != nil {
		return err
	}

	interfaces, err := ParseShowInterface(result)
	if err != nil {
		return err
	}

	return c.JSON(200, Res("", interfaces))
}

func ParseShowInterface(output string) ([]InterfaceInfo, error) {
	lines := strings.Split(output, "\n")
	var result []InterfaceInfo

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		// 跳过 prompt、标题等
		if strings.HasPrefix(line, "Switch_config") ||
			strings.HasPrefix(line, "port ") ||
			strings.HasPrefix(line, "HexDeep") {
			continue
		}

		// 按空白字符分割（自动处理多空格）
		fields := strings.Fields(line)
		if len(fields) < 5 {
			continue
		}

		// port 应为数字
		port, err := strconv.Atoi(fields[0])
		if err != nil {
			continue
		}

		info := InterfaceInfo{
			Port:   port,
			Status: fields[2],
		}

		result = append(result, info)
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

	sample := []Interface{
		{
			ID:          1,
			Active:      true,
			Drawer:      1,
			DiskUsed:    128_000_000,
			DiskTotal:   512_000_000,
			CPUUsage:    23,
			Temperature: 41.2,
			IP:          "10.0.0.1",
		},
		{
			ID:          2,
			Active:      true,
			Drawer:      1,
			DiskUsed:    980_000_000,
			DiskTotal:   1_000_000_000,
			CPUUsage:    77,
			Temperature: 56.8,
			IP:          "10.0.0.2",
		},
		{
			ID:          3,
			Active:      false,
			Drawer:      2,
			DiskUsed:    0,
			DiskTotal:   256_000_000,
			CPUUsage:    0,
			Temperature: 0.0,
			IP:          "10.0.0.3",
		},
		{
			ID:          4,
			Active:      true,
			Drawer:      2,
			DiskUsed:    730_000_000,
			DiskTotal:   750_000_000,
			CPUUsage:    65,
			Temperature: 49.5,
			IP:          "10.0.0.4",
		},
		{
			ID:          5,
			Active:      true,
			Drawer:      3,
			DiskUsed:    35_000_000,
			DiskTotal:   256_000_000,
			CPUUsage:    12,
			Temperature: 38.1,
			IP:          "10.0.0.5",
		},
		{
			ID:          6,
			Active:      true,
			Drawer:      3,
			DiskUsed:    512_000_000,
			DiskTotal:   1_000_000_000,
			CPUUsage:    91,
			Temperature: 63.3,
			IP:          "10.0.0.6",
		},
		{
			ID:          7,
			Active:      false,
			Drawer:      4,
			DiskUsed:    0,
			DiskTotal:   512_000_000,
			CPUUsage:    0,
			Temperature: 0.0,
			IP:          "10.0.0.7",
		},
		{
			ID:          8,
			Active:      true,
			Drawer:      4,
			DiskUsed:    860_000_000,
			DiskTotal:   1_024_000_000,
			CPUUsage:    42,
			Temperature: 52.6,
			IP:          "10.0.0.8",
		},
		{
			ID:          9,
			Active:      true,
			Drawer:      5,
			DiskUsed:    1_500_000_000,
			DiskTotal:   2_000_000_000,
			CPUUsage:    58,
			Temperature: 47.9,
			IP:          "10.0.0.9",
		},
		{
			ID:          10,
			Active:      true,
			Drawer:      5,
			DiskUsed:    99_000_000,
			DiskTotal:   250_000_000,
			CPUUsage:    19,
			Temperature: 36.4,
			IP:          "10.0.0.10",
		},
	}

	return c.JSON(200, Res("", sample))
}

func isBetween1And48(s string) bool {
	n, err := strconv.Atoi(s)
	if err != nil {
		return false // not a number
	}
	return n >= 1 && n <= 48
}

func (h *Handler) InterfacePowerOn(c echo.Context) error {

	id := c.Param("id")

	if !isBetween1And48(id) {
		return c.JSON(400, Res("非法的插槽编号", nil))
	}

	if err := os.WriteFile("/proc/hexdeep_sub_pwr/pwr_on", []byte(id), 0); err != nil {
		return err
	}

	return c.JSON(200, Res("执行成功", true))
}

func (h *Handler) InterfacePowerOff(c echo.Context) error {

	id := c.Param("id")

	if !isBetween1And48(id) {
		return c.JSON(400, Res("非法的插槽编号", nil))
	}

	if err := os.WriteFile("/proc/hexdeep_sub_pwr/pwr_off", []byte(id), 0); err != nil {
		return err
	}

	return c.JSON(200, Res("执行成功", true))
}

func (h *Handler) MainPowerOn(c echo.Context) error {

	id := c.Param("id")

	return nil
}
