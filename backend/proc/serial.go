package proc

import (
	"context"
	"errors"
	"strings"
	"time"

	"go.bug.st/serial"
)

func SerialCommand(mode *serial.Mode, portName string, ctx context.Context, command string) (string, error) {

	if strings.HasSuffix(command, "\n") {
		command = command + "\n"
	}

	port, err := serial.Open(portName, mode)
	if err != nil {
		return "", err
	}
	defer port.Close()

	_, err = port.Write([]byte(command))
	if err != nil {
		return "", err
	}

	buf := make([]byte, 256)
	result := make([]byte, 0)

	port.SetReadTimeout(time.Second)

	for {
		n, err := port.Read(buf)
		if err != nil {
			return "", err
		}
		if n == 0 {
			break
		}
		result = append(result, buf[:n]...)
	}

	if len(result) == 0 {
		return "", errors.New("no response from serial device within timeout")
	}

	lines := strings.Split(string(result), "\n")

	// 删除第一行（通常是命令回显，如 "cat /proc/stat"）
	if len(lines) > 0 {
		lines = lines[1:]
	}

	// 删除最后一行（通常是 shell 提示符，如 "HexDeep:~# "）
	if len(lines) > 0 {
		last := strings.TrimSpace(lines[len(lines)-1])
		if last == "" || strings.Contains(last, ":") || strings.HasSuffix(last, "#") {
			lines = lines[:len(lines)-1]
		}
	}

	// 合并干净的数据
	cleaned := strings.Join(lines, "\n")
	cleaned = strings.TrimSpace(cleaned)

	if len(cleaned) > 9 {
		cleaned = cleaned[9:]
	}

	return string(cleaned), nil
}
