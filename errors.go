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
	if e, ok := err.(errorWithCode); ok {
		return e.ErrorCode()
	}
	return -1
}

func ErrorMessage(err error) string {
	return err.Error()
}

////////////////////////////////////////////////////////////////////////////////
type errorWithCode interface {
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