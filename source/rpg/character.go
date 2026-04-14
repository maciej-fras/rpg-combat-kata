package rpg

import "fmt"

const (
	MaxHealth                       = 1000
	MinLevelDiffToModifyDamageDealt = 5
)

type Character struct {
	health float64
	Level  int
	Alive  bool
}

func NewCharacter() Character {
	return Character{health: MaxHealth, Level: 1, Alive: true}
}

func (c *Character) Health() float64 {
	return c.health
}

// Setter with constraints
func (c *Character) SetHealth(value float64) {
	if value <= 0 {
		c.health = 0
		c.Alive = false
	} else if value > MaxHealth {
		c.health = MaxHealth
	} else {
		c.health = value
	}
}

func (c *Character) Heal(target *Character, healValue float64) string {
	if c != target {
		return "Character cannot heal different character"
	}
	if !target.Alive {
		return "Dead character cannot be healed"
	}

	increasedHealth := target.health + healValue
	target.SetHealth(increasedHealth)
	return fmt.Sprintf("Character healed by %f up to %f", healValue, target.Health())
}

func (c *Character) DealDamage(target *Character, damage float64) {
	if c == target {
		return
	}
	levelDiff := c.Level - target.Level
	if levelDiff >= MinLevelDiffToModifyDamageDealt {
		damage = damage * 1.5
	} else if levelDiff <= -MinLevelDiffToModifyDamageDealt {
		damage = damage * 0.5
	}

	target.TakeDamage(damage)
}

func (c *Character) TakeDamage(damage float64) {
	decreasedHealth := c.health - damage
	c.SetHealth(decreasedHealth)
}
