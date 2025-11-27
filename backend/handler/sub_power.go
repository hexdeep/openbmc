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

	type Drawer struct {
		Name  string `json:"name"`
		Slots []Slot `json:"slots"`
	}

	type Pane struct {
		Name    string   `json:"name"`
		Drawers []Drawer `json:"drawers"`
	}

	data := []Pane{
		{
			Name: "二号板",
			Drawers: []Drawer{
				{Name: "16", Slots: []Slot{{ID: "48"}, {ID: "40"}}},
				{Name: "15", Slots: []Slot{{ID: "47"}, {ID: "39"}}},
				{Name: "14", Slots: []Slot{{ID: "46"}, {ID: "38"}}},
				{Name: "13", Slots: []Slot{{ID: "45"}, {ID: "37"}}},
				{Name: "12", Slots: []Slot{{ID: "44"}, {ID: "36"}}},
				{Name: "11", Slots: []Slot{{ID: "43"}, {ID: "35"}}},
				{Name: "10", Slots: []Slot{{ID: "42"}, {ID: "34"}}},
				{Name: "9", Slots: []Slot{{ID: "41"}, {ID: "33"}}},
			},
		},
		{
			Name: "一号板",
			Drawers: []Drawer{
				{Name: "8", Slots: []Slot{{ID: "32"}, {ID: "24"}, {ID: "16"}, {ID: "8"}}},
				{Name: "7", Slots: []Slot{{ID: "31"}, {ID: "23"}, {ID: "15"}, {ID: "7"}}},
				{Name: "6", Slots: []Slot{{ID: "30"}, {ID: "22"}, {ID: "14"}, {ID: "6"}}},
				{Name: "5", Slots: []Slot{{ID: "29"}, {ID: "21"}, {ID: "13"}, {ID: "5"}}},
				{Name: "4", Slots: []Slot{{ID: "28"}, {ID: "20"}, {ID: "12"}, {ID: "4"}}},
				{Name: "3", Slots: []Slot{{ID: "27"}, {ID: "19"}, {ID: "11"}, {ID: "3"}}},
				{Name: "2", Slots: []Slot{{ID: "26"}, {ID: "18"}, {ID: "10"}, {ID: "2"}}},
				{Name: "1", Slots: []Slot{{ID: "25"}, {ID: "17"}, {ID: "9"}, {ID: "1"}}},
			},
		},
	}

	for pi := range data {
		for di := range data[pi].Drawers {
			for si := range data[pi].Drawers[di].Slots {
				slot := &data[pi].Drawers[di].Slots[si]
				num, _ := strconv.Atoi(slot.ID)
				slot.Active = status.Status[num-1]
			}
		}
	}

	return c.JSON(200, Res("", data))
}
