package commands

import "github.com/wilriker/goduetapiclient/types"

type Resolve struct {
	BaseCommand
	Type    types.MessageType
	Content string
}

func NewResolve(mType types.MessageType, content string) *Resolve {
	return &Resolve{
		BaseCommand: *NewBaseCommand("Resolve"),
		Type:        mType,
		Content:     content,
	}
}
