package proc

import (
	"context"
	"strings"
	"sync"

	"go.bug.st/serial"
)

type SwitchSerial struct {
	Mu sync.Mutex
}

func (s *SwitchSerial) ShowInterface(ctx context.Context) (map[string]string, error) {

	s.Mu.Lock()
	rawResult, err := SerialCommand(
		&serial.Mode{
			BaudRate: 115200,
			DataBits: 8,
			Parity:   serial.NoParity,
			StopBits: serial.OneStopBit,
		},
		"/dev/ttyS3",
		ctx,
		"show interface\n",
	)
	s.Mu.Unlock()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(rawResult, "\n")
	result := make(map[string]string)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "Switch_config") ||
			strings.HasPrefix(line, "port ") ||
			strings.HasPrefix(line, "HexDeep") {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) < 5 {
			continue
		}

		result[fields[0]] = fields[2]
	}

	return result, nil
}
