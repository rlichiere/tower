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

func (g *Game) KillEnemy(enemyIndex int) {
	g.Enemies = Remove(g.Enemies, enemyIndex)
}

func (g *Game) CompileStage() {
	g.Stage.InitializeFromStage(g.StageGround)

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
		g.Stage[tower.X][tower.Y].Content = Symbols.Tower0
	}
}
