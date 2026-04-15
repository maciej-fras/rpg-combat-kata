package rpg

import (
	"fmt"
)

const (
	MaxHealth                       = 1000
	MinLevelDiffToModifyDamageDealt = 5
)

type Character struct {
	Health
	Level    int
	Type     FighterType
	Position Position
	Factions map[*Faction]bool
}

func NewCharacter() Character {
	return Character{
		Health:   NewHealth(MaxHealth),
		Level:    1,
		Type:     Melee,
		Position: Position{X: 0},
		Factions: make(map[*Faction]bool),
	}
}

func (c *Character) Heal(target *Character, healValue float64) string {
	if c != target && !c.IsAlly(target) {
		return "Character cannot heal character from different faction"
	}
	if !target.IsAlive() {
		return "Dead character cannot be healed"
	}

	increasedHealth := target.GetHealth() + healValue
	target.SetHealth(increasedHealth)
	return fmt.Sprintf("Character healed by %f up to %f", healValue, target.GetHealth())
}

func (c *Character) DealDamage(target Damagable, damage float64) {
	damage = c.adjustDamageForTarget(target, damage)
	if damage == 0 {
		return
	}
	target.TakeDamage(damage)
}

func (c *Character) adjustDamageForTarget(target Damagable, damage float64) float64 {
	targetChar, ok := target.(*Character)
	if !ok { // non-character
		return damage
	}
	if c == targetChar {
		return 0
	}
	if c.IsAlly(targetChar) {
		return 0
	}
	if !c.isInRange(targetChar) {
		return 0
	}

	levelDiff := c.Level - targetChar.Level
	if levelDiff >= MinLevelDiffToModifyDamageDealt {
		damage *= 1.5
	} else if levelDiff <= -MinLevelDiffToModifyDamageDealt {
		damage *= 0.5
	}
	return damage
}

func (c *Character) isInRange(target *Character) bool {
	distance := c.Position.X - target.Position.X
	if distance < 0 {
		distance = -distance
	}
	return distance <= int(c.Type.MaxRange())
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
