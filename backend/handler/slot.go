package handler

import (
	"context"
	"fmt"
	"slices"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

type Sloter interface {
	PowerOn(id string) error
	PowerOff(id string) error
	BootOn(id string) error
	BootOff(id string) error
	PowerStatus() (map[string]bool, error)
	IsActive(id string, ctx context.Context) bool
	GetMAC(id string, ctx context.Context) (string, error)
	GetIP(id string, ctx context.Context) (string, error)
	GetTemp(id string, ctx context.Context) (string, error)
	GetMem(id string, ctx context.Context) (int, int, error)
}

type SlotHandler struct {
	Sloter Sloter
}

func NewSlotHandler(sloter Sloter) *SlotHandler {
	return &SlotHandler{Sloter: sloter}
}

func (sh *SlotHandler) PowerOn(c echo.Context) error {

	if err := sh.Sloter.PowerOn(c.Param("id")); err != nil {
		return err
	}

	return c.JSON(200, Res("执行成功", true))
}

func (sh *SlotHandler) PowerOff(c echo.Context) error {

	if err := sh.Sloter.PowerOff(c.Param("id")); err != nil {
		return err
	}

	return c.JSON(200, Res("执行成功", true))
}

func (sh *SlotHandler) PowerStatus(c echo.Context) error {

	status, err := sh.Sloter.PowerStatus()
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
				slot.Active = status[slot.ID]
			}
		}
	}

	return c.JSON(200, Res("", data))
}

func ListPoweredSlot(h *Handler, c echo.Context, send chan<- any) {

	type SlotStatus struct {
		Slot      string  `json:"slot"`
		Port      string  `json:"port"`
		Active    bool    `json:"active"`
		TTYActive bool    `json:"ttyActive"`
		Mac       string  `json:"mac"`
		IP        string  `json:"ip"`
		Temp      string  `json:"temp"`
		MemUsed   int     `json:"memUsed"`
		MemTotal  int     `json:"memTotal"`
		UpTime    float64 `json:"uptime"`
		Load      float64 `json:"load"`
	}

	ctx := c.Request().Context()

	for {
		select {
		case <-ctx.Done():
			return
		default:

			status, err := h.Proc.SubPwrStatus.Get()
			if err != nil {
				continue
			}

			slots := status.GetPoweredSlots()

			ch := make(chan SlotStatus, len(slots))

			var wg sync.WaitGroup

			for _, slot := range slots {
				wg.Add(1)
				go func() {
					defer wg.Done()
					timeout := 100 * time.Millisecond
					mac, ip, _ := h.Proc.SlotSerial.GetMacIP(slot, timeout)
					temp, _ := h.Proc.SlotSerial.GetTemp(slot, timeout)
					memUsed, memTotal, _ := h.Proc.SlotSerial.GetMem(slot, timeout)
					item, _ := h.Proc.SlotSerial.GetItem(slot)
					var ttyActive bool
					if item.TTY != nil {
						ttyActive = true
					}
					port, _ := h.Proc.SlotSerial.GetPort(slot)
					uptime, _ := h.Proc.SlotSerial.GetUpTime(slot, timeout)
					load, _ := h.Proc.SlotSerial.GetLoad(slot, timeout)
					ch <- SlotStatus{
						Slot:      slot,
						TTYActive: ttyActive,
						Port:      port,
						Active:    h.Proc.SlotSerial.IsActive(slot, timeout),
						Mac:       mac,
						IP:        ip,
						Temp:      temp,
						MemUsed:   memUsed,
						MemTotal:  memTotal,
						UpTime:    uptime,
						Load:      load,
					}
				}()
			}

			wg.Wait()
			close(ch)

			results := make([]SlotStatus, 0, len(slots))
			for s := range ch {
				results = append(results, s)
			}

			slices.SortFunc(results, func(a, b SlotStatus) int {
				aid, _ := strconv.Atoi(a.Slot)
				bid, _ := strconv.Atoi(b.Slot)
				return aid - bid
			})

			select {
			case send <- results:
			case <-ctx.Done():
				return
			}
		}
		time.Sleep(time.Second)
	}
}

func (h *Handler) SlotOpenTTY(c echo.Context) error {

	if err := h.Proc.SlotSerial.OpenTTY(c.Param("id")); err != nil {
		return fmt.Errorf("failed to open tty: %w\n", err)
	}

	return c.JSON(200, Res("终端开启成功", true))
}

func (h *Handler) SlotCloseTTY(c echo.Context) error {

	if err := h.Proc.SlotSerial.CloseTTY(c.Param("id")); err != nil {
		return fmt.Errorf("failed to close tty: %w\n", err)
	}

	return c.JSON(200, Res("终端关闭成功", true))
}

/*
func (h *Handler) FlushSlot(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, Res("插槽标识不合法", nil))
	}

	slot, err := NewSlot(id)
	if err != nil {
		return c.JSON(400, Res("插槽超过范围", nil))
	}

	if err := slot.PowerOff(); err != nil {
		return err
	}

	if err := slot.BootOn(); err != nil {
		return err
	}
	time.Sleep(100 * time.Millisecond)

	if err := slot.PowerOn(); err != nil {
		return err
	}
	time.Sleep(100 * time.Millisecond)

	if err := slot.BootOff(); err != nil {
		return err
	}
	time.Sleep(100 * time.Millisecond)

	return c.JSON(200, Res("刷机模式进入成功", nil))

}
*/
