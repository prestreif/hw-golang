package hw05parallelexecution

import (
	"errors"
	"sync"
)

var (
	ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
	ErrWorkersLessThenOne  = errors.New("workers less then one")
	ErrTasksLessThenOne    = errors.New("count tasks less then one")
)

type Task func() error

// struct for safe count error
type CountErrors struct {
	countErrors int
	mutex       sync.Mutex
}

func (t *CountErrors) MoreThen(m int) bool {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.countErrors > m
}

func (t *CountErrors) AddCountError() {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.countErrors++
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	if n < 1 {
		return ErrWorkersLessThenOne
	}

	if len(tasks) < 1 {
		return ErrTasksLessThenOne
	}

	// for count error
	countErr := CountErrors{}

	// for flow tasks
	taskCh := make(chan Task, len(tasks))
	// sync groups
	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			// minus from wg when finished goroutine
			defer wg.Done()
			// get task from flow
			for task := range taskCh {
				// if countError More then max error
				if countErr.MoreThen(m) {
					break
				}
				// if execute task get any error
				if task() != nil {
					countErr.AddCountError()
				}
			}
		}()
	}

	// fill of flow tasks
	for i := range tasks {
		taskCh <- tasks[i]
	}
	// close chan so we can finished read from flow
	close(taskCh)
	// wite all goroutines
	wg.Wait()

	if countErr.MoreThen(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
