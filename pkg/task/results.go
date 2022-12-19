package task

import (
	"fmt"
	"time"
)

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

// Render the results
func (r *Results) Render() {
	fmt.Printf("\nTotal Duration: %v\nAvg. Duration %v\nMin. Duration %v\nMax Duration %v\nSuccess: %d\nFailed: %d", r.Duration, r.AverageDuration, r.MinDuration, r.MaxDuration, r.SuccessCount, r.FailedCount)
}
