package commands

import (
	"github.com/wilriker/goduetapiclient/types"
)

// SendHttpResponse responds to a received HTTP request
type SendHttpResponse struct {
	//
	StatusCode   uint16
	Response     string
	ResponseType types.HttpResponseType
}

func NewSendHttpResponse(statusCode uint16, response string, t types.HttpResponseType) *SendHttpResponse {
	return &SendHttpResponse{
		StatusCode:   statusCode,
		Response:     response,
		ResponseType: t,
	}
}
