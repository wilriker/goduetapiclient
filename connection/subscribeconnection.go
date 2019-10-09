package connection

import (
	"github.com/wilriker/goduetapiclient/commands"
	"github.com/wilriker/goduetapiclient/connection/initmessages"
	"github.com/wilriker/goduetapiclient/machine"
)

type SubscribeConnection struct {
	BaseConnection
	Mode   initmessages.SubscriptionMode
	Filter string
}

func (sc *SubscribeConnection) Connect(mode initmessages.SubscriptionMode, filter, socketPath string) error {
	sc.Mode = mode
	sc.Filter = filter
	sim := initmessages.NewSubscribeInitMessage(mode, filter)
	sc.BaseConnection.Connect(sim, socketPath)
	return nil
}

func (sc *SubscribeConnection) GetMachineModel() (*machine.MachineModel, error) {
	m := machine.NewMachineModel()
	err := sc.Receive(m)
	if err != nil {
		return nil, err
	}
	err = sc.Send(commands.NewAcknowledge())
	if err != nil {
		return nil, err
	}
	return m, nil
}
