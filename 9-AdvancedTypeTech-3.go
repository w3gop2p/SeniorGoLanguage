package main

import "fmt"

type Position12 struct {
	x, y int
}

func (p *Position12) Move(val int) {
	fmt.Println("The position is moved by: ", val)
}

type Player12 struct {
	Position12
}

func main() {
	p := Player12{}
	p.Move(1000)
}
