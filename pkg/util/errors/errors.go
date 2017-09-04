package errors

type notFoundError struct {
	err string
}

func (n *notFoundError) Error() string {
	return n.err
}

func NewNotFound(s string) error {
	return &notFoundError{s}
}

func IsNotFound(err error) bool {
	if _, ok := err.(*notFoundError); ok {
		return true
	}
	return false
}
