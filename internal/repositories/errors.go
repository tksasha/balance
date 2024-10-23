package repositories

type NotFoundError struct {
	message string
}

func NewNotFoundError() error {
	return &NotFoundError{"not found"}
}

func (e *NotFoundError) Error() string {
	return e.message
}
