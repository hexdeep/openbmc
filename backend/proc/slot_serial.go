package proc

import (
	"context"
	"strings"
	"sync"

	"go.bug.st/serial"
)

type SlotSerial struct {
	Mu [48]sync.Mutex
}

func (s *SlotSerial) IsActive(ctx context.Context, id string) bool {

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

func (s *SlotSerial) GetMacIP(ctx context.Context, id string) (string, string, error) {

	ttyId, err := SlotIDToTTY(id)
	if err != nil {
		return "", "", err
	}

	data, err := SerialCommand(&serial.Mode{BaudRate: 1500000}, ttyId, ctx, "ifconfig\n")
	if err != nil {
		return "", "", err
	}

	lines := strings.Split(data, "\n")

	return strings.Fields(lines[0])[4], strings.Split(strings.Fields(lines[1])[1], ":")[1], nil
}

func (s *SlotSerial) GetTemp(ctx context.Context, id string) (string, error) {

	ttyId, err := SlotIDToTTY(id)
	if err != nil {
		return "", nil
	}

	data, err := SerialCommand(&serial.Mode{BaudRate: 1500000}, ttyId, ctx, "cat /sys/class/thermal/thermal_zone0/temp\n")
	if err != nil {
		return "", err
	}

	return data[:len(data)-3], nil
}
