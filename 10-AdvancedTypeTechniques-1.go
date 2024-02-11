package main

import (
	"fmt"
	"strings"
)

type TransformFunc func(s string) string

func Uppercase(s string) string {
	return strings.ToUpper(s)
}

func Prefixer(prefix string) TransformFunc {
	return func(s string) string {
		return prefix + s
	}
}

func transformString(s string, fn TransformFunc) string {
	return fn(s)
}

func main() {
	fmt.Println(transformString("Hello Sailor!", Uppercase))
	fmt.Println(transformString("Hello Sailor!", Prefixer("FOO_")))
}

// when to use interface and type function? use interface if you want to store state, for simple things like transform functions or this stateless pure functions use typed function
