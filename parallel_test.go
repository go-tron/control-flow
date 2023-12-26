package controlFlow

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func Parallel1() error {
	time.Sleep(time.Second * 1)
	return nil //errors.New("Error occurred in Parallel1")
}

func Parallel2() error {
	time.Sleep(time.Second * 1)
	return nil //errors.New("Error occurred in Parallel2")
}

func Parallel3() error {
	time.Sleep(time.Second * 1)
	return errors.New("Error occurred in Parallel3")
}

func TestParallel(t *testing.T) {
	err := Parallel(Parallel1, Parallel2, Parallel3)
	if err != nil {
		fmt.Printf("Parallel failed with error: %v\n", err)
	} else {
		fmt.Println("Parallel all succeed")
	}

	time.Sleep(time.Second * 1111)
}
