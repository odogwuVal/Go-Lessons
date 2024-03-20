package main

import "fmt"

type ninjaStar struct {
	owner string
}

type ninjaSword struct {
	owner string
}

type ninjaWeapon interface {
	attack()
}

func (ninjaStar) attack() {
	fmt.Println("Throwing ninja stars")
}

func (ninjaSword) attack() {
	fmt.Println("Swinging ninja sword")
}

func attack(weapon ninjaWeapon) {
	weapon.attack()
}

func main() {
	weapons := []ninjaWeapon{ninjaStar{owner: "wallace"}, ninjaSword{owner: "wallace"}}
	for _, weapon := range weapons {
		attack(weapon)
	}
}
