package machine

import "github.com/wilriker/goduetapiclient/commands"

// MachineModel represents the full machine model as maintained by DuetControlServer
type MachineModel struct {
	Channels      Channels
	Directories   Directories
	Electronics   Electronics
	Fans          []Fan
	Heat          Heat
	HttpEndpoints []HttpEndpoint
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
	UserSessions  []UserSession
	UserVariables []UserVariable
}

func NewMachineModel() *MachineModel {
	return &MachineModel{}
}
