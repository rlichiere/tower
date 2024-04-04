package game

import (
	_ "embed"
	"fmt"
	"strings"
)

// ˥˩˪    ᒣᒥᒧᒪ—–—‗|⎟          ⎸⎹⎾⎺⎿⏋⎽⏌         ⌈⌉⌊⌋ ⎺⎻⎼⎽     ╔╗╚╝║═      ⯧    ▉

//go:embed maps/map_0.txt
var mapFile string

var Symbols = struct {
	CornerTopRight    string
	CornerBottomLeft  string
	CornerBottomRight string
	EdgeTop           string
	EdgeBottom        string
	EdgeLeft          string
	EdgeRight         string
	CornerTopLeft     string
	EmptyCell         string
	Wall              string
	InputCell         string
	OutputCell        string
	Tower0            string
	Tower1            string
	Tower2            string
	Tower3            string
	Tower4            string
	Tower5            string
	Tower6            string
	Tower7            string
	Tower8            string
	Tower9            string
	Enemy             string
	EnemyAtStart      string
	EnemyAtEnd        string
}{
	CornerTopLeft:     "╔═",
	CornerTopRight:    "═╗",
	CornerBottomLeft:  "╚═",
	CornerBottomRight: "═╝",
	EdgeTop:           "══",
	EdgeBottom:        "══",
	EdgeLeft:          "║ ",
	EdgeRight:         " ║",
	EmptyCell:         " .",
	Wall:              " ▉",
	InputCell:         "> ",
	OutputCell:        " >",
	Tower0:            " Ⅰ",
	Tower1:            " Ⅱ",
	Tower2:            " Ⅲ",
	Tower3:            " Ⅳ",
	Tower4:            " Ⅴ",
	Tower5:            " Ⅵ",
	Tower6:            " Ⅶ",
	Tower7:            " Ⅷ",
	Tower8:            " Ⅸ",
	Tower9:            " Ⅹ",
	Enemy:             " ⯧",
	EnemyAtStart:      "⯧ ",
	EnemyAtEnd:        " ⯧",
}

type StageMap [32][32]Cell

func (s *StageMap) Initialize() {
	fmt.Println("This is a 32x32 map")
	for x := 0; x < 32; x++ {
		for y := 0; y < 32; y++ {
			cell := Cell{X: x, Y: y}
			cell.Content = cell.DetectContent()
			s[x][y] = cell
		}
	}
	fmt.Println("Stage size:", len(s))
	fmt.Println("mapFile:", mapFile)
	s[16][31].Content = Symbols.OutputCell
	s[16][0].Content = Symbols.InputCell
	s.ReadMapFile()
}

func (s *StageMap) InitializeFromStage(source *StageMap) {
	for x := 0; x < 32; x++ {
		for y := 0; y < 32; y++ {
			cellSource := source[x][y]
			cell := Cell{X: cellSource.X, Y: cellSource.Y, Content: cellSource.Content}
			s[x][y] = cell
		}
	}
}

func (s *StageMap) ReadMapFile() {
	lines := strings.Split(mapFile, "\n")
	wallChar := "▉"
	for x, line := range lines {
		var y = 0
		for _, c := range line {
			cString := fmt.Sprintf("%c", c)
			if cString == wallChar {
				s[x][y].Content = Symbols.Wall
			}
			y++
		}
	}
}

func (s *StageMap) CheckEnemyOnOutput(enemies EnemiesList) int {
	for index, enemy := range enemies {
		if enemy.X == 16 && enemy.Y >= 32 {
			fmt.Println("Drop enemy", index, "from Enemies")
			return index
		}
	}
	return -1
}

func Remove(l EnemiesList, index int) EnemiesList {
	return append(l[:index], l[index+1:]...)
}
