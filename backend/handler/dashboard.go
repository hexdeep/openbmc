package handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func readCPU(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	if sc.Scan() {
		parts := strings.Fields(sc.Text())
		if parts[0] == "cpu" {
			return fmt.Sprintf(
				`{"user": %s, "nice": %s, "system": %s, "idle": %s}`,
				parts[1], parts[2], parts[3], parts[4],
			), nil
		}
	}
	return "", fmt.Errorf("no cpu line found")
}

func PushCPUStatus(h *Handler, c echo.Context, send chan<- string) error {

	for {

		cpu, err := readCPU("/proc/meminfo")
		if err != nil {
			return err
		}

		send <- fmt.Sprintf("{\"mem\": %s}", cpu)

		time.Sleep(time.Duration(h.Config.StatusDuration))
	}
}
