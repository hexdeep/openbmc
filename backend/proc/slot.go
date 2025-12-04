package proc

import (
	"context"
	"fmt"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"go.bug.st/serial"
)

type Slot struct {
	PowerOnMu     sync.Mutex
	PowerOffMu    sync.Mutex
	PowerStatusMu sync.Mutex
	BootOnMu      sync.Mutex
	BootOffMu     sync.Mutex
	Items         [48]SlotItem
}

type SlotItem struct {
	SerialMu sync.Mutex
	TTY      *os.Process
}

func NewSlot() *Slot {
	return &Slot{}
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

func (s *Slot) PowerOn(id string) error {

	if !IsSlotIDValid(id) {
		return ErrInvalidID
	}

	s.PowerOnMu.Lock()
	time.Sleep(100 * time.Millisecond)
	err := os.WriteFile("/proc/hexdeep_sub_pwr/pwr_on", []byte(id), 0)
	time.Sleep(100 * time.Millisecond)
	s.PowerOnMu.Unlock()
	return err
}

func (s *Slot) PowerOff(id string) error {

	if !IsSlotIDValid(id) {
		return ErrInvalidID
	}

	s.PowerOffMu.Lock()
	time.Sleep(100 * time.Millisecond)
	err := os.WriteFile("/proc/hexdeep_sub_pwr/pwr_off", []byte(id), 0)
	time.Sleep(100 * time.Millisecond)
	s.PowerOffMu.Unlock()
	return err
}

func (s *Slot) GetPort(id string) (string, error) {

	slot, err := strconv.Atoi(id)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(slot + 7500), nil
}

func (s *Slot) GetItem(id string) (*SlotItem, error) {

	slot, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("failed to convert id to int: %w\n", err)
	}

	if slot < 1 || slot > 48 {
		return nil, ErrInvalidID
	}

	return &s.Items[slot-1], nil
}

type SlotSerialItem struct {
	Mu  sync.Mutex
	TTY *os.Process
}

func (s *Slot) IsActive(id string, ctx context.Context) bool {

	ttyId, err := SlotIDToTTY(id)
	if err != nil {
		return false
	}

	item, err := s.GetItem(id)
	if err != nil {
		return false
	}

	if !item.SerialMu.TryLock() {
		return false
	}
	defer item.SerialMu.Unlock()

	port, err := serial.Open(ttyId, &serial.Mode{BaudRate: 1500000})
	if err != nil {
		return false
	}
	defer port.Close()

	if _, err := port.Write([]byte("\n")); err != nil {
		return false
	}

	ch := make(chan []byte, 1)

	go func() {
		buf := make([]byte, 10)
		port.Read(buf)
		ch <- buf
	}()

	select {
	case <-ctx.Done():
		return false
	case <-ch:
		return true
	}
}

func (s *Slot) GetMacIP(id string, ctx context.Context) (string, string, error) {

	ttyId, err := SlotIDToTTY(id)
	if err != nil {
		return "", "", err
	}

	item, err := s.GetItem(id)
	if err != nil {
		return "", "", err
	}

	if !item.SerialMu.TryLock() {
		return "", "", nil
	}
	data, err := SerialCommand(&serial.Mode{BaudRate: 1500000}, ttyId, ctx, "ifconfig\n")
	item.SerialMu.Unlock()
	if err != nil {
		return "", "", err
	}

	lines := strings.Split(data, "\n")
	target := -1
	for index, line := range lines {
		if strings.HasPrefix(line, "eth0") {
			target = index
		}
	}
	if target == -1 {
		return "", "", nil
	}

	return Idx(strings.Fields(lines[target]), 4), Idx(strings.Split(Idx(strings.Fields(Idx(lines, target+1)), 1), ":"), 1), nil
}

func (s *Slot) GetIP(id string, ctx context.Context) (string, error) {
	_, ip, err := s.GetMacIP(id, ctx)
	return ip, err
}

func (s *Slot) GetMAC(id string, ctx context.Context) (string, error) {
	mac, _, err := s.GetMacIP(id, ctx)
	return mac, err
}

func (s *Slot) GetTemp(id string, ctx context.Context) (string, error) {

	ttyId, err := SlotIDToTTY(id)
	if err != nil {
		return "", nil
	}

	item, err := s.GetItem(id)
	if err != nil {
		return "", fmt.Errorf("failed to get slot temperature: %w\n", err)
	}

	if !item.SerialMu.TryLock() {
		return "", nil
	}
	data, err := SerialCommand(&serial.Mode{BaudRate: 1500000}, ttyId, ctx, "cat /sys/class/thermal/thermal_zone0/temp\n")
	item.SerialMu.Unlock()
	if err != nil {
		return "", err
	}

	n := len(data)
	if n > 3 {
		return data[:len(data)-3], nil
	} else {
		return "", nil
	}
}

func (s *Slot) GetMem(id string, ctx context.Context) (int, int, error) {

	ttyId, err := SlotIDToTTY(id)
	if err != nil {
		return 0, 0, nil
	}

	item, err := s.GetItem(id)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to get slot mem: %w\n", err)
	}

	if !item.SerialMu.TryLock() {
		return 0, 0, nil
	}
	data, err := SerialCommand(&serial.Mode{BaudRate: 1500000}, ttyId, ctx, "cat /proc/meminfo\n")
	item.SerialMu.Unlock()
	if err != nil {
		return 0, 0, err
	}

	lines := strings.Split(data, "\n")

	total, err := strconv.Atoi(Idx(strings.Fields(Idx(lines, 0)), 1) + "000")
	if err != nil {
		return 0, 0, err
	}

	free, err := strconv.Atoi(Idx(strings.Fields(Idx(lines, 1)), 1) + "000")
	if err != nil {
		return 0, 0, err
	}

	return total - free, total, nil
}

func (s *Slot) GetUpTime(id string, ctx context.Context) (float64, error) {

	ttyId, err := SlotIDToTTY(id)
	if err != nil {
		return 0, err
	}

	item, err := s.GetItem(id)
	if err != nil {
		return 0, fmt.Errorf("failed to get slot up time: %v\n", err)
	}

	if !item.SerialMu.TryLock() {
		return 0, nil
	}
	data, err := SerialCommand(&serial.Mode{BaudRate: 1500000}, ttyId, ctx, "cat /proc/uptime\n")
	item.SerialMu.Unlock()
	if err != nil {
		return 0, err
	}

	value, err := strconv.ParseFloat(Idx(strings.Fields(data), 0), 64)
	if err != nil {
		return 0, err
	}

	return value, nil
}

func (s *Slot) GetLoad(id string, ctx context.Context) (float64, error) {

	ttyId, err := SlotIDToTTY(id)
	if err != nil {
		return 0, err
	}

	item, err := s.GetItem(id)
	if err != nil {
		return 0, err
	}

	if !item.SerialMu.TryLock() {
		return 0, nil
	}
	data, err := SerialCommand(&serial.Mode{BaudRate: 1500000}, ttyId, ctx, "cat /proc/loadavg\n")
	item.SerialMu.Unlock()
	if err != nil {
		return 0, err
	}

	value, err := strconv.ParseFloat(Idx(strings.Fields(data), 0), 64)
	if err != nil {
		return 0, err
	}

	return math.Round(value*1250) / 100, nil
}

func (s *Slot) OpenTTY(id string) error {

	item, err := s.GetItem(id)
	if err != nil {
		return fmt.Errorf("failed to get slot id in open tty: %v\n", err)
	}

	tty, err := SlotIDToTTY(id)
	if err != nil {
		return fmt.Errorf("failed to convert slot id to tty :%v\n", err)
	}

	num, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	item.SerialMu.Lock()
	cmd := exec.Command(
		"./ttyd.aarch64", "-p", strconv.Itoa(7500+num), "-W", "microcom", "-s", "1500000", tty,
	)
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start slot tty: %v\n", err)
	}

	item.TTY = cmd.Process
	return nil
}

func (s *Slot) CloseTTY(id string) error {

	item, err := s.GetItem(id)
	if err != nil {
		return fmt.Errorf("failed to get slot id in close tty: %v\n", err)
	}

	if err := item.TTY.Signal(syscall.SIGTERM); err != nil {
		return fmt.Errorf("failed to close slot tty: %v\n", err)
	}

	item.SerialMu.Unlock()
	item.TTY = nil
	return nil
}

func (t *Slot) Flash(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, "./flash.sh")
	result, err := cmd.Output()
	if err != nil {
		fmt.Println(string(result))
		return err
	}
	return nil
}

func (s *Slot) GetPoweredSlots() ([]string, error) {
	data, err := s.PowerStatus()
	if err != nil {
		return nil, err
	}
	result := make([]string, 0)
	for port, active := range data {
		if active {
			result = append(result, port)
		}
	}
	return result, nil
}

func (p *Slot) PowerStatus() (map[string]bool, error) {

	var result map[string]bool

	p.PowerStatusMu.Lock()
	dataBytes, err := os.ReadFile("/proc/hexdeep_sub_pwr/pwr_status")
	if err != nil {
		return nil, err
	}
	p.PowerStatusMu.Unlock()

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
			result[strconv.Itoa(idx)] = (value & (1 << bitPos)) != 0
			idx++
		}
	}

	return result, nil
}

func (s *Slot) BootOn(id string) error {

	if !IsSlotIDValid(id) {
		return ErrInvalidID
	}

	s.BootOnMu.Lock()
	time.Sleep(100 * time.Millisecond)
	err := os.WriteFile("/proc/hexdeep_sub_pwr/boot_on", []byte(id), 0)
	time.Sleep(100 * time.Millisecond)
	s.BootOnMu.Unlock()
	return err
}

func (s *Slot) BootOff(id string) error {

	if !IsSlotIDValid(id) {
		return ErrInvalidID
	}

	s.BootOffMu.Lock()
	time.Sleep(100 * time.Millisecond)
	err := os.WriteFile("/proc/hexdeep_sub_pwr/boot_off", []byte(id), 0)
	time.Sleep(100 * time.Millisecond)
	s.BootOffMu.Unlock()
	return err
}
