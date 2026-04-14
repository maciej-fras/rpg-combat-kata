package rpg

import "testing"

func TestNewCharacter(t *testing.T) {
	c := NewCharacter()
	if c.Health() != 1000 {
		t.Errorf("expected 1000, got %d", c.Health)
	}
}
