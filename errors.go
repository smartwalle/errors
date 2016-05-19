package errors

import (
	"fmt"
)

func New(text string) error {
	return NewWithCode(-1, text)
}

func NewWithCode(code int, text string) error {
	var err = &errorInfo{}
	err.Code    = code
	err.Message = text
	return err
}

func Code(err error) int {
	if e, ok := err.(*errorInfo); ok {
		return e.Code
	}
	return -1
}

func Message(err error) string {
	if e, ok := err.(*errorInfo); ok {
		return e.Message
	}
	return ""
}


type errorInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (this *errorInfo) Error() string {
	return fmt.Sprintf("[%d]%s", this.Code, this.Message)
}