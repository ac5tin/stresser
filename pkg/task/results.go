package task

import (
	"fmt"
	"os"
	"time"

	"github.com/gocarina/gocsv"
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
	// Task Responses
	responses []response
}

// Render the results
func (r *Results) Render() {
	fmt.Printf("\nTotal Duration: %v\nAvg. Duration %v\nMin. Duration %v\nMax Duration %v\nSuccess: %d\nFailed: %d", r.Duration, r.AverageDuration, r.MinDuration, r.MaxDuration, r.SuccessCount, r.FailedCount)
}

// ExportResponsesToFile exports the responses to a CSV file
func (r *Results) ExportResponsesToFile(filepath string) error {
	// create file if not already exist
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		file, err = os.Create(filepath)
		if err != nil {
			return err
		}
	}
	if err := gocsv.MarshalFile(&r.responses, file); err != nil {
		return err
	}
	return nil
}
