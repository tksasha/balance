package errors

type BadRequestError struct {
	message string
}

func NewBadRequestError() error {
	return &BadRequestError{
		message: "bad request",
	}
}

func (e *BadRequestError) Error() string {
	return e.message
}
