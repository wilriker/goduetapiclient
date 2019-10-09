package initmessages

type SubscriptionMode string

const (
	SubscriptionModeFull  SubscriptionMode = "Full"
	SubscriptionModePatch                  = "Patch"
)

type SubscribeInitMessage struct {
	BaseInitMessage
	SubscriptionMode SubscriptionMode
	Filter           string
}

func NewSubscribeInitMessage(subMode SubscriptionMode, filter string) ClientInitMessage {
	return &SubscribeInitMessage{
		BaseInitMessage:  BaseInitMessage{Mode: ConnectionModeSubscribe},
		SubscriptionMode: subMode,
		Filter:           filter,
	}
}
