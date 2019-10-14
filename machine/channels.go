package machine

import (
	"errors"

	"github.com/wilriker/goduetapiclient/types"
)

type Channels struct {
	HTTP      Channel
	Telnet    Channel
	File      Channel
	USB       Channel
	AUX       Channel
	Daemon    Channel
	CodeQueue Channel
	LCD       Channel
	SPI       Channel
	AutoPause Channel
}

func NewChannels() Channels {
	return Channels{
		Telnet: Channel{Compatibility: Marlin},
		USB:    Channel{Compatibility: Marlin},
	}
}

func (ch *Channels) Get(cc types.CodeChannel) Channel {
	switch cc {
	case types.HTTP:
		return ch.HTTP
	case types.Telnet:
		return ch.Telnet
	case types.File:
		return ch.File
	case types.USB:
		return ch.USB
	case types.AUX:
		return ch.AUX
	case types.Daemon:
		return ch.Daemon
	case types.CodeQueue:
		return ch.CodeQueue
	case types.LCD:
		return ch.LCD
	case types.SPI:
		return ch.SPI
	case types.AutoPause:
		return ch.AutoPause
	default:
		return ch.SPI
	}
}

func (ch *Channels) Assign(other interface{}) error {
	if other == nil {
		return errors.New("Cannot assign from nil value")
	}
	co, ok := other.(Channels)
	if !ok {
		return errors.New("Invalid type")
	}
	ch.HTTP.Assign(co.HTTP)
	ch.Telnet.Assign(co.Telnet)
	ch.File.Assign(co.File)
	ch.USB.Assign(co.USB)
	ch.AUX.Assign(co.AUX)
	ch.Daemon.Assign(co.Daemon)
	ch.CodeQueue.Assign(co.CodeQueue)
	ch.LCD.Assign(co.LCD)
	ch.SPI.Assign(co.SPI)
	ch.AutoPause.Assign(co.AutoPause)
	return nil
}
