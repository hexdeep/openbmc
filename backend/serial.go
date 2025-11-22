package main

import (
	"errors"
	"time"

	"go.bug.st/serial"
)

func (h *Handler) SerialCommand(command string) (string, error) {

	port, err := serial.Open(h.Config.SerialFile, h.Config.Serial)
	if err != nil {
		return "", err
	}
	defer port.Close()

	_, err = port.Write([]byte(command))
	if err != nil {
		return "", err
	}

	port.SetReadTimeout(time.Duration(h.Config.SerialTimeout) * time.Millisecond)

	buffer := make([]byte, 256)
	result := make([]byte, 0, 1024)

	for {
		n, err := port.Read(buffer)
		if err != nil {
			return "", err
		}
		if n == 0 {
			break
		}
		result = append(result, buffer[:n]...)
	}

	if len(result) == 0 {
		return "", errors.New("no response from serial device within timeout")
	}

	return string(result), nil

}
