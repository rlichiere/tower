package game

type Game struct {
	Player      *Player
	StageGround *StageMap // static data of the stage
	Stage       *StageMap
	Enemies     EnemiesList
	EnemyPath   EnemyPath
	Towers      TowersList
	iteration   int
}

func NewGame(player *Player) *Game {
	game := Game{
		Player:      player,
		StageGround: &StageMap{},
		Stage:       &StageMap{},
		Enemies:     EnemiesList{},
		EnemyPath:   NewEnemyPath(),
		Towers:      TowersList{},
		iteration:   0,
	}
	game.StageGround.Initialize()
	return &game
}

func (g *Game) Iterate() {
	g.iteration++
}

func (g *Game) GetIteration() int {
	return g.iteration
}

func (g *Game) AddEnemyOnStage(life int) {
	enemy := Enemy{Life: life, X: 16, Y: 0, PositionInPath: 0, Path: g.EnemyPath}
	g.Enemies = append(g.Enemies, &enemy)
}

func (g *Game) MoveEnemies() {
	for _, enemy := range g.Enemies {
		enemy.MoveNext()
	}
}

func (g *Game) ManagerTowersShots() {
	for _, tower := range g.Towers {
		tower.ManageShot(g.Enemies, g)
	}
}

func (g *Game) BuildTower(x int, y int) (ok bool, err string) {
	done := g.Player.SpendMoney(Config.Game.MoneyPerTowerBuild)
	if !done {
		return false, "Not enough money to build"
	}
	tower := NewTower(x, y)
	g.Towers = append(g.Towers, tower)
	return true, ""
}

func (g *Game) UpgradeTower(x int, y int) (ok bool, err string) {
	for _, tower := range g.Towers {
		if tower.X == x && tower.Y == y {
			if tower.Strength == 10 {
				return false, "Tower is already maxed"
			}
			done := g.Player.SpendMoney((tower.Strength + 1) * Config.Game.MoneyPerTowerUpgrade)
			if !done {
				return false, "Not enough money to upgrade"
			}
			tower.Strength += 1
			break
		}
	}
	return true, ""
}

func (g *Game) ShootEnemy(enemy *Enemy, shotStrenght int, enemyIndex int) {
	enemy.Life -= shotStrenght
	if enemy.Life <= 0 {
		g.KillEnemy(enemy, enemyIndex)
	}
}

func (g *Game) KillEnemy(enemy *Enemy, enemyIndex int) {
	//fmt.Println("Kill enemy X:", enemy.X, "Y:", enemy.Y, "Position in path:", enemy.PositionInPath)
	g.Enemies = Remove(g.Enemies, enemyIndex)
	g.Player.EarnMoney(Config.Game.MoneyPerKill)
	g.Player.EarnScore(enemy.PositionInPath)
}

func (g *Game) CompileStage() {
	g.Stage.InitializeFromGround(g.StageGround)

	// enemies
	for _, enemy := range g.Enemies {
		if enemy.X == 16 && enemy.Y == 0 {
			g.Stage[enemy.X][enemy.Y].Content = Symbols.EnemyAtStart
		} else if enemy.X == 16 && enemy.Y == 31 {
			g.Stage[enemy.X][enemy.Y].Content = Symbols.EnemyAtEnd
		} else {
			g.Stage[enemy.X][enemy.Y].Content = Symbols.Enemy
		}
	}

	// towers
	for _, tower := range g.Towers {
		switch tower.Strength {
		case 1:
			g.Stage[tower.X][tower.Y].Content = Symbols.Tower0
		case 2:
			g.Stage[tower.X][tower.Y].Content = Symbols.Tower1
		case 3:
			g.Stage[tower.X][tower.Y].Content = Symbols.Tower2
		case 4:
			g.Stage[tower.X][tower.Y].Content = Symbols.Tower3
		case 5:
			g.Stage[tower.X][tower.Y].Content = Symbols.Tower4
		case 6:
			g.Stage[tower.X][tower.Y].Content = Symbols.Tower5
		case 7:
			g.Stage[tower.X][tower.Y].Content = Symbols.Tower6
		case 8:
			g.Stage[tower.X][tower.Y].Content = Symbols.Tower7
		case 9:
			g.Stage[tower.X][tower.Y].Content = Symbols.Tower8
		default:
			g.Stage[tower.X][tower.Y].Content = Symbols.Tower9
		}
		g.Stage[tower.X][tower.Y].Kind = "tower"
	}
}
