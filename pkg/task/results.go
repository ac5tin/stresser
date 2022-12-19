package task

// Results contain the results of a task execution
type Results struct {
	// Total time taken to execute the task
	Duration uint32
	// Total failed requests
	FailedCount uint32
	// Total successful requests
	SuccessCount uint32
	// Minimum time taken to execute a request
	MinDuration uint32
	// Maximum time taken to execute a request
	MaxDuration uint32
	// Average time taken to execute a request
	AverageDuration uint32
}
