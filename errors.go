package errors

import (
	"fmt"
	"runtime"
)

var (
	showLocation = true
)

func ShowLocation() {
	showLocation = true
}

func HideLocation() {
	showLocation = false
}

// New 最多支持 2 个参数
// 当只有一个参数的时候，默认给 message 赋值
// 当有两个参数的时候，第一个参数为 code， 第二个参数为 message
func New(args ...string) *Error {
	argsLen := len(args)
	err := &Error{}
	if argsLen == 1 {
		err.Code = "0"
		err.Message = args[0]
	} else if argsLen >= 2 {
		err.Code = args[0]
		err.Message = args[1]
	}
	return err
}

func WithError(err error) *Error {
	var nErr *Error
	switch e := err.(type) {
	case *Error:
		nErr = New(e.Code, e.Message)
		nErr.File = e.File
		nErr.Line = e.Line
	case nil:
		nErr = nil
	default:
		nErr = New(e.Error())
	}
	return nErr
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message,omitempty"`
	File    string `json:"file,omitempty"`
	Line    int    `json:"line,omitempty"`
}

func (this *Error) Error() string {
	if this.File != "" {
		return fmt.Sprintf("[%s - %d] %s - %s", this.File, this.Line, this.Code, this.Message)
	}
	return fmt.Sprintf("%s - %s", this.Code, this.Message)
}

func (this *Error) Location() *Error {
	if showLocation {
		_, file, line, ok := runtime.Caller(1)
		if ok == false {
			file = "???"
			line = -1
		}
		this.File = file
		this.Line = line
	}
	return this
}
