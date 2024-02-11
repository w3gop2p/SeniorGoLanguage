package main

import "fmt"

// constants
const acceleration = 9.8

// Global Variables
var x = 10
var (
	name             = "Foo"
	firstName string = "Foo"
	lastName  string
)

func main() {
	// local variables
	var versions int
	version := 1           // infer int
	otherVersion := "Bar"  // infer string
	anotherVersion := 10.1 // infer float32 a float64
	fmt.Println(version, otherVersion, anotherVersion)
	fmt.Println(versions)

	// constants
	fmt.Println(acceleration)

}
