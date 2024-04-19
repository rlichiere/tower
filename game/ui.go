package game

import (
	"fmt"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func (g *Game) Display(wave int) {
	screen, err := pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("Go ", pterm.FgCyan.ToStyle()),
		putils.LettersFromStringWithStyle("TOWER", pterm.FgLightMagenta.ToStyle()),
	).Srender()

	if err != nil {
		screen = ""
	}
	screen += "\n"
	header := fmt.Sprintf("Iteration: %12d                      Wave: %11d\n", g.GetIteration(), wave)
	header += fmt.Sprintf("Enemies: %14d\n", len(g.Enemies))
	lifeText := fmt.Sprintf("%16d", g.Player.GetLifes())
	switch g.Player.GetLifes() {
	case 10:
		lifeText = pterm.FgLightMagenta.Sprint(lifeText)
	case 9, 8, 7, 6, 5, 4:
		lifeText = pterm.FgLightYellow.Sprint(lifeText)
	case 3, 2, 1, 0:
		lifeText = pterm.FgLightRed.Sprint(lifeText)
	}
	header += fmt.Sprintf("Lifes: %s                      Money: %10d", lifeText, g.Player.GetMoney())
	title := fmt.Sprintf("Score: %15d", g.Player.GetScore())
	screen += pterm.DefaultBox.WithTitle(pterm.FgLightCyan.Sprint(title)).Sprint(pterm.FgLightMagenta.Sprint(header))
	screen += "\n"

	// display grid
	for x := 0; x < len(g.Stage); x++ {
		if x < 1 || x > 30 {
			screen += "   "
		} else {
			screen += pterm.FgLightYellow.Sprintf("%2d", x)
			screen += " "
		}
		for y := 0; y < len(g.Stage[x]); y++ {
			screen += g.Stage[x][y].Content
		}
		screen += "\n"
	}
	screen += "    "
	screen += pterm.FgLightYellow.Sprint(" a b c d e f g h i j k l m n o p q r s t u v w x y z A B C D \n")
	fmt.Println(screen)
}
