package main

type TV interface {
	GetStatus() bool
	GetModel() string
	switchOFF() bool
	switchOn() bool
}
type LgTV struct {
	status bool
	model  string
}

type SamsungTV struct {
	status bool
	model  string
}

func (tv *SamsungTV) GetStatus() bool {
	return tv.status
}

func (tv *SamsungTV) GetModel() string {
	return tv.model
}

func (tv *SamsungTV) switchOn() bool {
	tv.status = true
	return true
}

func (tv *SamsungTV) switchOFF() bool {
	tv.status = false
	return true
}

func (tv *SamsungTV) SamsungHub() string {
	return "SamsungHub"
}
func (tv *LgTV) GetStatus() bool {
	return tv.status
}

func (tv *LgTV) GetModel() string {
	return tv.model
}

func (tv *LgTV) switchOn() bool {
	tv.status = true
	return true
}

func (tv *LgTV) switchOFF() bool {
	tv.status = false
	return true
}

func (tv *LgTV) LGHub() string {
	return "LGHub"
}

