package game

import "fmt"

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
		Index:  -15,
		Armor:  armor,
		game:   game,
	}
	return &wave
}

func (w *Wave) Iterate() {
	if !w.IsFinished() {
		w.Index++
		if w.Index > 0 {
			fmt.Println("Add enemy on stage")
			w.game.AddEnemyOnStage(10 + w.Armor)
		}
	}
}

func (w *Wave) IsFinished() bool {
	return w.Index > w.Length
}
