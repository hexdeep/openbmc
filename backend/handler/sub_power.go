package handler

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) ListSubPower(c echo.Context) error {

	status, err := h.Proc.SubPwrStatus.Get()
	if err != nil {
		return err
	}

	type Slot struct {
		ID     string `json:"id"`
		Active bool   `json:"active"`
	}

	data := [][][]Slot{
		{
			{{ID: "48"}, {ID: "40"}},
			{{ID: "47"}, {ID: "39"}},
			{{ID: "46"}, {ID: "38"}},
			{{ID: "45"}, {ID: "37"}},
			{{ID: "44"}, {ID: "36"}},
			{{ID: "43"}, {ID: "35"}},
			{{ID: "42"}, {ID: "34"}},
			{{ID: "41"}, {ID: "33"}},
		},
		{
			{{ID: "32"}, {ID: "24"}, {ID: "16"}, {ID: "8"}},
			{{ID: "31"}, {ID: "23"}, {ID: "15"}, {ID: "7"}},
			{{ID: "30"}, {ID: "22"}, {ID: "14"}, {ID: "6"}},
			{{ID: "29"}, {ID: "21"}, {ID: "13"}, {ID: "5"}},
			{{ID: "28"}, {ID: "20"}, {ID: "12"}, {ID: "4"}},
			{{ID: "27"}, {ID: "19"}, {ID: "11"}, {ID: "3"}},
			{{ID: "26"}, {ID: "18"}, {ID: "10"}, {ID: "2"}},
			{{ID: "25"}, {ID: "17"}, {ID: "9"}, {ID: "1"}},
		},
	}

	for _, pane := range data {
		for _, drawer := range pane {
			for index, slot := range drawer {
				num, _ := strconv.Atoi(slot.ID)
				drawer[index].Active = status.Status[num-1]
			}
		}
	}

	return c.JSON(200, Res("", data))
}
