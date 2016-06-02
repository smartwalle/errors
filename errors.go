package errors

func New(text string) error {
	return NewWithCode(-1, text)
}

func NewWithCode(code int, text string) error {
	var err = &errorInfo{}
	err.Code = code
	err.Message = text
	return err
}

func ErrorCode(err error) int {
	if err == nil {
		return -1
	}
	if e, ok := err.(ErrorWithCode); ok {
		return e.ErrorCode()
	}
	return -1
}

func ErrorMessage(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

////////////////////////////////////////////////////////////////////////////////
type ErrorWithCode interface {
	ErrorCode() int
}

////////////////////////////////////////////////////////////////////////////////
type errorInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (this *errorInfo) ErrorCode() int {
	return this.Code
}

func (this *errorInfo) Error() string {
	return this.Message
}