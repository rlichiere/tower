package game

type Player struct {
	lifes int
	money int
	score int
}

func NewPlayer() Player {
	player := Player{lifes: Config.Player.InitialLifes, money: Config.Player.InitialMoney, score: 0}
	return player
}

func (p *Player) LooseLife() {
	p.lifes -= 1
}

func (p *Player) EarnMoney(money int) {
	p.money += money
}

func (p *Player) SpendMoney(money int) bool {
	if p.money < money {
		return false
	}
	p.money -= money
	return true
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
