package commands

type ReceivedHttpRequest struct {
	SessionId   int64
	Queries     map[string]string
	Headers     map[string]string
	ContentType string
	Body        string
}

func NewReceivedHttpRequest() *ReceivedHttpRequest {
	return &ReceivedHttpRequest{}
}
