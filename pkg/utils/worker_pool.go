package utils

import (
	"context"
	"sync"

	"crypto-dashboard/pkg/response"

	"golang.org/x/sync/errgroup"
)

type (
	TaskIf[T, V any] interface {
		Process() (V, *response.AppError)
	}

	WorkerTask[T any, V any] struct {
		Concurency        int
		StopWhenErrorFlag bool
		WaitFlag          bool
		mu                sync.Mutex
		data              []V
		TaskChain         chan TaskIf[T, V]
	}
)

func (wt *WorkerTask[T, V]) Run(ctx context.Context) (*[]V, *response.AppError) {
	eg, ctx := errgroup.WithContext(ctx)
	eg.SetLimit(wt.getConcurrency())
	defer close(wt.TaskChain)

	process := wt.processUnskipError
	if !wt.StopWhenErrorFlag {
		process = wt.processSkipError
	}

	if wt.WaitFlag {
		process(ctx, eg)
		err := eg.Wait()
		return &wt.data, response.ConvertError(err)
	}
	go process(ctx, eg)

	return nil, nil
}

func (wt *WorkerTask[T, V]) getConcurrency() int {
	if wt.Concurency > 100 || wt.Concurency == 0 {
		return 10
	}
	return wt.Concurency
}

func (wt *WorkerTask[T, V]) processSkipError(_ context.Context, g *errgroup.Group) {
	for task := range wt.TaskChain {
		g.Go(func() error {
			res, err := task.Process()
			wt.atomicAppend(res)
			if err != nil {
				return err
			}
			return nil
		})
	}
}

func (wt *WorkerTask[T, V]) atomicAppend(res V) {
	wt.mu.Lock()
	wt.data = append(wt.data, res)
	wt.mu.Unlock()
}

func (wt *WorkerTask[T, V]) processUnskipError(ctx context.Context, g *errgroup.Group) {
	for task := range wt.TaskChain {
		g.Go(func() error {
			select {
			case <-ctx.Done():
				return nil
			default:
				res, err := task.Process()
				wt.atomicAppend(res)
				if err != nil {
					return err
				}
				return nil
			}
		})
	}
}
