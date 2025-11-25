package proc

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"go.bug.st/serial"
)

func SerialCommand(mode *serial.Mode, portName string, ctx context.Context, command string) (string, error) {
	port, err := serial.Open(portName, mode)
	if err != nil {
		return "", err
	}
	defer port.Close()

	// 写入指令
	_, err = port.Write([]byte(command))
	if err != nil {
		return "", err
	}

	buf := make([]byte, 256)
	result := make([]byte, 0)

	readErrCh := make(chan error, 1)
	dataCh := make(chan []byte, 1)

	// 读取 goroutine
	go func() {
		for {
			n, err := port.Read(buf)
			if err != nil {
				readErrCh <- err
				return
			}
			if n == 0 {
				// 认为串口返回结束
				dataCh <- result
				return
			}
			result = append(result, buf[:n]...)
		}
	}()

	// 监听 context 或读取结束
	select {
	case <-ctx.Done():
		return "", fmt.Errorf("serial read canceled or timed out: %w", ctx.Err())
	case err := <-readErrCh:
		return "", err
	case data := <-dataCh:
		// data 就是完整输出
		result = data
	}

	if len(result) == 0 {
		return "", errors.New("no response from serial device before context done")
	}

	// 按你原来的逻辑清洗数据
	lines := strings.Split(string(result), "\n")

	// 删除第一行（命令回显）
	if len(lines) > 0 {
		lines = lines[1:]
	}

	// 删除最后一行（shell 提示符）
	if len(lines) > 0 {
		last := strings.TrimSpace(lines[len(lines)-1])
		if last == "" || strings.Contains(last, ":") || strings.HasSuffix(last, "#") {
			lines = lines[:len(lines)-1]
		}
	}

	cleaned := strings.Join(lines, "\n")
	cleaned = strings.TrimSpace(cleaned)

	// 原代码中 additional slicing cleaned[9:]，保留原意
	if len(cleaned) > 9 {
		return cleaned[9:], nil
	}

	return cleaned, nil
}
