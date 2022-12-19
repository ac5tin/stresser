package task

import (
	"context"
	"sync/atomic"
	"time"

	"golang.org/x/sync/errgroup"
)

// Execute task
func (t *Task) Execute(cfg *Config) (*Results, error) {
	success := uint32(0)
	fails := uint32(0)
	durations := make([]time.Duration, 0, cfg.Concurrency)
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
		Duration:     uint32(totalDuration),
		FailedCount:  fails,
		SuccessCount: success,
	}

	return &r, nil
}

// exec executes the task
func (t *Task) exec() error {
	return ErrNotImplemented
}
