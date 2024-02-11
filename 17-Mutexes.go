package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type State struct {
	mu    sync.Mutex
	count int32
}

func (s *State) setState(i int32) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count = i
}

// or using sync/atomic
func (s *State) setState1(i int32) {
	atomic.AddInt32(&s.count, i)
}

func main() {
	state := &State{}
	for i := 0; i < 10; i++ {
		go func(i int32) {
			state.setState(i + 1)
		}(int32(i))
	}

	for i := 0; i < 10; i++ {
		go func(i int32) {
			state.setState1(i)
		}(int32(i))
	}

	fmt.Printf("%+v\n", state)
}
