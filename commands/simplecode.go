package commands

import "github.com/wilriker/goduetapiclient/types"

// SimpleCode performs a simple G/M/T-code
type SimpleCode struct {
	BaseCommand
	// Code to parse and execute
	Code string
	// Channel to execute this code on
	Channel types.CodeChannel
}

// NewSimpleCode creates a new SimpleCode for the given code and channel.
func NewSimpleCode(code string, channel types.CodeChannel) *SimpleCode {
	return &SimpleCode{
		BaseCommand: *NewBaseCommand("SimpleCode"),
		Code:        code,
		Channel:     channel,
	}
}
