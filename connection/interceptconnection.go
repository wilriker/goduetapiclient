package connection

import (
	"github.com/wilriker/goduetapiclient/commands"
	"github.com/wilriker/goduetapiclient/connection/initmessages"
	"github.com/wilriker/goduetapiclient/types"
)

type InterceptConnection struct {
	BaseCommandConnection
	Mode initmessages.InterceptionMode
}

func (ic *InterceptConnection) Connect(mode initmessages.InterceptionMode, socketPath string) error {
	ic.Mode = mode
	iim := initmessages.NewInterceptInitMessage(mode)
	return ic.BaseConnection.Connect(iim, socketPath)
}

func (ic *InterceptConnection) ReceiveCode() (*commands.Code, error) {
	c := commands.NewCode()
	err := ic.Receive(&c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (ic *InterceptConnection) CancelCode() error {
	return ic.Send(commands.NewCancel())
}

func (ic *InterceptConnection) IgnoreCode() error {
	return ic.Send(commands.NewIgnore())
}

func (ic *InterceptConnection) ResolveCode(mType types.MessageType, content string) error {
	return ic.Send(commands.NewResolve(mType, content))
}
