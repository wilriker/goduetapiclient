package machine

type ScannerStatus string

const (
	Disconnected   ScannerStatus = "D"
	Idle                         = "I"
	Scanning                     = "S"
	PostProcessing               = "P"
	Calibrating                  = "C"
	Uploading                    = "U"
)

type Scanner struct {
	Progress float64
	Status   ScannerStatus
}
