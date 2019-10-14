package machine

import "github.com/wilriker/goduetapiclient/commands"

type MachineModel struct {
	Channels      Channels
	Electronics   Electronics
	Fans          []Fan
	Heat          Heat
	Job           Job
	Lasers        []Laser
	MessageBox    MessageBox
	Messages      []commands.Message
	Move          Move
	Network       Network
	Scanner       Scanner
	Sensors       Sensors
	Spindles      []Spindle
	State         State
	Storages      []Storage
	Tools         []Tool
	UserVariables []UserVariable
}

func NewMachineModel() *MachineModel {
	return &MachineModel{}
}
