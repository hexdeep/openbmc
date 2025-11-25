package main

import (
	"strconv"
	"sync"

	"github.com/labstack/echo/v4"
)

type SubPwrStatusProc struct {
	Mu sync.Mutex
}

type SubPower struct {
	ID     string `json:"id"`
	Active bool   `json:"active"`
}

func (h *Handler) ListSubPower(c echo.Context) error {

	status, err := h.Proc.SubPwrStatus.Get()
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
