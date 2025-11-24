package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type SubPower struct {
	ID     string `json:"id"`
	Active bool   `json:"active"`
}

func (h *Handler) ListSubPower(c echo.Context) error {

	status, err := GetSubPowerStatus()
	if err != nil {
		return err
	}

	subPowers := make([]SubPower, 0)
	for i, v := range status {
		subPowers = append(subPowers, SubPower{
			ID:     strconv.Itoa(i + 1),
			Active: v,
		})
	}

	return c.JSON(200, Res("", subPowers))
}

type SubPowerStatus [48]bool

func (s *SubPowerStatus) GetPoweredSlots() []Slot {
	result := make([]Slot, 0)
	for port, active := range s {
		if active {
			result = append(result, Slot(port+1))
		}
	}
	return result
}

func GetSubPowerStatus() (*SubPowerStatus, error) {
	var result SubPowerStatus

	/*
		dataBytes, err := os.ReadFile("/proc/hexdeep_sub_pwr/pwr_status")
		if err != nil {
			return nil, err
		}
	*/

	f, err := os.Open("/proc/hexdeep_sub_pwr/pwr_status")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf := make([]byte, 256)
	n, err := f.Read(buf)
	if err != nil && err != io.EOF {
		return nil, err
	}

	dataBytes := make([]byte, 256)
	f.Read(dataBytes)

	// e.g. "0xffff,0xaaab,0xaaaa,0xaaaa,0xaaaa,0xeaaa\n"
	content := strings.TrimSpace(string(buf[:n]))
	parts := strings.Split(content, ",")
	if len(parts) != 6 {
		return nil, fmt.Errorf("invalid pwr_status format: %q", content)
	}

	idx := 0 // 0..47
	for _, p := range parts {
		p = strings.TrimSpace(p)

		// strip "0x" prefix if present
		if strings.HasPrefix(p, "0x") || strings.HasPrefix(p, "0X") {
			p = p[2:]
		}

		// parse 16-bit value
		value, err := strconv.ParseUint(p, 16, 16)
		if err != nil {
			return nil, fmt.Errorf("invalid hex value %q: %w", p, err)
		}

		// 8 channels per group → bits 0,2,4,...,14
		for bit := 0; bit < 8; bit++ {
			bitPos := bit * 2
			result[idx] = (value & (1 << bitPos)) != 0
			idx++
		}
	}

	return &result, nil
}
