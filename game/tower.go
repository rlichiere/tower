package game

import (
	"fmt"
	"math"
)

type Tower struct {
	X             int
	Y             int
	Reload        int
	initialReload int
}

func NewTower(x int, y int, reload int) *Tower {
	tower := Tower{X: x, Y: y, Reload: reload, initialReload: reload}
	return &tower
}

func (t *Tower) IsReady() bool {
	return t.Reload == 0
}

func (t *Tower) ResetReload() {
	t.Reload = t.initialReload
}

func (t *Tower) ManageShot(enemies EnemiesList, g *Game) {
	if t.Reload > 0 {
		t.Reload--
	}
	fmt.Println("tower reload:", t.Reload)
	killAtIndex := -1
	for enemyIndex, enemy := range enemies {
		if t.IsReady() && t.CheckEnemyRange(enemy) {
			fmt.Println("Kill enemy !")
			killAtIndex = enemyIndex
			t.ResetReload()
		}
	}
	if killAtIndex >= 0 {
		g.KillEnemy(killAtIndex)
	}
}

func (t *Tower) CheckEnemyRange(enemy *Enemy) bool {

	return math.Abs(float64(t.X-enemy.X)) == 1 && math.Abs(float64(t.Y-enemy.Y)) == 1
}

type TowersList []*Tower
