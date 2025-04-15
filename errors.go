package errors

import (
	"bytes"
	"strconv"
)

func New(code int, message string) *Error {
	var err = &Error{}
	err.Code = code
	err.Message = message
	return err
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	var buf bytes.Buffer
	buf.WriteString(strconv.Itoa(e.Code))
	buf.WriteByte('-')
	buf.WriteString(e.Message)
	return buf.String()
}

func (e *Error) GetCode() int {
	return e.Code
}

func (e *Error) GetMessage() string {
	return e.Message
}
