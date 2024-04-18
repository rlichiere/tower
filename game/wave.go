package game

type Wave struct {
	Id     int
	Length int
	Index  int
	Armor  int
	game   *Game
}

func NewWave(id int, length int, armor int, game *Game) *Wave {
	wave := Wave{
		Id:     id,
		Length: length,
		Index:  Config.Wave.StartDecay,
		Armor:  armor,
		game:   game,
	}
	return &wave
}

func (w *Wave) Iterate() {
	if !w.IsFinished() {
		w.Index++
		if w.Index > 0 {
			w.game.AddEnemyOnStage(w.Id)
		}
	}
}

func (w *Wave) IsFinished() bool {
	return w.Index > w.Length
}
