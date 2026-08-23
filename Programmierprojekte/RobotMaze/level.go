package main

import "fmt"

var level = []string{
	"#####################",
	"#S#       #         #",
	"# # ##### # ####### #",
	"# #     # #       # #",
	"# ##### # ####### # #",
	"#     # #       # # #",
	"##### # ####### # # #",
	"#     #       # #   #",
	"# ########### # ### #",
	"#           # #   # #",
	"########### # ### # #",
	"#         # # #   # #",
	"# ####### # # # ### #",
	"#       #   # #     #",
	"####### ##### #######",
	"#                   #",
	"# ################# #",
	"#                 #G#",
	"#####################",
}

func LoadWorld() (World, Position, Position) {
	cells := make([][]rune, len(level))

	var start Position
	var goal Position

	for y, row := range level {
		cells[y] = []rune(row)

		for x, cell := range cells[y] {
			switch cell {
			case 'S':
				start = Position{X: x, Y: y}
				cells[y][x] = ' '
			case 'G':
				goal = Position{X: x, Y: y}
				cells[y][x] = ' '
			}
		}
	}

	if len(cells) == 0 {
		panic("level contains no rows")
	}

	width := len(cells[0])
	for _, row := range cells {
		if len(row) != width {
			panic(fmt.Sprintf("level rows have different widths"))
		}
	}

	return World{Cells: cells}, start, goal
}
