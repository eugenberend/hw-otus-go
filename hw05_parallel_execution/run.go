package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

var ErrWorkersNumberNegative = errors.New("number of workers should be positive")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var ignoreErr bool
	var result error

	if m < 0 {
		ignoreErr = true
	}

	if n <= 0 {
		return ErrWorkersNumberNegative
	}

	if len(tasks) < n {
		n = len(tasks)
	}

	wg := new(sync.WaitGroup)
	taskCh := make(chan Task)
	var errCounter int32

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for task := range taskCh {
				err := task()
				if !ignoreErr && err != nil {
					atomic.AddInt32(&errCounter, 1)
				}
			}
		}()
	}

	for _, task := range tasks {
		taskCh <- task
		curErrCount := int(atomic.LoadInt32(&errCounter))
		if !ignoreErr && curErrCount >= m {
			result = ErrErrorsLimitExceeded
			break
		}
	}
	close(taskCh)
	wg.Wait()

	return result
}
