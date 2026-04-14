package main

import (
	"fmt"
	"rpg-combat/rpg"
)

func main() {
	c1 := rpg.NewCharacter()
	c2 := rpg.NewCharacter()
	c3 := rpg.NewCharacter()

	alliance := &rpg.Faction{Name: "Alliance"}
	horde := &rpg.Faction{Name: "Horde"}

	c1.JoinFaction(alliance)
	c2.JoinFaction(alliance)
	c3.JoinFaction(horde)

	fmt.Println("=== Allies cannot deal damage ===")
	c1.DealDamage(&c2, 100)
	fmt.Printf("c2 health (should be 1000): %.0f\n", c2.Health())

	fmt.Println("\n=== Non-allies can deal damage ===")
	c1.DealDamage(&c3, 100)
	fmt.Printf("c3 health (should be 900): %.0f\n", c3.Health())

	fmt.Println("\n=== Allies can heal each other ===")
	c2.TakeDamage(200)
	fmt.Printf("c2 health before heal: %.0f\n", c2.Health())
	c1.Heal(&c2, 50)
	fmt.Printf("c2 health after ally heal (should be 850): %.0f\n", c2.Health())

	fmt.Println("\n=== Non-allies cannot heal ===")
	fmt.Printf("c3 health before heal: %.0f\n", c3.Health())
	c1.Heal(&c3, 50)
	fmt.Printf("c3 health after enemy heal (should be 900, unchanged): %.0f\n", c3.Health())

	fmt.Println("\n=== Self-heal works ===")
	c3.TakeDamage(200)
	fmt.Printf("c3 health before self-heal: %.0f\n", c3.Health())
	c3.Heal(&c3, 100)
	fmt.Printf("c3 health after self-heal (should be 800): %.0f\n", c3.Health())

	fmt.Println("\n=== Leave faction ===")
	c2.LeaveFaction(alliance)
	c1.DealDamage(&c2, 100)
	fmt.Printf("c2 health after leaving (should be 750): %.0f\n", c2.Health())
}
