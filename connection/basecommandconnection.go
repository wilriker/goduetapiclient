package connection

import (
	"encoding/json"

	"github.com/wilriker/goduetapiclient/commands"
	"github.com/wilriker/goduetapiclient/machine"
	"github.com/wilriker/goduetapiclient/types"
)

type BaseCommandConnection struct {
	BaseConnection
}

func (bcc *BaseCommandConnection) Flush(channel types.CodeChannel) (bool, error) {
	r, err := bcc.PerformCommand(commands.NewFlush(channel))
	if err != nil {
		return false, err
	}
	return r.IsSuccess(), nil
}

func (bcc *BaseCommandConnection) GetFileInfo(fileName string) (*types.ParsedFileInfo, error) {
	r, err := bcc.PerformCommand(commands.NewGetFileInfo(fileName))
	if err != nil {
		return nil, err
	}
	pfi := r.GetResult().(types.ParsedFileInfo)
	return &pfi, nil
}

func (bcc *BaseCommandConnection) PerformCode(code commands.Code) (*commands.CodeResult, error) {
	r, err := bcc.PerformCommand(&code)
	if err != nil {
		return nil, err
	}
	cr := r.GetResult().(commands.CodeResult)
	return &cr, nil
}

func (bcc *BaseCommandConnection) PerformSimpleCode(code string, channel types.CodeChannel) (string, error) {
	r, err := bcc.PerformCommand(commands.NewSimpleCode(code, channel))
	if err != nil {
		return "", err
	}
	return r.GetResult().(string), nil
}

func (bcc *BaseCommandConnection) GetMachineModel() (*machine.MachineModel, error) {
	r, err := bcc.PerformCommand(commands.NewGetMachineModel())
	if err != nil {
		return nil, err
	}
	mm := r.GetResult().(machine.MachineModel)
	return &mm, nil
}

func (bcc *BaseCommandConnection) GetSerializedMachineModel() (json.RawMessage, error) {
	var raw json.RawMessage
	err := bcc.Send(commands.NewGetMachineModel())
	if err != nil {
		return raw, err
	}
	err = bcc.Receive(&raw)
	if err != nil {
		return raw, err
	}
	return raw, nil
}

func (bcc *BaseCommandConnection) ResolvePath(path string) (string, error) {
	r, err := bcc.PerformCommand(commands.NewResolvePath(path))
	if err != nil {
		return "", err
	}
	return r.GetResult().(string), nil
}

func (bcc *BaseCommandConnection) SyncMachineModel() error {
	_, err := bcc.PerformCommand(commands.NewSyncMachineModel())
	return err
}
