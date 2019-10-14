package machine

import "github.com/wilriker/goduetapiclient/types"

type Job struct {
	File              types.ParsedFileInfo
	FilePosition      uint64
	LastFileName      string
	LastFileAborted   bool
	LastFileCancelled bool
	LastFileSimulated bool
	ExtrudedRaw       []float64
	Duration          float64
	Layer             float64
	LayerTime         float64
	Layers            []Layer
	WarmUpDuration    float64
	TimesLeft         TimesLeft
}

type Layer struct {
	Duration        float64
	Height          float64
	Filament        []float64
	FractionPrinted float64
}

type TimesLeft struct {
	File     float64
	Filament float64
	Layer    float64
}
