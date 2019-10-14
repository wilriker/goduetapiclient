package machine

type HeaterState uint64

const (
	Off HeaterState = iota
	Standby
	Active
	Tuning
	Offline
)

type Heat struct {
	Beds              []BedOrChamber
	Chambers          []BedOrChamber
	ColdExtrusionTemp float64
	ColdRetractTemp   float64
	Extra             []ExtraHeater
	Heaters           []Heater
}

type BedOrChamber struct {
	Active  []float64
	Standby []float64
	Name    string
	Heaters []int64
}

type ExtraHeater struct {
	Current float64
	Name    string
	State   HeaterState
	Sensor  int64
}

type Heater struct {
	Current float64
	Name    string
	State   HeaterState
	Model   HeaterModel
	Max     float64
	Sensor  int64
}

type HeaterModel struct {
	Gain            float64
	TimeConstant    float64
	DeadTime        float64
	MaxPwm          float64
	StandardVoltage float64
	UsePID          bool
	CustomPID       bool
	P               float64
	I               float64
	D               float64
}
