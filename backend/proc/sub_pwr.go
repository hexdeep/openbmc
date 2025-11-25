package proc

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type SubPwrStatus struct {
	Mu sync.Mutex
}

type SubPowerStatus [48]bool

func (s *SubPowerStatus) GetPoweredSlots() []string {
	result := make([]string, 0)
	for port, active := range s {
		if active {
			result = append(result, strconv.Itoa(port+1))
		}
	}
	return result
}

func (p *SubPwrStatus) Get() (*SubPowerStatus, error) {
	var result SubPowerStatus

	p.Mu.Lock()
	dataBytes, err := os.ReadFile("/proc/hexdeep_sub_pwr/pwr_status")
	if err != nil {
		return nil, err
	}
	p.Mu.Unlock()

	// e.g. "0xffff,0xaaab,0xaaaa,0xaaaa,0xaaaa,0xeaaa\n"
	content := strings.TrimSpace(string(dataBytes))
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
		for bit := range 8 {
			bitPos := bit * 2
			result[idx] = (value & (1 << bitPos)) != 0
			idx++
		}
	}

	return &result, nil
}

type SlotPwrOn struct {
	Mu sync.Mutex
}

func (s *SlotPwrOn) Do(id string) error {

	if !IsSlotIDValid(id) {
		return ErrInvalidID
	}

	s.Mu.Lock()
	defer s.Mu.Unlock()
	return os.WriteFile("/proc/hexdeep_sub_pwr/pwr_on", []byte(id), 0)
}

type SlotPwrOff struct {
	Mu sync.Mutex
}

func (s *SlotPwrOff) Do(id string) error {

	if !IsSlotIDValid(id) {
		return ErrInvalidID
	}

	s.Mu.Lock()
	defer s.Mu.Unlock()
	return os.WriteFile("/proc/hexdeep_sub_pwr/pwr_off", []byte(id), 0)
}

func IsSlotIDValid(id string) bool {

	slot, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	if slot < 1 || slot > 48 {
		return false
	}

	return true
}

func SlotIDToTTY(id string) (string, error) {

	slot, err := strconv.Atoi(id)
	if err != nil {
		return "", err
	}

	usbMap := []int{7, 4, 5, 1, 6, 3, 2, 0}

	group := (slot - 1) / 8
	pos := (slot - 1) % 8

	usb := group*8 + usbMap[pos]

	return "/dev/ttyCH9344USB" + strconv.Itoa(usb), nil
}
