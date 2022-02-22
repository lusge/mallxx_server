package merrorx

import (
	"encoding/json"
	"net/http"

	"google.golang.org/grpc/status"
)

const (
	defaultCode = 500
)

type CodeError struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
}

func ErrorHandler(err error) (int, interface{}) {

	if e, ok := err.(*CodeError); ok {
		return http.StatusOK, e
	} else {
		ret, ok := ConvertRpcError(err)
		if ok {
			return http.StatusOK, ret
		}

	}
	return http.StatusOK, &CodeError{
		Code:   defaultCode,
		Detail: err.Error(),
	}
}

func ConvertRpcError(err error) (*CodeError, bool) {
	st := status.Convert(err)
	var serviceErr CodeError
	jsonErr := json.Unmarshal([]byte(st.Message()), &serviceErr)

	if jsonErr != nil {
		return nil, false
	} else {
		return &CodeError{
			Code:   serviceErr.Code,
			Detail: serviceErr.Detail,
		}, true
	}
}

func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Detail: msg}
}

func NewDefaultError(msg string) error {
	return NewCodeError(defaultCode, msg)
}

func (e *CodeError) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}
