package player

type Player struct {
    Level int
	Experience int
}

func (p Player) AddExp(e int) {
   p.Experience += e
}