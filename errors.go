package errors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"runtime"
)

func New(code int32, message string) *Error {
	var err = &Error{}
	err.Code = code
	err.Message = message
	return err
}

func Parse(s string) *Error {
	var e *Error
	if err := json.Unmarshal([]byte(s), &e); err != nil {
		return New(0, s)
	}
	return e
}

type Error struct {
	Code    int32       `json:"code"`
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
	buf.WriteString(fmt.Sprintf("%d", this.Code))
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
