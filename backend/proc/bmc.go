package proc

import (
	"os"
	"os/exec"
	"sync"
	"syscall"
)

type BMCSerial struct {
	Mu  sync.Mutex
	TTY *os.Process
}

func (s *BMCSerial) OpenTTY() error {

	s.Mu.Lock()

	cmd := exec.Command("./ttyd.aarch64", "-p", "7500", "-W", "sh")
	if err := cmd.Start(); err != nil {
		return err
	}

	s.TTY = cmd.Process
	return nil
}

func (s *BMCSerial) CloseTTY() error {

	if err := s.TTY.Signal(syscall.SIGTERM); err != nil {
		return err
	}

	s.Mu.Unlock()
	s.TTY = nil
	return nil
}
