package main

import (
	"context"
	"errors"
	"sync"
)

type Result[R any] struct {
	Value R
	Err   error
}

var ErrInvalidWorkers = errors.New("workers must be positive")

func ParallelMap[T any, R any](ctx context.Context, workers int, in []T, fn func(context.Context, T) (R, error)) ([]Result[R], error) {
	if workers <= 0 {
		return nil, ErrInvalidWorkers
	}

	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	if len(in) == 0 {
		return []Result[R]{}, nil
	}

	results := make([]Result[R], len(in))
	numWorkers := workers
	if numWorkers > len(in) {
		numWorkers = len(in)
	}

	jobs := make(chan int)
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for w := 0; w < numWorkers; w++ {
		go func() {
			defer wg.Done()
			for idx := range jobs {
				val, err := fn(ctx, in[idx])
				if err != nil {
					var zero R
					results[idx] = Result[R]{Value: zero, Err: err}
				} else {
					results[idx] = Result[R]{Value: val, Err: nil}
				}
			}
		}()
	}

	go func() {
		defer close(jobs)
		for i := 0; i < len(in); i++ {
			select {
			case <-ctx.Done():
				return
			case jobs <- i:
			}
		}
	}()
	wg.Wait()
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	return results, nil
}
