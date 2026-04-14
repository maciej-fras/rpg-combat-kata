package rpg

import "fmt"

const MaxHealth = 1000

type Character struct {
	health int
	Level  int
	Alive  bool
}

func NewCharacter() Character {
	return Character{health: MaxHealth, Level: 1, Alive: true}
}

func (c *Character) Health() int {
	return c.health
}

// Setter with constraints
func (c *Character) SetHealth(value int) {
	if value <= 0 {
		c.health = 0
		c.Alive = false
	} else if value > MaxHealth {
		c.health = MaxHealth
	} else {
		c.health = value
	}
}

func (c *Character) TakeDamage(damage int) {
	decreasedHealth := c.health - damage
	c.SetHealth(decreasedHealth)
}

func (c *Character) Heal(target *Character, healValue int) string {
	if !target.Alive {
		return "Dead character cannot be healed"
	} else {
		increasedHealth := target.health + healValue
		target.SetHealth(increasedHealth)
		return fmt.Sprintf("Character healed by %d up to %d", healValue, target.Health())
	}
}

func (c *Character) DealDamage(target *Character, damage int) {
	target.TakeDamage(damage)
}
