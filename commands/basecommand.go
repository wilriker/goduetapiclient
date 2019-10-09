package commands

type Command interface {
	GetCommand() string
	GetSourceConnection() int64
}

type BaseCommand struct {
	Command          string
	SourceConnection int64
}

func (bc *BaseCommand) GetCommand() string {
	return bc.Command
}

func (bc *BaseCommand) GetSourceConnection() int64 {
	return bc.SourceConnection
}

func NewBaseCommand(command string) *BaseCommand {
	return &BaseCommand{Command: command}
}

func NewAcknowledge() *BaseCommand {
	return NewBaseCommand("Acknowledge")
}

func NewCancel() *BaseCommand {
	return NewBaseCommand("Cancel")
}

func NewGetMachineModel() *BaseCommand {
	return NewBaseCommand("GetMachineModel")
}

func NewIgnore() *BaseCommand {
	return NewBaseCommand("Ignore")
}

func NewSyncMachineModel() *BaseCommand {
	return NewBaseCommand("SyncMachineModel")
}
