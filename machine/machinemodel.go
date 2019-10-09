package machine

import "github.com/wilriker/goduetapiclient/commands"

// TODO: Make these real impls
type Channels map[string]interface{}
type Electronics map[string]interface{}
type Fan map[string]interface{}
type Heat map[string]interface{}
type Job map[string]interface{}
type Laser map[string]interface{}
type MessageBox map[string]interface{}
type Move map[string]interface{}
type Network map[string]interface{}
type Scanner map[string]interface{}
type Sensors map[string]interface{}
type Spindle map[string]interface{}
type State map[string]interface{}
type Storage map[string]interface{}
type Tool map[string]interface{}
type UserVariable map[string]interface{}

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
