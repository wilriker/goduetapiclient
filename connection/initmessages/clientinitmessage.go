package initmessages

type ConnectionMode string

const (
	ConnectionModeUnknown   ConnectionMode = "Unknown"
	ConnectionModeCommand                  = "Command"
	ConnectionModeIntercept                = "Intercept"
	ConnectionModeSubscribe                = "Subscribe"
)

type ClientInitMessage interface {
	GetMode() ConnectionMode
}

type BaseInitMessage struct {
	Mode ConnectionMode
}

func (bim *BaseInitMessage) GetMode() ConnectionMode {
	return bim.Mode
}

func NewCommandInitMessage() ClientInitMessage {
	return &BaseInitMessage{Mode: ConnectionModeCommand}
}
