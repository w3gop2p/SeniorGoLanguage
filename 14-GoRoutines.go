package main

import (
	"fmt"
	"time"
)

// a go routine is a function that is getting scheduled by the go scheduler at the run time
// goroutines don't run in parallel, it is about concurrency

func main() {
	/*result := fetchResource()
	fmt.Println("the result:", result)*/

	// 1 st way to call a goroutine
	//	go fetchResource(1) // this is an async call

	// 2 nd way to call a goroutine
	go func() {
		result := fetchResource(2)
		fmt.Println(result)
	}()
}

func fetchResource(n int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("result %d", n)
}
