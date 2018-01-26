package errors

import "fmt"

// New 最多支持 2 个参数
// 当只有一个参数的时候，默认给 message 赋值
// 当有两个参数的时候，第一个参数为 code， 第二个参数为 message
func New(args ...string) *Error {
	argsLen := len(args)
	err := &Error{}
	if argsLen == 1 {
		err.Code = "0"
		err.Message = args[0]
	}
	if argsLen >= 2 {
		err.Code = args[0]
		err.Message = args[1]
	}
	return err
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message,omitempty"`
}

func (this *Error) Error() string {
	if this.Code != "" {
		fmt.Sprintf("%s - %s", this.Code, this.Message)
	}
	return this.Message
}
