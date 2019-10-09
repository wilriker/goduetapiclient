package commands

import "github.com/wilriker/goduetapiclient/types"

type WriteMessage struct {
	BaseCommand
	Type          types.MessageType
	Content       string
	OutputMessage bool
	LogMessage    bool
}

func NewWriteMessage(mType types.MessageType, content string, outputMessage, logMessage bool) *WriteMessage {
	return &WriteMessage{
		BaseCommand:   *NewBaseCommand("WriteMessage"),
		Type:          mType,
		Content:       content,
		OutputMessage: outputMessage,
		LogMessage:    logMessage,
	}
}
