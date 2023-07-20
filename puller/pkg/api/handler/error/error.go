package error

type ErrorBadRequest struct {
	Err error
}

func (e *ErrorBadRequest) Error() string {
	return e.Err.Error()
}

func NewErrorBadRequest(err error) error {
	return &ErrorBadRequest{Err: err}
}
