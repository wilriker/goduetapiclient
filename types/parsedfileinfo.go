package types

import (
	"time"
)

type ParsedFileInfo struct {
	FileName         string
	Size             uint64
	LastModified     time.Time // TODO: This will probably need adjustment/custom type
	Height           float64
	FirstLayerHeight float64
	LayerHeight      float64
	NumLayers        uint64
	Filament         []float64
	GeneratedBy      string
	PrintTime        uint64
	SimulatedTime    uint64
}
