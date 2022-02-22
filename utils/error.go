package utils

import (
	"encoding/json"
	"errors"
	"fmt"
)

type ServiceError struct {
	Id     string `json:"id"`
	Code   int    `json:"code"`
	Detail string `json:"detail"`
}

func NewServiceError(id string, code int, detail string) error {
	serr := ServiceError{
		Id:     id,
		Code:   code,
		Detail: detail,
	}

	str, _ := json.Marshal(&serr)
	fmt.Println(string(str))
	return errors.New(string(str))
}
