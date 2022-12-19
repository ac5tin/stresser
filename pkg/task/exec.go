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
	// ensure config is valid
	{
		if cfg.Concurrency == 0 {
			return nil, ErrInvalidConcurrencyNumber
		}
		if cfg.Total == 0 {
			return nil, ErrInvalidTotalNumber
		}
		if cfg.Concurrency > uint16(cfg.Total) {
			return nil, ErrConcurrencyGreaterThanTotal
		}
	}

	success := uint32(0)
	fails := uint32(0)
	durations := make([]time.Duration, cfg.Total, cfg.Total)
	totalStart := time.Now()
	responses := make([]response, cfg.Total, cfg.Total)

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
			resp, err := t.exec()
			responses[i] = *resp
			if err != nil {
				atomic.AddUint32(&fails, 1)
				return nil
			}
			atomic.AddUint32(&success, 1)
			durations[i] = resp.Duration
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
		responses:       responses,
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
func (t *Task) exec() (*response, error) {
	startTime := time.Now()
	res := response{}
	req, err := http.NewRequest(t.Method, t.URL, bytes.NewBuffer(t.Payload))
	if err != nil {
		res.Error = err
		return &res, err
	}
	for k, v := range t.Headers {
		req.Header.Set(k, v)
	}
	client := &http.Client{
		Timeout: time.Second * time.Duration(t.Timeout),
	}
	resp, err := client.Do(req)
	if err != nil {
		res.Error = err
		return &res, err
	}
	defer resp.Body.Close()

	// fill in response
	{
		// duration
		res.Duration = time.Since(startTime)
		// status code
		res.StatusCode = uint32(resp.StatusCode)
		// body
		buf := bytes.NewBuffer(nil)
		buf.ReadFrom(resp.Body)
		res.Body = buf.String()
	}

	// ensure statuscode is in AcceptedStatusCodes
	for _, code := range t.AcceptedStatusCodes {
		if resp.StatusCode == int(code) {
			return &res, nil
		}
	}

	return &res, UnacceptableStatusCode
}
