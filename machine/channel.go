package machine

import (
	"errors"
)

type Compatibility uint64

const (
	Me Compatibility = iota
	RepRapFirmware
	Marlin
	Teacup
	Sprinter
	Repetier
	NanoDLP
)

type Channel struct {
	Compatibility       Compatibility
	Feedrate            float64
	RelativeExtrusion   bool
	VolumetricExtrusion bool
	RelativePositioning bool
	UsingInches         bool
	StackDepth          uint8
	LineNumber          int64
}

func (c *Channel) Assign(other interface{}) error {
	if other == nil {
		return errors.New("Cannot assign from nil value")
	}
	co, ok := other.(Channel)
	if !ok {
		return errors.New("Invalid type")
	}
	c.Compatibility = co.Compatibility
	c.Feedrate = co.Feedrate
	c.RelativeExtrusion = co.RelativeExtrusion
	c.VolumetricExtrusion = co.VolumetricExtrusion
	c.RelativePositioning = co.RelativePositioning
	c.UsingInches = co.UsingInches
	c.StackDepth = co.StackDepth
	c.LineNumber = co.LineNumber
	return nil
}
