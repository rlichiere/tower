package game

import (
	"math"
)

type Tower struct {
	X        int
	Y        int
	Cooling  int
	cooldown int
	Strength int
}

func NewTower(x int, y int) *Tower {
	tower := Tower{X: x, Y: y, Cooling: Config.Tower.InitialCooldown, cooldown: Config.Tower.InitialCooldown, Strength: 1}
	return &tower
}

func (t *Tower) IsReadyToShoot() bool {
	return t.Cooling == 0
}

func (t *Tower) ResetCooldown() {
	t.Cooling = t.cooldown
}

func (t *Tower) ManageShot(enemies EnemiesList, g *Game) {
	if t.Cooling > 0 {
		t.Cooling--
	}
	if !t.IsReadyToShoot() {
		return
	}

	for enemyIndex, enemy := range enemies {
		if t.CheckEnemyRange(enemy) {
			g.ShootEnemy(enemy, t.Strength, enemyIndex)
			t.ResetCooldown()
			break
		}
	}
}

func (t *Tower) CheckEnemyRange(enemy *Enemy) bool {
	return math.Abs(float64(t.X-enemy.X)) == 1 && math.Abs(float64(t.Y-enemy.Y)) == 1
}

type TowersList []*Tower
