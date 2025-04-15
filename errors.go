package errors

import (
	"bytes"
	"strconv"
)

func New(code int, message string) *Error {
	var err = &Error{}
	err.code = code
	err.message = message
	return err
}

type Error struct {
	code    int
	message string
}

func (e *Error) Error() string {
	var buf bytes.Buffer
	buf.WriteString(strconv.Itoa(e.code))
	buf.WriteByte('-')
	buf.WriteString(e.message)
	return buf.String()
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Message() string {
	return e.message
}
