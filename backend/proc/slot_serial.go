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

type SlotSerial struct {
	Items [48]SlotSerialItem
}

func (s *SlotSerial) GetPort(id string) (string, error) {

	slot, err := strconv.Atoi(id)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(slot + 7500), nil
}

func (s *SlotSerial) GetItem(id string) (*SlotSerialItem, error) {

	slot, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("failed to convert id to int: %w\n", err)
	}

	if slot < 1 || slot > 48 {
		return nil, ErrInvalidID
	}

	return &s.Items[slot], nil
}

type SlotSerialItem struct {
	Mu  sync.Mutex
	TTY *os.Process
}

func (s *SlotSerial) IsActive(id string, timeout time.Duration) bool {

	ttyId, err := SlotIDToTTY(id)
	if err != nil {
		return false
	}

	item, err := s.GetItem(id)
	if err != nil {
		return false
	}

	if !item.Mu.TryLock() {
		return false
	}
	defer item.Mu.Unlock()

	port, err := serial.Open(ttyId, &serial.Mode{BaudRate: 1500000})
	if err != nil {
		return false
	}
	defer port.Close()

	if _, err := port.Write([]byte("\n")); err != nil {
		return false
	}

	buf := make([]byte, 256)

	port.SetReadTimeout(timeout)
	n, err := port.Read(buf)
	if err != nil || n == 0 {
		return false
	} else {
		return true
	}
}

func (s *SlotSerial) GetMacIP(id string, timeout time.Duration) (string, string, error) {

	ttyId, err := SlotIDToTTY(id)
	if err != nil {
		return "", "", err
	}

	item, err := s.GetItem(id)
	if err != nil {
		return "", "", err
	}

	if !item.Mu.TryLock() {
		return "", "", nil
	}
	data, err := SerialCommand(&serial.Mode{BaudRate: 1500000}, ttyId, timeout, "ifconfig\n")
	item.Mu.Unlock()
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

func (s *SlotSerial) GetTemp(id string, timeout time.Duration) (string, error) {

	ttyId, err := SlotIDToTTY(id)
	if err != nil {
		return "", nil
	}

	item, err := s.GetItem(id)
	if err != nil {
		return "", fmt.Errorf("failed to get slot temperature: %w\n", err)
	}

	if !item.Mu.TryLock() {
		return "", nil
	}
	data, err := SerialCommand(&serial.Mode{BaudRate: 1500000}, ttyId, timeout, "cat /sys/class/thermal/thermal_zone0/temp\n")
	item.Mu.Unlock()
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

func (s *SlotSerial) GetMem(id string, timeout time.Duration) (int, int, error) {

	ttyId, err := SlotIDToTTY(id)
	if err != nil {
		return 0, 0, nil
	}

	item, err := s.GetItem(id)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to get slot mem: %w\n", err)
	}

	if !item.Mu.TryLock() {
		return 0, 0, nil
	}
	data, err := SerialCommand(&serial.Mode{BaudRate: 1500000}, ttyId, timeout, "cat /proc/meminfo\n")
	item.Mu.Unlock()
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

func (s *SlotSerial) GetUpTime(id string, timeout time.Duration) (float64, error) {

	ttyId, err := SlotIDToTTY(id)
	if err != nil {
		return 0, err
	}

	item, err := s.GetItem(id)
	if err != nil {
		return 0, fmt.Errorf("failed to get slot up time: %v\n", err)
	}

	if !item.Mu.TryLock() {
		return 0, nil
	}
	data, err := SerialCommand(&serial.Mode{BaudRate: 1500000}, ttyId, timeout, "cat /proc/uptime\n")
	item.Mu.Unlock()
	if err != nil {
		return 0, err
	}

	value, err := strconv.ParseFloat(Idx(strings.Fields(data), 0), 64)
	if err != nil {
		return 0, err
	}

	return value, nil
}

func (s *SlotSerial) GetLoad(id string, timeout time.Duration) (float64, error) {

	ttyId, err := SlotIDToTTY(id)
	if err != nil {
		return 0, err
	}

	item, err := s.GetItem(id)
	if err != nil {
		return 0, err
	}

	if !item.Mu.TryLock() {
		return 0, nil
	}
	data, err := SerialCommand(&serial.Mode{BaudRate: 1500000}, ttyId, timeout, "cat /proc/loadavg\n")
	item.Mu.Unlock()
	if err != nil {
		return 0, err
	}

	value, err := strconv.ParseFloat(Idx(strings.Fields(data), 0), 64)
	if err != nil {
		return 0, err
	}

	return math.Round(value*1250) / 100, nil
}

func (s *SlotSerial) OpenTTY(id string) error {

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

	item.Mu.Lock()
	cmd := exec.Command(
		"./ttyd.aarch64", "-p", strconv.Itoa(7500+num), "-W", "microcom", "-s", "1500000", tty,
	)
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start slot tty: %v\n", err)
	}

	item.TTY = cmd.Process
	return nil
}

func (s *SlotSerial) CloseTTY(id string) error {

	item, err := s.GetItem(id)
	if err != nil {
		return fmt.Errorf("failed to get slot id in close tty: %v\n", err)
	}

	if err := item.TTY.Signal(syscall.SIGTERM); err != nil {
		return fmt.Errorf("failed to close slot tty: %v\n", err)
	}

	item.Mu.Unlock()
	item.TTY = nil
	return nil
}

func (t *SlotSerial) Flash(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, "./flash.sh")
	result, err := cmd.Output()
	if err != nil {
		fmt.Println(string(result))
		return err
	}
	return nil
}
