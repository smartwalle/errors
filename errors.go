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
	Code    int         `json:"code"`
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (e *Error) Error() string {
	var buf bytes.Buffer
	buf.WriteString(strconv.Itoa(e.Code))
	buf.WriteByte('-')
	buf.WriteString(e.Message)
	return buf.String()
}
