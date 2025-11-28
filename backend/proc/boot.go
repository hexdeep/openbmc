package proc

import (
	"os"
	"sync"
	"time"
)

type SlotBootOn struct {
	Mu sync.Mutex
}

func (b *SlotBootOn) Do(id string) error {

	if !IsSlotIDValid(id) {
		return ErrInvalidID
	}

	b.Mu.Lock()
	time.Sleep(100 * time.Millisecond)
	err := os.WriteFile("/proc/hexdeep_sub_pwr/boot_on", []byte(id), 0)
	time.Sleep(100 * time.Millisecond)
	b.Mu.Unlock()
	return err
}

type SlotBootOff struct {
	Mu sync.Mutex
}

func (b *SlotBootOff) Do(id string) error {

	if !IsSlotIDValid(id) {
		return ErrInvalidID
	}

	b.Mu.Lock()
	time.Sleep(100 * time.Millisecond)
	err := os.WriteFile("/proc/hexdeep_sub_pwr/boot_off", []byte(id), 0)
	time.Sleep(100 * time.Millisecond)
	b.Mu.Unlock()
	return err
}
