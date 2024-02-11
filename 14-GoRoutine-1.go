package main

import (
	"fmt"
	"time"
)

func main() {
	// there are 2 types of channels: Unbuffered channel and Buffered Channel
	// a channel in goland  will always going to block if it is full()
	resultCh := make(chan string) // -> unbuffered channel,
	//	resultCh <- "foo"             // -> is now FULL -> IT WILL BLOCK

	go func() {
		result := <-resultCh
		fmt.Println(result)
	}()

	resultCh <- "foo"

}

func fetchResource1(n int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("result %d", n)
}
