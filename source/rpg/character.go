package rpg

import (
	"fmt"
)

const (
	MaxHealth                       = 1000
	MinLevelDiffToModifyDamageDealt = 5
)

type Character struct {
	health   float64
	Level    int
	Alive    bool
	Type     FighterType
	Position Position
	Factions map[*Faction]bool
}

func NewCharacter() Character {
	return Character{
		health:   MaxHealth,
		Level:    1,
		Alive:    true,
		Type:     Melee,
		Position: Position{X: 0},
		Factions: make(map[*Faction]bool),
	}
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
	if c != target && !c.IsAlly(target) {
		return "Character cannot heal character from different faction"
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

	if c.IsAlly(target) {
		return
	}

	distance := c.Position.X - target.Position.X
	if distance < 0 {
		distance = -distance
	}

	if distance > int(c.Type.MaxRange()) {
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

func (c *Character) JoinFaction(faction *Faction) {
	c.Factions[faction] = true
}

func (c *Character) LeaveFaction(faction *Faction) {
	delete(c.Factions, faction)
}

func (c *Character) IsAlly(otherCharacter *Character) bool {
	for faction := range c.Factions {
		if otherCharacter.Factions[faction] {
			return true
		}
	}
	return false
}
