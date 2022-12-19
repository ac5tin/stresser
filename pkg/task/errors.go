package task

import "errors"

var (
	ErrNotImplemented      = errors.New("not implemented")
	UnacceptableStatusCode = errors.New("unacceptable status code")
)
