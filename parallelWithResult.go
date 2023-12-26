package controlFlow

import (
	"errors"
	"fmt"
)

func NewParallelWithResultTask[T any](funcs ...func() *Result[T]) *ParallelWithResultTask[T] {
	return &ParallelWithResultTask[T]{
		funcs: funcs,
	}
}

type Result[T any] struct {
	Val T
	Err error
}

type ParallelWithResultTask[T any] struct {
	funcs []func() *Result[T]
}

func (pt *ParallelWithResultTask[T]) Add(funcs ...func() *Result[T]) {
	pt.funcs = append(pt.funcs, funcs...)
}

func (pt *ParallelWithResultTask[T]) Run() []*Result[T] {
	return ParallelWithResult[T](pt.funcs...)
}

func ParallelWithResult[T any](funcs ...func() *Result[T]) []*Result[T] {
	var ch = make(chan *Result[T], len(funcs))
	defer close(ch)
	for _, fn := range funcs {
		go func(f func() *Result[T]) {
			defer func() {
				if e := recover(); e != nil {
					ch <- &Result[T]{
						Err: errors.New(fmt.Sprintf("%v", e)),
					}
				}
			}()
			ch <- f()
		}(fn)
	}
	var results []*Result[T]
	for i := 0; i < len(funcs); i++ {
		results = append(results, <-ch)
	}
	return results
}
