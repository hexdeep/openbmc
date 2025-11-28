package proc

type Proc struct {
	SubPwrStatus  SubPwrStatus
	SlotPwrOn     SlotPwrOn
	SlotPwrOff    SlotPwrOff
	MainPwrStatus MainPwrStatus
	MainPwrOn     MainPwrOn
	MainPwrOff    MainPwrOff
	SlotBootOn    SlotBootOn
	SlotBootOff   SlotBootOff
	SlotSerial    SlotSerial
	SwitchSerial  SwitchSerial
	BMCSerial     BMCSerial
}
