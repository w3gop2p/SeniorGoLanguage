package main

import (
	"fmt"
	"time"
)

func main() {
	msgch := make(chan string, 128)
	msgch <- "A"
	msgch <- "B"
	msgch <- "C"
	close(msgch)

	for {
		msg, ok := <-msgch
		if !ok {
			break
		}
		fmt.Println("the message ->", msg)
	}

	/*for msg := range msgch {
		fmt.Println(msg)
	}*/
	fmt.Println("done reading all the messages from the channel")
}

func fetchResource2(n int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("result %d", n)
}
