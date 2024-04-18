package main

import (
	"fmt"
	"time"
	"tower/game"
)

func main() {
	fmt.Println("Starting tower")

	player := game.NewPlayer()
	g := game.NewGame(&player)

	gameServer := game.GameServer{Game: g}
	fmt.Println("Start HTTP server...")
	ctx := gameServer.Start()
	fmt.Println("HTTP server started.")

	addEnemyEvery := 5

	// to be removed
	// tower := game.NewTower(2, 2)
	// g.Towers = append(g.Towers, tower)

	wave := game.NewWave(1, 10, 0, g)
	for {
		g.Iterate()
		fmt.Printf("Iteration: %4d    Wave: %4d\n", g.GetIteration(), wave.Id)

		// move enemies
		g.MoveEnemies()

		if g.GetIteration()%addEnemyEvery == 1 {
			wave.Iterate()
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
			break
		}

		if g.GetIteration() > game.Config.Game.MaximumIterations {
			fmt.Println("Maximum iteration reached")
			fmt.Println("Money:", player.GetMoney())
			fmt.Println("Score:", player.GetScore())
			break
		}

		if wave.IsFinished() {
			wave = game.NewWave(wave.Id+1, wave.Length, wave.Armor+1, g)
			g.Player.EarnMoney(game.Config.Game.MoneyPerIteration)
		}

		// sleep iteration time
		time.Sleep(time.Duration(game.Config.Game.IterationDuration) * time.Millisecond)
	}

	gameServer.End(ctx)

}
