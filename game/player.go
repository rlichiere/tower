package game

type Player struct {
	lifes int
	money int
	score int
}

func NewPlayer(lifes int, money int) Player {
	player := Player{lifes: lifes, money: money, score: 0}
	return player
}

func (p *Player) LooseLife() {
	p.lifes -= 1
}

func (p *Player) EarnMoney(money int) {
	p.money += money
}

func (p *Player) EarnScore(score int) {
	p.score += score
}

func (p *Player) GetLifes() int {
	return p.lifes
}

func (p *Player) GetMoney() int {
	return p.money
}

func (p *Player) GetScore() int {
	return p.score
}

func (p *Player) IsDead() bool {
	return p.lifes == 0
}
