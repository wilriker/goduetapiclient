package machine

// EndstopAction is the action performed when an endstop is hit
type EndstopAction uint64

const (
	// EndstopActionNone means no action
	EndstopActionNone EndstopAction = iota
	// ReduceSpeed because an endstop or Z-probe is close to triggering (analog only)
	ReduceSpeed
	// StopDriver stops a single motor driver
	StopDriver
	// StopAxis stops all drivers for an axis
	StopAxis
	// StopAll stops all drivers in the system
	StopAll
)

// EndstopType represents the type of a configured enstop
type EndstopType uint64

const (
	// ActiveLow will pull the endstop signal from HIGH to LOW when hit
	ActiveLow EndstopType = iota
	// ActiveHigh will pull the endstops signal from LOW to HIGH when hit
	ActiveHigh
	// EndstopTypeProbe represents a probe being used for this endstop (most often on Z)
	EndstopTypeProbe
	// MotorStallAny uses motor load detection and stops all when one motor stalls
	MotorStallAny
	// MotorStallIndividual uses motor load detection and stops each motor acoording to its
	// stall detection status
	MotorStallIndividual
)

// ProbeType represents supported probe types
type ProbeType uint64

const (
	// ProbeTypeNone for no probe
	ProbeTypeNone ProbeType = iota
	// Unmodulated i a simple unmodulated probe (like dc42's infrared probe)
	Unmodulated
	// Modulated probe (like the original one shipped with RepRapPro Ormerod)
	Modulated
	// Switch that is triggered when probe is activated
	Switch
	// BLTouch probe
	BLTouch
	// MotorLoadDetection provided by the stepper driver
	MotorLoadDetection
)

// Sensors holds information about sensors
type Sensors struct {
	// Endstops is a list of configured endstops
	Endstops []Endstop
	// Probes is a list of configured probes
	Probes []Probe
}

// Endstop holds information about an endstop
type Endstop struct {
	// Action to perform when this endstop is hit
	Action EndstopAction
	// Triggered represents the curent trigger state
	Triggered bool
	// Type of this endstop
	Type EndstopType
	// Probe is the index of the use probe (only valid if Type == EndstopTypeProbe)
	Probe int64
}

const (
	// DefaultTriggerThreshold of an analog probe
	DefaultTriggerThreshold = 500
	// DefaultProbingSpeed at which a probing move is performed (in mm/s)
	DefaultProbingSpeed = 2.0
	// DefaultTriggerHeight at which the probe is triggered (in mm)
	DefaultTriggerHeight = 0.7
	// DefaultTravelSpeed between probing locations (in mm/s)
	DefaultTravelSpeed = 100.0
	// DefaultTolerance for a deviation between to measurements (in mm)
	DefaultTolerance = 0.03
)

// Probe holds information about a configured probe
type Probe struct {
	// Type of the configured probe
	Type ProbeType
	// Value is the current analog value of the probe
	Value int64
	// SecondaryValues of the probe
	SecondaryValues []int64
	// Threshold at which the probe is considered to be triggered (0..1023)
	Threshold int64
	// Speed at which probing is performed (in mm/s)
	Speed float64
	// DiveHeight is how far above the probe point a probing move starts (in mm)
	DiveHeight float64
	// Offsets for X and Y (in mm)
	Offsets []float64
	// TriggerHeight is th  Z height at which the probe is triggered (in mm)
	TriggerHeight float64
	// Filtered is true if the probe signal is filtered
	Filtered bool
	// Inverted is true if the probe signal is inverted
	Inverted bool
	// RecoveryTime (in s)
	RecoveryTime float64
	// TravelSpeed when probing multiple points (in mm/s)
	TravelSpeed float64
	// MaxProbeCount is the maximum number of times to probe after a bad reading
	// was determined
	MaxProbeCount uint64
	// Tolerance is the allowed deviation between two measurements (in mm)
	Tolerance float64
	// DisablesBed is true if the bed heater(s) are disabled while probing
	DisablesBed bool
	// Persistent indicates if the probe parameters are supposed to be
	// saved to config-override.g
	Persistent bool
}
