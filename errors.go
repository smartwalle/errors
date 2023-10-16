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
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	File    string      `json:"file,omitempty"`
	Func    string      `json:"func,omitempty"`
	Line    int         `json:"line,omitempty"`
	Code    int32       `json:"code"`
}

func (e *Error) Error() string {
	var buf bytes.Buffer
	//if e.File != "" {
	//	buf.WriteString(fmt.Sprintf("[%s - %s : %d] ", e.File, e.Func, e.Line))
	//}
	buf.WriteString(fmt.Sprintf("%d", e.Code))
	buf.WriteString(sep)
	buf.WriteString(e.Message)
	return buf.String()

	//bytes, _ := json.Marshal(e)
	//return string(bytes)
}

func (e Error) Format(args ...interface{}) *Error {
	e.Message = fmt.Sprintf(e.Message, args...)
	return &e
}

func (e Error) Location() *Error {
	pc, file, line, ok := runtime.Caller(1)
	if ok == false {
		file = "???"
		line = -1
	}
	f := runtime.FuncForPC(pc)
	e.File = file
	e.Line = line
	e.Func = f.Name()
	return &e
}

func (e Error) WithData(data interface{}) *Error {
	e.Data = data
	return &e
}

func (e Error) WithMessage(msg string) *Error {
	e.Message = msg
	return &e
}
