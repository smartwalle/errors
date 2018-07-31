package errors

import (
	"fmt"
	"runtime"
	"bytes"
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
	Err     error  `json:"err,omitempty"`
	File    string `json:"file,omitempty"`
	Line    int    `json:"line,omitempty"`
	Func    string `json:"func,omitempty"`
}

func (this *Error) Error() string {
	var buf bytes.Buffer
	if this.File != "" {
		buf.WriteString(fmt.Sprintf("[%s - %s : %d] ", this.File, this.Func, this.Line))
	}
	buf.WriteString(this.Code)
	buf.WriteString(" - ")
	buf.WriteString(this.Message)
	if this.Err != nil {
		buf.WriteString(" (")
		buf.WriteString(this.Err.Error())
		buf.WriteString(")")
	}
	return buf.String()
}

func (this *Error) Location() *Error {
	var err = &Error{}
	err.Code = this.Code
	err.Message = this.Message
	err.Err = this.Err
	if showLocation {
		pc, file, line, ok := runtime.Caller(1)
		if ok == false {
			file = "???"
			line = -1
		}
		f := runtime.FuncForPC(pc)
		err.File = file
		err.Line = line
		err.Func = f.Name()
	}
	return err
}

func (this *Error) WithError(err error) *Error {
	var e = &Error{}
	e.Err = err
	e.Code = this.Code
	e.Message = this.Message
	e.File = this.File
	e.Line = this.Line
	e.Func = this.Func
	return e
}