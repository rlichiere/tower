package game

type Enemy struct {
	X              int
	Y              int
	Life           int
	PositionInPath int
	Path           EnemyPath
}

func (e *Enemy) MoveNext() {
	if e.PositionInPath+1 == e.Path.PathLength {
		//fmt.Println("Enemy out of path")
		return
	}
	e.PositionInPath++
	var nextCell = e.Path.Path[e.PositionInPath]
	e.X = nextCell.X
	e.Y = nextCell.Y
}

type EnemiesList []*Enemy
