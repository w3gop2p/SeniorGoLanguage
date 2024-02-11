package main

import "fmt"

type WeaponType int

func (w WeaponType) String() string {
	switch w {
	case Sword:
		return "SWORD"
	case Axe:
		return "AXE"
	case WoodenStick:
		return "WOODENSTICK"
	}
	return ""
}

const (
	Axe         WeaponType = 1 // also we can do: Axe WeaponType = iota, and for the rest initialization will be automatic incremental from 1 by 1
	Sword       WeaponType = 2
	WoodenStick WeaponType = 3
	Knife       WeaponType = 4
)

func getDamage(weaponType WeaponType) int {
	switch weaponType {
	case Axe:
		return 100
	case Sword:
		return 90
	case WoodenStick:
		return 40
	case Knife:
		return 40
	default:
		panic("weapon does not exists")
	}
}

func main() {
	fmt.Println("Damage of weapon:", getDamage(Axe))
	fmt.Printf("Damage of weapon (%d) (%d):\n", Axe, getDamage(Axe))
	fmt.Printf("Damage of weapon (%d) (%d):\n", Sword, getDamage(Sword))
	fmt.Printf("Damage of weapon (%d) (%d):\n", WoodenStick, getDamage(WoodenStick))

	// because of method String attached to the type WeaponType we can do like this
	fmt.Printf("Damage of weapon (%s) (%d):\n", Axe, getDamage(Axe))
	fmt.Printf("Damage of weapon (%s) (%d):\n", Sword, getDamage(Sword))
	fmt.Printf("Damage of weapon (%s) (%d):\n", WoodenStick, getDamage(WoodenStick))
}
