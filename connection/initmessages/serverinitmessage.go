package initmessages

const (
	ExpectedServerVersion = 2
)

type ServerInitMessage struct {
	Version int64
	Id      int64
}
