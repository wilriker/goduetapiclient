package machine

type EndstopAction uint64

const (
	EndstopActionNone EndstopAction = iota
	ReduceSpeed
	StopDriver
	StopAxis
	StopAll
)

type EndstopType uint64

const (
	ActiveLow EndstopType = iota
	ActiveHigh
	EntstopTypeProbe
	MotorStallAny
	MotorStallIndividual
)

type ProbeType uint64

const (
	ProbeTypeNone ProbeType = iota
	Unmodulated
	Modulated
	Switch
	BLTouch
	MotorLoadDetection
)

type Sensors struct {
	Endstops []Endstop
	Probes   []Probe
}

type Endstop struct {
	Action    EndstopAction
	Triggered bool
	Type      EndstopType
	Probe     int64
}

type Probe struct {
	Type          ProbeType
	Value         int64
	Threshold     int64
	Speed         float64
	DiveHeight    float64
	TriggerHeight float64
	Filtered      bool
	Inverted      bool
	RecoveryTime  float64
	TravelSpeed   float64
	MaxProbeCount uint64
	Tolerance     float64
	DisablesBed   bool
	Persistent    bool
}
