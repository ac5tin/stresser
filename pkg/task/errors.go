package task

import "errors"

var (
	ErrNotImplemented              = errors.New("not implemented")
	UnacceptableStatusCode         = errors.New("unacceptable status code")
	ErrConcurrencyGreaterThanTotal = errors.New("concurrency is greater than total")
	ErrInvalidConcurrencyNumber    = errors.New("concurrency must be greater than 0")
	ErrInvalidTotalNumber          = errors.New("total must be greater than 0")
)
