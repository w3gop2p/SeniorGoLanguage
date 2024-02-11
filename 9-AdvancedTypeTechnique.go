package main

import "fmt"

type Position struct {
	x, y int
}

type Entity struct {
	name    string
	id      string
	version string
	Position
}

type SpecialEntity struct {
	Entity
	specialField float64
}

func main() {
	e := Entity{
		name:    "my entity",
		id:      "id 1",
		version: "version 1.1",
		Position: Position{
			x: 100,
			y: 200,
		},
	}
	fmt.Printf("%+v\n", e)

	k := SpecialEntity{
		specialField: 88.88,
		Entity:       e,
	}
	fmt.Printf("%+v\n", k.x)
}
