package machine

type GeometryType string

const (
	Cartesian   GeometryType = "cartesian"
	CoreXY                   = "corexy"
	CoreXYU                  = "corexyu"
	CoreXYUV                 = "corexyuv"
	CoreXZ                   = "corexz"
	Handprinter              = "hangprinter"
	Delta                    = "delta"
	Polar                    = "polar"
	RotaryDelta              = "rotary delta"
	Scara                    = "scara"
	Unknown                  = "unknown"
)

// Move holds information about the move subsystem
type Move struct {
	Axes         []Axis
	BabystepZ    float64
	CurrentMove  CurrentMove
	Compensation string
	// HeightmapFile is the path to the current heightmap file if Compensation is "Mesh"
	HeightmapFile string
	Drives        []Drive
	Extruders     []Extruder
	Geometry      Geometry
	Idle          MotorsIdleControl
	// ProbeGrid holds information about the configured mesh compensation (see M557)
	ProbeGrid            ProbeGrid
	SpeedFactor          float64
	CurrentWorkplace     int64
	WorkplaceCoordinates [][]float64
}

type Axis struct {
	Letter          string
	Drives          []int64
	Homed           bool
	MachinePosition *float64
	Min             float64
	MinEndstop      *int64
	MinProbed       bool
	Max             float64
	MaxEndstop      *int64
	MaxProbed       bool
	Visible         bool
}

type CurrentMove struct {
	RequestedSpeed float64
	TopSpeed       float64
}

type Drive struct {
	Position      float64
	Microstepping DriveMicrostepping
	Current       uint64
	Acceleration  float64
	MinSpeed      float64
	MaxSpeed      float64
}

type DriveMicrostepping struct {
	Value        uint64
	Interpolated bool
}

type Extruder struct {
	Drives    []int64
	Factor    float64
	NonLinear ExtruderNonLinear
}

type ExtruderNonLinear struct {
	A           float64
	B           float64
	UpperLimit  float64
	Temperature float64
}

type Geometry struct {
	Type               GeometryType
	Anchors            []float64
	PrintRadius        float64
	Diagonals          []float64
	Radius             float64
	HomedHeight        float64
	AngleCorrections   []float64
	EndstopAdjustments []float64
	Tilt               []float64
}

type MotorsIdleControl struct {
	Timeout float64
	Factor  float64
}
