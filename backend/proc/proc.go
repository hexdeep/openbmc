package proc

type Proc struct {
	SubPwrStatus  SubPwrStatus
	SlotPwrOn     SlotPwrOn
	SlotPwrOff    SlotPwrOff
	MainPwrStatus MainPwrStatus
	MainPwrOn     MainPwrOn
	MainPwrOff    MainPwrOff
	BootOn        BootOn
	BootOff       BootOff
	SlotSerial    SlotSerial
	SwitchSerial  SwitchSerial
	BMCSerial     BMCSerial
}
