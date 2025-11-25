package proc

import (
	"errors"
	"os"
	"slices"
	"sync"
)

var (
	ErrInvalidID = errors.New("invalid id")
)

type MainPwrStatus struct {
	Mu sync.Mutex
}

type MainPowerStatus struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

var MainPowerStatusMap = map[string][]MainPowerStatus{
	"0x0022\r\n": {
		{"3", "上方主电源", false},
		{"7", "下方主电源", false},
	},
	"0x002b\r\n": {
		{"3", "上方主电源", true},
		{"7", "下方主电源", false},
	},
	"0x00b2\r\n": {
		{"3", "上方主电源", false},
		{"7", "下方主电源", true},
	},
	"0x00bb\r\n": {
		{"3", "上方主电源", true},
		{"7", "下方主电源", true},
	},
}

func (m *MainPwrStatus) Get() ([]MainPowerStatus, error) {

	m.Mu.Lock()
	data, err := os.ReadFile("/proc/hexdeep_main_pwr/pwr_status")
	m.Mu.Unlock()
	if err != nil {
		return nil, err
	}

	return MainPowerStatusMap[string(data)], nil
}

type MainPwrOn struct {
	Mu sync.Mutex
}

func (m *MainPwrOn) Do(id string) error {

	if !IsMainPwrIDValid(id) {
		return ErrInvalidID
	}

	m.Mu.Lock()
	defer m.Mu.Unlock()
	return os.WriteFile("/proc/hexdeep_main_pwr/pwr_on", []byte(id), 0)
}

type MainPwrOff struct {
	Mu sync.Mutex
}

func (m *MainPwrOff) Do(id string) error {

	if !IsMainPwrIDValid(id) {
		return ErrInvalidID
	}

	m.Mu.Lock()
	defer m.Mu.Unlock()
	return os.WriteFile("/proc/hexdeep_main_pwr/pwr_off", []byte(id), 0)
}

func IsMainPwrIDValid(id string) bool {
	return slices.Contains([]string{"3", "7"}, id)
}
