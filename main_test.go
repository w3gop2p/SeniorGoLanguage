package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestCalculateValues(t *testing.T) {
	fmt.Println("Hello from our test")
}

func TestState(t *testing.T) {
	state := State{}
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		// doing some work
		go func(i int32) {
			state.setState(i + 1)
			wg.Done() // done working
		}(int32(i))
	}
	wg.Wait()
	fmt.Printf("%+v\n", state)
}
