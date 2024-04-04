package main

import (
	"fmt"
	"os"
	"time"
	"tower/game"
)

func main() {
	fmt.Println("Starting tower")

	player := game.NewPlayer(10, 100)
	g := game.NewGame(&player)

	// create enemies list

	addEnemyEvery := 5
	// waveLength := 20

	// to be removed
	tower := game.NewTower(2, 2, 15)
	g.Towers = append(g.Towers, tower)

	wave := game.NewWave(1, 10, 1, g)
	for {
		g.Iterate()
		fmt.Printf("Iteration: %4d    Wave: %4d\n", g.GetIteration(), wave.Id)

		// move enemies
		g.MoveEnemies()

		if g.GetIteration()%addEnemyEvery == 1 {
			// add enemy on stage
			// g.AddEnemyOnStage(10)
			wave.Iterate()
		} else {
			fmt.Println("Enemies on stage:", len(g.Enemies))
		}

		// manage tower shots
		g.ManagerTowersShots()

		// manage enemy output
		outputEnemy := g.Stage.CheckEnemyOnOutput(g.Enemies)
		if outputEnemy >= 0 {
			player.LooseLife()
			g.Enemies = game.Remove(g.Enemies, outputEnemy)
		}

		// manage towers shot
		g.CompileStage()
		g.Display()

		if player.IsDead() {
			fmt.Println("Player is dead")
			fmt.Println("Money:", player.GetMoney())
			fmt.Println("Score:", player.GetScore())
			os.Exit(0)
		}

		if g.GetIteration() > 100000 {
			fmt.Println("Maximum iteration reached")
			fmt.Println("Money:", player.GetMoney())
			fmt.Println("Score:", player.GetScore())
			os.Exit(0)
		}

		// increase money
		player.EarnMoney(100)

		if wave.IsFinished() {
			wave = game.NewWave(wave.Id+1, wave.Length, wave.Armor+1, g)
		}

		// sleep iteration time
		time.Sleep(1 * time.Second / 5)
	}
}
