package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := NewGame()

	ebiten.SetWindowSize(WindowWidth, WindowHeight)
	ebiten.SetWindowTitle("Lagerverwaltung - Go Workshop")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
