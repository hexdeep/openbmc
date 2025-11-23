package main

import (
	"encoding/json"
	"log"
	"math"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/v4/cpu"
)

func Usage(h *Handler, c echo.Context, send chan<- string) {
	type Usage struct {
		CPU      float64 `json:"cpu"`
		MemTotal uint64  `json:"memTotal"`
		MemUsed  uint64  `json:"memUsed"`
	}

	ctx := c.Request().Context()

	for {
		select {
		case <-ctx.Done():
			return

		default:
			// CPU usage（1 秒平均）
			percent, err := cpu.Percent(time.Second, false)
			if err != nil || len(percent) == 0 {
				continue
			}

			// Memory usage
			vm, err := mem.VirtualMemory()
			if err != nil {
				log.Printf("failed to get memory info: %v\n", err)
				continue
			}

			usage := Usage{
				CPU:      math.Round(percent[0]*10) / 10,
				MemTotal: vm.Total,
				MemUsed:  vm.Used,
			}

			data, err := json.Marshal(usage)
			if err != nil {
				log.Printf("failed to marshal usage: %v\n", err)
				continue
			}

			select {
			case send <- string(data):
			case <-ctx.Done():
				return
			}
		}
	}
}
