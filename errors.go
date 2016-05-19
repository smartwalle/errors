package errors

func New(text string) error {
	return NewWithCode(-1, text)
}

func NewWithCode(code int, text string) error {
	var err = &errorInfo{}
	err.code    = code
	err.message = text
	return err
}

func Code(err error) int {
	if e, ok := err.(errorWithCode); ok {
		return e.Code()
	}
	return -1
}

func Message(err error) string {
	return err.Error()
}

////////////////////////////////////////////////////////////////////////////////
type errorWithCode interface {
	Code() int
}

////////////////////////////////////////////////////////////////////////////////
type errorInfo struct {
	code    int    `json:"code"`
	message string `json:"message"`
}

func (this *errorInfo) Code() int {
	return this.code
}

func (this *errorInfo) Error() string {
	return this.message
}