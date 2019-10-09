package initmessages

type InterceptionMode string

const (
	InterceptionModePre      InterceptionMode = "Pre"
	InterceptionModePost                      = "Post"
	InterceptionModeExecuted                  = "Executed"
)

type InterceptInitMessage struct {
	BaseInitMessage
	InterceptionMode InterceptionMode
}

func NewInterceptInitMessage(iMode InterceptionMode) ClientInitMessage {
	return &InterceptInitMessage{
		BaseInitMessage:  BaseInitMessage{Mode: ConnectionModeIntercept},
		InterceptionMode: iMode,
	}
}
