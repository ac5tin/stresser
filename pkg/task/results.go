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
	// Total Transfer
	TotalTransfer float64
	// Throughput
	Throughput float64
	// RequestsPerSec
	RequestsPerSec float64
	// Task Responses
	responses []response
}

// Render the results
func (r *Results) Render() {
	fmt.Printf("\nTotal Duration: %v\nAvg. Duration %v\nMin. Duration %v\nMax Duration %v\nSuccess: %d\nFailed: %d\nTotal Transfer %.2f Mib\nThroughput: %.2f MiB/sec\nRequests/sec %2.f", r.Duration, r.AverageDuration, r.MinDuration, r.MaxDuration, r.SuccessCount, r.FailedCount, (r.TotalTransfer / 1000000), (r.Throughput / 1000000), r.RequestsPerSec)
}

// ExportResponsesToFile exports the responses to a CSV file
func (r *Results) ExportResponsesToFile(filepath string, onlyErr bool) error {
	// create file if not already exist, else overwrite

	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		file, err = os.Create(filepath)
		if err != nil {
			return err
		}
	}
	// skip errors
	exp := r.responses
	if onlyErr {
		exp = make([]response, 0)
		for _, resp := range r.responses {
			if resp.Error != nil {
				exp = append(exp, resp)
			}
		}
	}

	if err := gocsv.MarshalFile(&exp, file); err != nil {
		return err
	}
	return nil
}
