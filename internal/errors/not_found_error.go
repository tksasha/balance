package errors

type NotFoundError struct {
	err error
}

func NewNotFoundError(err error) error {
	return &NotFoundError{
		err: err,
	}
}

func (e *NotFoundError) Error() string {
	return e.err.Error()
}
