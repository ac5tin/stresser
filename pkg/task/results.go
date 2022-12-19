package task

import "time"

// Results contain the results of a task execution
type Results struct {
	// Total time taken to execute the task
	Duration time.Duration
	// Total failed requests
	FailedCount uint32
	// Total successful requests
	SuccessCount uint32
	// Minimum time taken to execute a request
	MinDuration time.Duration
	// Maximum time taken to execute a request
	MaxDuration time.Duration
	// Average time taken to execute a request
	AverageDuration time.Duration
}
