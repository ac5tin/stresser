package task

import "time"

// response is the response from a task
type response struct {
	StatusCode uint32        `csv:"status_code"`
	Duration   time.Duration `csv:"duration"`
	Error      error         `csv:"-"`
	Body       responseBody  `csv:"body"`
}

type responseBody []byte

func (rb *responseBody) MarshalCSV() (string, error) {
	return string(*rb), nil
}
