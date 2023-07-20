package error

type ErrorBusinessException struct {
	Err error
}

func (e *ErrorBusinessException) Error() string {
	return e.Err.Error()
}

func NewErrorBusinessException(err error) error {
	return &ErrorBusinessException{Err: err}
}
