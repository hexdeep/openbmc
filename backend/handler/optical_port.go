package handler

import (
	"time"

	"github.com/labstack/echo/v4"
)

func (h *Handler) ListOpticalPort(c echo.Context) error {

	type OpticalPort struct {
		ID     string `json:"id"`
		Active bool   `json:"active"`
	}

	interfaces, err := h.Proc.SwitchSerial.ShowInterface(time.Second)
	if err != nil {
		return err
	}

	ports := [][]OpticalPort{
		{{ID: "50"}, {ID: "51"}},
		{{ID: "49"}, {ID: "52"}},
	}

	for _, group := range ports {
		for index, port := range group {
			group[index].Active = interfaces[port.ID]
		}
	}

	return c.JSON(200, Res("", ports))
}
