package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

func (h *Handler) ListPoweredSlot(c echo.Context) error {

	status, err := h.Proc.SubPwrStatus.Get()
	if err != nil {
		return err
	}

	slots := status.GetPoweredSlots()

	type SlotStatus struct {
		Slot   string `json:"slot"`
		Active bool   `json:"active"`
		Mac    string `json:"mac"`
		IP     string `json:"ip"`
		Temp   string `json:"temp"`
	}

	ch := make(chan SlotStatus, len(slots))

	var wg sync.WaitGroup

	for _, slot := range slots {
		wg.Add(1)
		go func() {
			defer wg.Done()
			timeout := 100 * time.Millisecond
			mac, ip, err := h.Proc.SlotSerial.GetMacIP(slot, timeout)
			if err != nil {
				fmt.Printf("failed to get mac and ip: %v\n", err)
				mac, ip = "", ""
			}
			temp, err := h.Proc.SlotSerial.GetTemp(slot, timeout)
			if err != nil {
				fmt.Printf("failed to get temp: %v\n", err)
				temp = ""
			}
			ch <- SlotStatus{
				Slot:   slot,
				Active: h.Proc.SlotSerial.IsActive(slot, timeout),
				Mac:    mac,
				IP:     ip,
				Temp:   temp,
			}
		}()
	}

	wg.Wait()
	close(ch)

	results := make([]SlotStatus, 0, len(slots))
	for s := range ch {
		results = append(results, s)
	}

	return c.JSON(200, Res("", results))
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
