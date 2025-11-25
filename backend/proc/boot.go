package proc

import (
	"os"
	"sync"
)

type BootOn struct {
	Mu sync.Mutex
}

func (b *BootOn) Do(id string) error {

	if !IsSlotIDValid(id) {
		return ErrInvalidID
	}

	b.Mu.Lock()
	defer b.Mu.Unlock()
	return os.WriteFile("/proc/hexdeep_sub_pwr/boot_on", []byte(id), 0)
}

type BootOff struct {
	Mu sync.Mutex
}

func (b *BootOff) Do(id string) error {

	if !IsSlotIDValid(id) {
		return ErrInvalidID
	}

	b.Mu.Lock()
	defer b.Mu.Unlock()
	return os.WriteFile("/proc/hexdeep_sub_pwr/boot_off", []byte(id), 0)
}
