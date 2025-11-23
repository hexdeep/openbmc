package main

import (
	"context"
	"errors"
	"os"
	"strconv"

	"go.bug.st/serial"
)

type Slot int

func NewSlot(slot int) (Slot, error) {
	if slot < 1 || slot > 48 {
		return 0, errors.New("invalid slot")
	}
	return Slot(slot), nil
}

func (s Slot) ToBytes() []byte {
	return []byte(strconv.Itoa(int(s)))
}

func (s Slot) ToTTY() string {
	usbMap := []int{7, 4, 5, 1, 6, 3, 2, 0}

	slot := int(s)
	group := (slot - 1) / 8
	pos := (slot - 1) % 8

	usb := group*8 + usbMap[pos]

	return "/dev/ttyCH9344USB" + strconv.Itoa(usb)
}

func (s Slot) PowerOn() error {
	return os.WriteFile("/proc/hexdeep_sub_pwr/pwr_on", s.ToBytes(), 0)
}

func (s Slot) PowerOff() error {
	return os.WriteFile("/proc/hexdeep_sub_pwr/pwr_off", s.ToBytes(), 0)
}

func (s Slot) BootOn() error {
	return os.WriteFile("/proc/hexdeep_sub_pwr/boot_on", s.ToBytes(), 0)
}

func (s Slot) BootOff() error {
	return os.WriteFile("/proc/hexdeep_sub_pwr/boot_off", s.ToBytes(), 0)
}

func (s Slot) IsActive(ctx context.Context) bool {
	port, err := serial.Open(s.ToTTY(), &serial.Mode{BaudRate: 1500000})
	if err != nil {
		return false
	}
	defer port.Close()

	if _, err := port.Write([]byte("\n")); err != nil {
		return false
	}

	buf := make([]byte, 256)
	resultCh := make(chan bool, 1)

	go func() {
		n, err := port.Read(buf)
		if err != nil || n == 0 {
			resultCh <- false
			return
		}
		resultCh <- true
	}()

	select {
	case <-ctx.Done():
		return false
	case ok := <-resultCh:
		return ok
	}
}
