package task

import "time"

// response is the response from a task
type response struct {
	StatusCode uint32        `csv:"status_code"`
	Duration   time.Duration `csv:"duration"`
	Error      error         `csv:"error"`
	Body       string        `csv:"body"`
}
