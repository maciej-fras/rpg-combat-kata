package rpg

var _ Damagable = (*Character)(nil)

type Damagable interface {
	TakeDamage(damage float64)
	GetHealth() float64
	IsAlive() bool
}

type Health struct {
	health    float64
	maxHealth float64
	alive     bool
}

func NewHealth(maxHealth float64) Health {
	return Health{health: maxHealth, maxHealth: maxHealth, alive: true}
}

func (h *Health) GetHealth() float64 {
	return h.health
}

func (h *Health) IsAlive() bool {
	return h.alive
}

func (h *Health) SetHealth(value float64) {
	if value <= 0 {
		h.health = 0
		h.alive = false
	} else if value > h.maxHealth {
		h.health = h.maxHealth
	} else {
		h.health = value
	}
}

func (h *Health) TakeDamage(damage float64) {
	h.SetHealth(h.health - damage)
}
