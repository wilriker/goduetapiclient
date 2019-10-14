package machine

type Tool struct {
	Number           int64
	Active           []float64
	Standby          []float64
	Name             string
	FilamentExtruder int64
	Filament         string
	Fans             []int64
	Heaters          []int64
	Extruders        []int64
	Mix              []float64
	Spindle          int64
	Axes             [][]uint64
	Offsets          []float64
	OffsetsProbed    int64
}
