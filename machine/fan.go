package machine

type Fan struct {
	Value        float64
	Name         string
	Rpm          uint64
	Inverted     bool
	Frequency    float64
	Min          float64
	Max          float64
	Blip         float64
	Thermostatic Thermostatic
	Pin          uint64
}

type Thermostatic struct {
	Control     bool
	Heaters     []int64
	Temperature float64
}
