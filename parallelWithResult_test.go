package controlFlow

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

type A string

func ParallelWithResult1() *Result[A] {
	time.Sleep(time.Second * 1)
	return &Result[A]{
		Val: "ParallelWithResult1 executed successfully",
		Err: nil,
	}
}

func ParallelWithResult2() *Result[A] {
	time.Sleep(time.Second * 2)
	return &Result[A]{
		Val: "ParallelWithResult2 executed successfully",
		Err: errors.New("Error occurred in ParallelWithResult2"),
	}
}

func ParallelWithResult3() *Result[A] {
	time.Sleep(time.Second * 1)
	panic("panic 3")
	return &Result[A]{
		Val: "ParallelWithResult3 executed successfully",
		Err: nil,
	}
}

func TestParallelWithResult(t *testing.T) {
	results := ParallelWithResult(ParallelWithResult1, ParallelWithResult2, ParallelWithResult3)
	for i, result := range results {
		if result.Err != nil {
			fmt.Printf("Function %d failed with error: %v\n", i+1, result.Err)
		} else {
			fmt.Printf("Function %d result: %v\n", i+1, result.Val)
		}
	}
	time.Sleep(time.Second * 1111)
}
