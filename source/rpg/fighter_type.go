package rpg

type FighterType int

const (
	Melee  FighterType = iota // 0
	Ranged                    // 1
)

func (ft FighterType) MaxRange() float64 {
	switch ft {
	case Ranged:
		return 20
	default:
		return 2
	}
}
