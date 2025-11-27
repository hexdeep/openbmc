package proc

import (
	"os"
	"os/exec"
	"strings"
	"sync"
	"syscall"
	"time"

	"go.bug.st/serial"
)

type SwitchSerial struct {
	Mu  sync.Mutex
	TTY *os.Process
}

func (s *SwitchSerial) ShowInterface(timeout time.Duration) (map[string]bool, error) {

	s.Mu.Lock()
	rawResult, err := SerialCommand(
		&serial.Mode{
			BaudRate: 115200,
			DataBits: 8,
			Parity:   serial.NoParity,
			StopBits: serial.OneStopBit,
		},
		"/dev/ttyS3",
		timeout,
		"show interface\n",
	)
	s.Mu.Unlock()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(rawResult, "\n")
	result := make(map[string]bool)

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

		if fields[2] == "up" {
			result[Idx(fields, 0)] = true
		} else {
			result[Idx(fields, 0)] = false
		}
	}

	return result, nil
}

func (s *SwitchSerial) OpenTTY() error {

	s.Mu.Lock()
	cmd := exec.Command("./ttyd.aarch64", "-p", "7600", "-W", "microcom", "-s", "115200", "/dev/ttyS3")
	if err := cmd.Start(); err != nil {
		return err
	}

	s.TTY = cmd.Process
	return nil
}

func (s *SwitchSerial) CloseTTY() error {

	if err := s.TTY.Signal(syscall.SIGTERM); err != nil {
		return err
	}

	s.Mu.Unlock()
	s.TTY = nil
	return nil
}
