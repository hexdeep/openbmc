package main

import (
	"context"
	"strconv"
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
		Slot   int  `json:"slot"`
		Active bool `json:"active"`
	}

	ch := make(chan SlotStatus, len(slots))

	var wg sync.WaitGroup

	for _, slot := range slots {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ctx, canc := context.WithTimeout(c.Request().Context(), 1*time.Second)
			defer canc()
			ch <- SlotStatus{
				Slot:   int(slot),
				Active: slot.IsActive(ctx),
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
