package exception

type NotFoundError struct {
	Err  string
	Type string
	Meta string
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{
		Err:  error,
		Type: "NotFoundError",
		Meta: "NotFoundError",
	}
}

func (e NotFoundError) Error() string {
	return e.Err
}
