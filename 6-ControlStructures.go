package main

import "fmt"

func main() {
	// for loop
	for i := 0; i < 10; i++ {
		fmt.Println("it:", i)
	}
	// loop throw a slice
	fmt.Println("---------------------")
	numbers := []int{1, 2, 3, 5, 6, 7}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// with range
	fmt.Println("---------------------")
	names := []string{"a", "b", "c", "d"}

	for i := range names {
		fmt.Println(names[i])
	}

	for _, name := range names {
		if name == "a" {
			// break
			// return
			fmt.Println(name)
		}
	}
	// maps
	fmt.Println("---------------------")
	users := map[string]int{
		"foo":   1,
		"bar":   2,
		"baz":   3,
		"Alice": 33,
		"Bob":   88,
	}
	for key, value := range users {
		fmt.Printf("key %s value %d\n", key, value)
	}

	//  switch
	fmt.Println("---------------------")
	name := "foo"
	switch name {
	case "Alice":
		fmt.Println("The name Alice")
	case "Bob":
		fmt.Println("The name Bob")
	default:
		fmt.Println("the name is default =>", name)
	}
}
