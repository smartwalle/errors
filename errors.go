package errors

import (
	"bytes"
	"encoding/json"
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

func Parse(s string) *Error {
	var e *Error
	if err := json.Unmarshal([]byte(s), &e); err != nil {
		return New(s)
	}
	return e
}

type Error struct {
	Code    string      `json:"code"`
	Message string      `json:"message,omitempty"`
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
	return buf.String()
}

func (this Error) Format(args ...interface{}) *Error {
	this.Message = fmt.Sprintf(this.Message, args...)
	return &this
}

func (this Error) Location() *Error {
	pc, file, line, ok := runtime.Caller(1)
	if ok == false {
		file = "???"
		line = -1
	}
	f := runtime.FuncForPC(pc)
	this.File = file
	this.Line = line
	this.Func = f.Name()
	return &this
}

func (this Error) WithData(data interface{}) *Error {
	this.Data = data
	return &this
}
