package types

type MessageType int64

const (
	Success MessageType = iota
	Warning
	Error
)
