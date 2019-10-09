package types

type CodeChannel byte

const (
	HTTP CodeChannel = iota
	Telnet
	File
	USB
	AUX
	Daemon
	CodeQueue
	LCD
	SPI
	AutoPause
	DefaultChannel CodeChannel = SPI
)
