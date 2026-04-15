package rpg

type Prop struct {
	Health
	Name string
}

func NewTree() Prop {
	return Prop{Health: NewHealth(2000), Name: "Tree"}
}
