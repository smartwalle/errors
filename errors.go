package errors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

const (
	sep = " - "
)

func New(code int32, message string) *Error {
	var err = &Error{}
	err.Code = code
	err.Message = message
	return err
}

func Parse(s string) *Error {
	var e *Error

	var bytes = []byte(s)

	if bytes[0] == '{' {
		if err := json.Unmarshal(bytes, &e); err != nil {
			return New(0, s)
		}
	} else {
		var ss = strings.SplitN(s, sep, 2)
		if len(ss) > 1 {
			var code, _ = strconv.Atoi(ss[0])
			return New(int32(code), ss[1])
		}
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
	//if this.File != "" {
	//	buf.WriteString(fmt.Sprintf("[%s - %s : %d] ", this.File, this.Func, this.Line))
	//}
	buf.WriteString(fmt.Sprintf("%d", this.Code))
	buf.WriteString(sep)
	buf.WriteString(this.Message)
	return buf.String()

	//bytes, _ := json.Marshal(this)
	//return string(bytes)
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

func (this Error) WithMessage(msg string) *Error {
	this.Message = msg
	return &this
}
