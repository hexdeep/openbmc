package proc

import (
	"errors"
	"os"
	"slices"
	"sync"
)

type Power struct {
	PowerStatusMu sync.Mutex
	PowerOnMu     sync.Mutex
	PowerOffMu    sync.Mutex
}

func NewPower() *Power {
	return &Power{}
}

var (
	ErrInvalidID = errors.New("invalid power id")
)

var PowerStatusMap = map[string]map[string]bool{
	"0x0022\r\n": {"3": false, "7": false},
	"0x002b\r\n": {"3": true, "7": false},
	"0x00b2\r\n": {"3": false, "7": true},
	"0x00bb\r\n": {"3": true, "7": true},
}

func (p *Power) PowerStatus() (map[string]bool, error) {

	p.PowerStatusMu.Lock()
	defer p.PowerStatusMu.Unlock()
	data, err := os.ReadFile("/proc/hexdeep_main_pwr/pwr_status")
	if err != nil {
		return nil, err
	}

	return PowerStatusMap[string(data)], nil
}

func (p *Power) PowerOn(id string) error {

	if !IsMainPwrIDValid(id) {
		return ErrInvalidID
	}

	p.PowerOnMu.Lock()
	defer p.PowerOnMu.Unlock()
	return os.WriteFile("/proc/hexdeep_main_pwr/pwr_on", []byte(id), 0)
}

func (p *Power) PowerOff(id string) error {

	if !IsMainPwrIDValid(id) {
		return ErrInvalidID
	}

	p.PowerOffMu.Lock()
	defer p.PowerOffMu.Unlock()
	return os.WriteFile("/proc/hexdeep_main_pwr/pwr_off", []byte(id), 0)
}

func IsMainPwrIDValid(id string) bool {
	return slices.Contains([]string{"3", "7"}, id)
}
