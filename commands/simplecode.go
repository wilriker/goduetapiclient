package commands

import "github.com/wilriker/goduetapiclient/types"

type SimpleCode struct {
	BaseCommand
	Code    string
	Channel types.CodeChannel
}

func NewSimpleCode(code string, channel types.CodeChannel) *SimpleCode {
	return &SimpleCode{
		BaseCommand: *NewBaseCommand("SimpleCode"),
		Code:        code,
		Channel:     channel,
	}
}
