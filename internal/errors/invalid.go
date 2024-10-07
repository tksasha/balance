package errors

type InvalidError struct {
	message string
}

func NewInvalidError() error {
	return &InvalidError{
		message: "is invalid",
	}
}

func (e *InvalidError) Error() string {
	return e.message
}
