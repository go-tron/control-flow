package controlFlow

import (
	"errors"
	"fmt"
	"sync"
)

func NewParallelTask(funcs ...func() error) *ParallelTask {
	return &ParallelTask{
		funcs: funcs,
	}
}

type ParallelTask struct {
	funcs []func() error
}

func (pt *ParallelTask) Add(funcs ...func() error) {
	pt.funcs = append(pt.funcs, funcs...)
}

func (pt *ParallelTask) Run() error {
	return Parallel(pt.funcs...)
}

func (pt *ParallelTask) RunBreakOnError() error {
	return ParallelBreakOnError(pt.funcs...)
}

func Parallel(funcs ...func() error) error {
	var ch = make(chan error, len(funcs))
	defer close(ch)
	for _, fn := range funcs {
		go func(f func() error) {
			defer func() {
				if e := recover(); e != nil {
					ch <- errors.New(fmt.Sprintf("%v", e))
				}
			}()
			ch <- f()
		}(fn)
	}
	var err error
	for i := 0; i < len(funcs); i++ {
		if e := <-ch; e != nil {
			err = e
		}
	}
	return err
}

func ParallelBreakOnError(funcs ...func() error) error {
	if len(funcs) == 0 {
		return nil
	}
	var wg sync.WaitGroup
	ch := make(chan error, len(funcs))

	for _, fn := range funcs {
		wg.Add(1)
		go func(f func() error) {
			defer wg.Done()
			defer func() {
				if e := recover(); e != nil {
					ch <- errors.New(fmt.Sprintf("%v", e))
				}
			}()
			if err := f(); err != nil {
				ch <- err
			}
		}(fn)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for err := range ch {
		return err
	}
	return nil
}
