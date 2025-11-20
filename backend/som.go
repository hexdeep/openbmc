package main

import "github.com/labstack/echo/v4"

func (h *Handler) ListSOMStatus(c echo.Context) error {

	return nil
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
