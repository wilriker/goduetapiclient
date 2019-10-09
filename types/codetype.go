package types

type CodeType string

const (
	Comment CodeType = "C"
	GCode            = "G"
	MCode            = "M"
	TCode            = "T"
)
