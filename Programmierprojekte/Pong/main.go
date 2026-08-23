package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := NewPongGame()

	ebiten.SetWindowSize(WindowWidth, WindowHeight)
	ebiten.SetWindowTitle("Pong - Go Programming Workshop")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
