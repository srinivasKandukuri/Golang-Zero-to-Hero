package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

/*
you are given a slice (array) of functinos in go each functin has func error as singnature

exectue all the function concurrently using goroutines
collect the results error from each function execution

if a function completes successfully, store nil
if it returns an error, store the error

return a slic of error values where each elelemt corresponds to the results of the function at the same index in the original input slice

*/

func main() {
	tasks := []Task{
		func() error {
			time.Sleep(time.Millisecond * 100)
			return nil
		},
		func() error {
			time.Sleep(time.Millisecond * 500)
			return errors.New("task 2 error")
		},
		func() error {
			time.Sleep(time.Millisecond * 800)
			return nil
		},
	}
	results := RunAllTasks(tasks)

	for i, err := range results {
		if err != nil {
			fmt.Printf("Task %d error: %v\n", i, err)
		} else {
			fmt.Printf("Task %d completed successfully\n", i)
		}
	}
}

type Task func() error

func RunAllTasks(tasks []Task) []error {

	var wg sync.WaitGroup
	errs := make([]error, len(tasks))

	for i, task := range tasks {
		wg.Add(1)

		go func(idx int, f Task) {
			defer wg.Done()
			err := f()
			errs[idx] = err

		}(i, task)
	}
	wg.Wait()
	return errs
}
