package task

import (
	"bytes"
	"context"
	"net/http"
	"sync/atomic"
	"time"

	"golang.org/x/sync/errgroup"
)

// Execute task
func (t *Task) Execute(cfg *Config) (*Results, error) {
	success := uint32(0)
	fails := uint32(0)
	durations := make([]time.Duration, cfg.Total, cfg.Total)
	totalStart := time.Now()

	// concurrently run task
	g, ctx := errgroup.WithContext(context.Background())
	g.SetLimit(int(cfg.Concurrency))
	for i := 0; i < int(cfg.Total); i++ {
		i := i
		g.Go(func() error {
			select {
			case <-ctx.Done():
				// something went wrong in the other goroutine
				return nil
			default:
				// dont block
			}
			start := time.Now()
			if err := t.exec(); err != nil {
				atomic.AddUint32(&fails, 1)
				return nil
			}
			atomic.AddUint32(&success, 1)
			durations[i] = time.Since(start)
			return nil
		})

	}
	if err := g.Wait(); err != nil {
		return nil, err
	}
	totalDuration := time.Since(totalStart)

	// calculate results
	r := Results{
		Duration:        totalDuration,
		FailedCount:     fails,
		SuccessCount:    success,
		MinDuration:     durations[0],
		MaxDuration:     durations[0],
		AverageDuration: totalDuration / time.Duration(cfg.Total),
	}

	for _, d := range durations {
		if d == 0 {
			continue
		}
		if d > r.MaxDuration {
			r.MaxDuration = d
		}
		if d < r.MinDuration {
			r.MinDuration = d
		}
	}

	return &r, nil
}

// exec executes the task
func (t *Task) exec() error {
	req, err := http.NewRequest(t.Method, t.URL, bytes.NewBuffer(t.Payload))
	if err != nil {
		return err
	}
	for k, v := range t.Headers {
		req.Header.Set(k, v)
	}
	client := &http.Client{
		Timeout: time.Second * time.Duration(t.Timeout),
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// ensure statuscode is in AcceptedStatusCodes
	for _, code := range t.AcceptedStatusCodes {
		if resp.StatusCode == int(code) {
			return nil
		}
	}

	return UnacceptableStatusCode
}
