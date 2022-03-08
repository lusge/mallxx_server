package merrorx

import (
	"encoding/json"
	"fmt"
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

type ResultError struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
}

func ErrorHandler(err error) (int, interface{}) {

	if e, ok := err.(*CodeError); ok {
		return http.StatusOK, &ResultError{
			Code:   e.Code,
			Detail: e.Detail,
		}
	} else {
		ret, ok := ConvertRpcError(err)
		if ok {
			fmt.Println(ret, "bbbb")
			return http.StatusOK, &ResultError{
				Code:   ret.Code,
				Detail: ret.Detail,
			}
		}

	}
	return http.StatusOK, &ResultError{
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
