package connection

import (
	"github.com/wilriker/goduetapiclient/connection/initmessages"
)

type CommandConnection struct {
	BaseCommandConnection
}

func (cc *CommandConnection) Connect(socketPath string) error {
	return cc.BaseConnection.Connect(initmessages.NewCommandInitMessage(), socketPath)
}
