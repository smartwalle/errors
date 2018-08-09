package errors

import (
	"bytes"
	"fmt"
	"runtime"
)

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
		nErr = &Error{}
		nErr.Code = e.Code
		nErr.Message = e.Message
		nErr.Err = e.Err
		nErr.File = e.File
		nErr.Line = e.Line
		nErr.Func = e.Func
		nErr.Data = e.Data
	case nil:
		nErr = nil
	default:
		nErr = New(e.Error())
	}
	return nErr
}

func WithData(data interface{}) *Error {
	var nErr = &Error{}
	nErr.Code = "0"
	nErr.Data = data
	return nErr
}

type Error struct {
	Code    string      `json:"code"`
	Message string      `json:"message,omitempty"`
	Err     error       `json:"err,omitempty"`
	File    string      `json:"file,omitempty"`
	Line    int         `json:"line,omitempty"`
	Func    string      `json:"func,omitempty"`
	Data    interface{} `json:"data,omitempty"`
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
		buf.WriteString(" {")
		buf.WriteString(this.Err.Error())
		buf.WriteString("}")
	}
	return buf.String()
}

func (this *Error) Format(args ...interface{}) *Error {
	var nErr = &Error{}
	nErr.Code = this.Code
	nErr.Message = fmt.Sprintf(this.Message, args...)
	nErr.Err = this.Err
	nErr.File = this.File
	nErr.Line = this.Line
	nErr.Func = this.Func
	nErr.Data = this.Data
	return nErr
}

func (this *Error) Location() *Error {
	var nErr = &Error{}
	nErr.Code = this.Code
	nErr.Message = this.Message
	nErr.Err = this.Err

	pc, file, line, ok := runtime.Caller(1)
	if ok == false {
		file = "???"
		line = -1
	}
	f := runtime.FuncForPC(pc)
	nErr.File = file
	nErr.Line = line
	nErr.Func = f.Name()
	nErr.Data = this.Data
	return nErr
}

func (this *Error) WithError(err error) *Error {
	var nErr = &Error{}
	nErr.Code = this.Code
	nErr.Message = this.Message
	nErr.Err = err
	nErr.File = this.File
	nErr.Line = this.Line
	nErr.Func = this.Func
	nErr.Data = this.Data
	return nErr
}

func (this *Error) WithData(data interface{}) *Error {
	var nErr = &Error{}
	nErr.Code = this.Code
	nErr.Message = this.Message
	nErr.Err = this.Err
	nErr.File = this.File
	nErr.Line = this.Line
	nErr.Func = this.Func
	nErr.Data = data
	return nErr
}
