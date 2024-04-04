package game

import "fmt"

func (g *Game) Display() {
	var screen string
	// display player data
	screen = fmt.Sprintf("Score: %15d        Lifes: %2d        Money: %10d\n", g.Player.GetScore(), g.Player.GetLifes(), g.Player.GetMoney())

	// display grid
	for x := 0; x < len(g.Stage); x++ {
		if x < 1 || x > 30 {
			screen += "   "
		} else {
			screen += fmt.Sprintf("%2d ", x)
		}
		for y := 0; y < len(g.Stage[x]); y++ {
			screen += g.Stage[x][y].Content
		}
		screen += "\n"
	}
	screen += "      a b c d e f g h i j k l m n o p q r s t u v x y z A B C D E\n"
	fmt.Println(screen)
}