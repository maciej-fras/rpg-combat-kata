package main

import (
	"rpg-combat/rpg"
	// "os"
)

func main() {
	c1 := rpg.NewCharacter()
	c2 := rpg.NewCharacter()

	c1.DealDamage(&c2, 100)
	println("C2 Health: ", c2.Health())
	c2.Heal(&c2, 50)
	println("C2 Health: ", c2.Health())
}
