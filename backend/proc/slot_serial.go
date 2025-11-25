package proc

import (
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

	return strings.Fields(lines[0])[4], strings.Split(strings.Fields(lines[1])[1], ":")[1], nil
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

	return data[:len(data)-3], nil
}
