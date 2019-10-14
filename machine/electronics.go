package machine

type Electronics struct {
	Version         string
	Type            string
	ShortName       string
	Name            string
	Revision        string
	Firmware        Firmware
	ProcessorId     string
	VIn             MinMaxCurrent
	McuTemp         MinMaxCurrent
	ExpansionBoards []ExpansionBoard
}

type Firmware struct {
	Name    string
	Version string
	Date    string
}

type MinMaxCurrent struct {
	Current float64
	Min     float64
	Max     float64
}

type ExpansionBoard struct {
	ShortName  string
	Name       string
	Revision   string
	Firmware   Firmware
	VIn        MinMaxCurrent
	McuTemp    MinMaxCurrent
	MaxHeaters int64
	MaxMotors  int64
}
