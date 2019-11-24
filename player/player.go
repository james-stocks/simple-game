package player

import "math"

type Player struct {
	Level      int
	Experience int
}

func (p *Player) AddExp(e int) {
	p.Experience += e
	p.Level = int(math.Max(1, float64(p.Experience/(100+(p.Level*3)))))
}
