package commands

type Response interface {
	IsSuccess() bool
	GetResult() interface{}
	GetErrorType() string
	GetErrorMessage() string
}

type BaseResponse struct {
	Success bool
	// }

	// type Response struct {
	// BaseResponse
	Result interface{}
	// }

	// type ErrorResponse struct {
	// BaseResponse
	ErrorType    string
	ErrorMessage string
}

func (br *BaseResponse) IsSuccess() bool {
	return br.Success
}

func (br *BaseResponse) GetResult() interface{} {
	return br.Result
}

func (br *BaseResponse) GetErrorType() string {
	return br.ErrorType
}

func (br *BaseResponse) GetErrorMessage() string {
	return br.ErrorMessage
}
