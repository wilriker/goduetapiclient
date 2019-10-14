package machine

type MachineMode string

const (
	FFF              MachineMode = "FFF"
	CNC                          = "CNC"
	MachineModeLaser             = "Laser"
)

type MachineStatus string

const (
	Updating          MachineStatus = "Updating"
	MachineStatusOff                = "Off"
	Halted                          = "Halted"
	Pausing                         = "Pausing"
	Paused                          = "Paused"
	Resuming                        = "Resuming"
	Processing                      = "Processing"
	Simulating                      = "Simulating"
	Busy                            = "Busy"
	ChangingTool                    = "ChangingTool"
	MachineStatusIdle               = "Idle"
)

type State struct {
	AtxPower       *bool
	Beep           BeepDetails
	CurrentTool    int64
	DisplayMessage string
	LogFile        string
	Mode           MachineMode
	Status         MachineStatus
}

type BeepDetails struct {
	Frequency uint64
	Duration  float64
}
