package main

import (
	"fmt"
	"log"
	"time"

	"go.bug.st/serial"
)

func main() {
	// 串口配置
	mode := &serial.Mode{
		BaudRate: 115200,
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
	}

	// 打开串口
	port, err := serial.Open("/dev/ttyS3", mode)
	if err != nil {
		log.Fatalf("Failed to open serial port: %v", err)
	}
	defer port.Close()

	// 要发送的命令
	cmd := "show interface\r\n"
	fmt.Println("Sending:", cmd)

	_, err = port.Write([]byte(cmd))
	if err != nil {
		log.Fatalf("Failed to write to port: %v", err)
	}

	// 设置读取超时
	port.SetReadTimeout(3 * time.Second)

	buffer := make([]byte, 256)

	fmt.Println("Reading response...")

	for {
		n, err := port.Read(buffer)
		if err != nil {
			log.Fatalf("Read error: %v", err)
		}

		if n == 0 {
			// 超时，没有更多输出
			fmt.Println("No more data, exit.")
			break
		}

		fmt.Print(string(buffer[:n]))
	}
}
