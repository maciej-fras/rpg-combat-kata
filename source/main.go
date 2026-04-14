package main

import (
	"rpg-combat/rpg"
	// "os"
)

func main() {
	c1 := rpg.NewCharacter()
	c2 := rpg.NewCharacter()

	c2.Type = rpg.Ranged
	c1.Position.X = 0
	c2.Position.X = 20

	// Damage not dealt - target is too far away
	c1.DealDamage(&c2, 100)
	println("C2 Health: ", c2.Health())
	// Damage dealt (ranger attacks warrior)
	c2.DealDamage(&c1, 100)
	println("C1 Health: ", c1.Health())
}
