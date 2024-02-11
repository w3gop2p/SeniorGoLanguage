package main

import "fmt"

// Pointer - 8 bytes (64 bits) long integer and it points to a memory address

type Player1 struct {
	HP int
}

// if player is not a pointer we are adjusting the copy of the player
// not the actual player.

func TakeDamage(player *Player1, amount int) {
	player.HP -= amount
	fmt.Println("player is taking damage. New HP ->", player.HP)
}

// function receiver. Even in this kind of function if Player is not a pointer
// function will make a copy of the Player

func (p *Player1) TakeDamage(amount int) {
	p.HP -= amount
	fmt.Println("player is taking damage. New HP ->", p.HP)
}

func main() {
	player := &Player1{
		HP: 100,
	}
	TakeDamage(player, 10)
	fmt.Printf("%+v\n", player)

	fmt.Println("-------")
	player.TakeDamage(21)
	fmt.Printf("%+v\n", player)
}
