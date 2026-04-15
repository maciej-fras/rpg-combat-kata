package rpg

type FighterType int

const (
	Melee FighterType = iota
	Ranged
)

func (ft FighterType) MaxRange() float64 {
	switch ft {
	case Ranged:
		return 20
	default:
		return 2
	}
}
