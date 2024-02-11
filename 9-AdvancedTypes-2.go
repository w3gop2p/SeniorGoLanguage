package main

import "fmt"

type Color int

func (c Color) String() string {
	switch c {
	case ColorBlue:
		return "BLUE"
	case ColorBlack:
		return "BLACK"
	case ColorYellow:
		return "YELLOW"
	case ColorPink:
		return "PINK"
	default:
		panic("invalid color given")

	}
}

const (
	ColorBlue Color = iota
	ColorBlack
	ColorYellow
	ColorPink
)

func main() {
	fmt.Println(ColorYellow)                   // YELLOW
	fmt.Printf("The color is %d", ColorYellow) // 2
	fmt.Printf("The color is %s", ColorYellow) // YELLOW
}
