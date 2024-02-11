package main

import (
	"fmt"
	"time"
)

type Server3 struct {
	quitch chan struct{} // 0 bytes
	msgch  chan string
}

func newServer1() *Server3 {
	return &Server3{
		quitch: make(chan struct{}),
		msgch:  make(chan string, 128),
	}
}

func (s *Server3) start() {
	fmt.Println("server starting")
	s.loop() // block
}

func (s *Server3) sendMessage1(msg string) {
	s.msgch <- msg
}

func (s *Server3) quit() {
	close(s.quitch)
	// s.quitch <- struct{}{}    // it is the same
}

func (s *Server3) loop() {
mainloop:
	for {
		select {
		case <-s.quitch:
			// do some stuff when we need to quit
			fmt.Println("quiting server")
			break mainloop
		case msg := <-s.msgch:
			// do some stuff when we have a message
			s.handleMessage1(msg)
		default:

		}
	}
	fmt.Println("server is shuting down gracefully")
}

func (s *Server3) handleMessage1(msg string) {
	fmt.Println("we received a message: ", msg)
}

func main() {
	server := newServer1()

	go func() {
		time.Sleep(time.Second * 5)
		server.quit()
	}()
	server.start()

}
