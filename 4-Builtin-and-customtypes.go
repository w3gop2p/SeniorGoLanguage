package main

import "fmt"

var (
	floatVar32 float32 = 0.1
	floatVar64 float64 = 0.1
	surname    string  = "Foo"
	intVar32   int32   = 1
	intVar64           = 48484
	intVar     int     = -1
	uintVar    uint    = 1
	uint32Var  uint32  = 1
	uint64Var  uint64  = 10
	uint8Var   uint8   = 0x1 // (X1 is 1 in hexadecimal
	byteVar    byte    = 0x2 // basically uint8 and byte are the same
	runVar     rune    = 'a'
)

type Player struct {
	name        string
	health      int
	attackPower float64
}

func getHealth(player Player) int {
	return player.health
}

func (player Player) getHealth1() int {
	return player.health
}

// Custom types

type version int

type Weapon string // here I have a type weapon and it will be a string

func getWeapon(weapon Weapon) string {
	return string(weapon)
}

func main() {
	player := Player{
		name:        "Captain Jack",
		health:      100,
		attackPower: 45.1,
	}
	fmt.Printf("this is the player: %+v\n", player) // here we can use without + before v
	fmt.Printf("this is the player health: %d\n", getHealth(player))
	fmt.Printf("this is the player health: %d\n", player.getHealth1())

	// maps
	// initialize a map
	// users := make(map[string]int)
	// users := map[string]int{}
	users := map[string]int{
		"John":  37,
		"Klare": 35,
	}
	users["foo"] = 10
	users["bar"] = 11
	fmt.Printf("users map is %+v\n", users)
	fmt.Printf("age of Johh is %d\n", users["John"])
	fmt.Printf("age of Johh is %d\n", users["buz"])

	// if a key doesn't exists for example user["buzz"], Golang for this value will show 0
	// so will solve this in this manner
	age, ok := users["baz"]
	if !ok {
		fmt.Println("baz not exist in the map")
	} else {
		fmt.Println("baz exists in the map", age)
	}

	// delete a key from the map
	delete(users, "foo")
	fmt.Println(users)

	// iterate throw the map
	for k, v := range users {
		fmt.Printf("the key %s and the value %d\n", k, v)
	}
	// also we can iterate just throw the keys of the map
	for k := range users {
		fmt.Printf("the key %s\n", k)
	}

	// another way to inialize a map
	utilizers := make(map[string]int)
	utilizers["foo"] = 10

	// Slices can be expandable
	// initialization of the slice
	numbersA := []int{}

	numbers := []int{1, 2, 3}
	numbers = append(numbers, 4)
	numbers = append(numbers, 5)
	numbers = append(numbers, 6)

	otherNumbers := make([]int, 5) // it is specified the length of the slice, also here instead of 5 can be zero

	fmt.Println(numbers)
	fmt.Println(numbersA)
	fmt.Println(otherNumbers)

	// arays has a fixed size
	numbersB := [2]int{}
	numbersB[0] = 1
	numbersB[1] = 2

	fmt.Println(numbersB)
}
