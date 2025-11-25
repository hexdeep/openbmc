package proc

import (
	"math"
	"strconv"
	"strings"
	"sync"
	"time"

	"go.bug.st/serial"
)

type SlotSerial struct {
	Mu [48]sync.Mutex
}

func (s *SlotSerial) IsActive(id string, timeout time.Duration) bool {

	ttyId, err := SlotIDToTTY(id)
	if err != nil {
		return false
	}

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

	data, err := SerialCommand(&serial.Mode{BaudRate: 1500000}, ttyId, timeout, "ifconfig\n")
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

	data, err := SerialCommand(&serial.Mode{BaudRate: 1500000}, ttyId, timeout, "cat /sys/class/thermal/thermal_zone0/temp\n")
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

	data, err := SerialCommand(&serial.Mode{BaudRate: 1500000}, ttyId, timeout, "cat /proc/meminfo\n")
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

	data, err := SerialCommand(&serial.Mode{BaudRate: 1500000}, ttyId, timeout, "cat /proc/uptime\n")
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

	data, err := SerialCommand(&serial.Mode{BaudRate: 1500000}, ttyId, timeout, "cat /proc/loadavg\n")
	if err != nil {
		return 0, err
	}

	value, err := strconv.ParseFloat(Idx(strings.Fields(data), 0), 64)
	if err != nil {
		return 0, err
	}

	return math.Round(value*1250) / 100, nil
}
