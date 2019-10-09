package commands

import "github.com/wilriker/goduetapiclient/types"

type Flush struct {
	BaseCommand
	Channel types.CodeChannel
}

func NewFlush(channel types.CodeChannel) *Flush {
	return &Flush{
		BaseCommand: *NewBaseCommand("Flush"),
		Channel:     channel,
	}
}
