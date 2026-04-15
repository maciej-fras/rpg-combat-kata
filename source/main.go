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
	fmt.Printf("c2 health (should be 1000): %.0f\n", c2.GetHealth())

	fmt.Println("\n=== Non-allies can deal damage ===")
	c1.DealDamage(&c3, 100)
	fmt.Printf("c3 health (should be 900): %.0f\n", c3.GetHealth())

	fmt.Println("\n=== Allies can heal each other ===")
	c2.TakeDamage(200)
	fmt.Printf("c2 health before heal: %.0f\n", c2.GetHealth())
	c1.Heal(&c2, 50)
	fmt.Printf("c2 health after ally heal (should be 850): %.0f\n", c2.GetHealth())

	fmt.Println("\n=== Non-allies cannot heal ===")
	fmt.Printf("c3 health before heal: %.0f\n", c3.GetHealth())
	c1.Heal(&c3, 50)
	fmt.Printf("c3 health after enemy heal (should be 900, unchanged): %.0f\n", c3.GetHealth())

	fmt.Println("\n=== Self-heal works ===")
	c3.TakeDamage(200)
	fmt.Printf("c3 health before self-heal: %.0f\n", c3.GetHealth())
	c3.Heal(&c3, 100)
	fmt.Printf("c3 health after self-heal (should be 800): %.0f\n", c3.GetHealth())

	fmt.Println("\n=== Leave faction ===")
	c2.LeaveFaction(alliance)
	c1.DealDamage(&c2, 100)
	fmt.Printf("c2 health after leaving (should be 750): %.0f\n", c2.GetHealth())

	fmt.Println("\n=== Prop: Character can damage a tree ===")
	tree := rpg.NewTree()
	fmt.Printf("Tree health (should be 2000): %.0f\n", tree.GetHealth())
	c1.DealDamage(&tree, 500)
	fmt.Printf("Tree health after attack (should be 1500): %.0f\n", tree.GetHealth())

	fmt.Println("\n=== Prop: Tree is destroyed at 0 ===")
	c1.DealDamage(&tree, 1500)
	fmt.Printf("Tree health (should be 0): %.0f\n", tree.GetHealth())
	fmt.Printf("Tree alive (should be false): %v\n", tree.IsAlive())
}
