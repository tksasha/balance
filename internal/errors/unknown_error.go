package errors

type UnknownError struct {
	err error
}

func NewUnknownError(err error) error {
	return &UnknownError{
		err: err,
	}
}

func (e *UnknownError) Error() string {
	return e.err.Error()
}
