package exception

type NotFoundError struct {
	Err string
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{
		Err: error,
	}
}

func (e NotFoundError) Error() string {
	return e.Err
}
