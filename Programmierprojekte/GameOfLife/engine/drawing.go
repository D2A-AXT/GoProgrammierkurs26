package engine

import (
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// DrawGrid zeichnet ein rechteckiges Zellraster.
// Die eigentliche Spiellogik kennt diese Funktion nicht.
func DrawGrid(
	screen *ebiten.Image,
	x, y, cellSize, rows, cols int,
	isAlive func(row, col int) bool,
	aliveColor, deadColor, gridColor color.Color,
) {
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			fill := deadColor
			if isAlive(row, col) {
				fill = aliveColor
			}

			px := float64(x + col*cellSize)
			py := float64(y + row*cellSize)

			ebitenutil.DrawRect(screen, px, py, float64(cellSize), float64(cellSize), gridColor)
			ebitenutil.DrawRect(screen, px+1, py+1, float64(cellSize-2), float64(cellSize-2), fill)
		}
	}
}

func DrawPanel(screen *ebiten.Image, x, y, width, height int, fill color.Color) {
	ebitenutil.DrawRect(screen, float64(x), float64(y), float64(width), float64(height), fill)
}

func DrawLabel(screen *ebiten.Image, x, y int, text string, textColor color.Color) {
	// DebugPrintAt verwendet die eingebaute Debug-Schrift von Ebitengine.
	// Sie ist für Workshop-Werkzeuge ausreichend und benötigt keine Font-Datei.
	_ = textColor // Farbe wird von DebugPrintAt nicht unterstützt; Parameter hält API erweiterbar.
	ebitenutil.DebugPrintAt(screen, text, x, y)
}

// DrawWrappedText ist eine bewusst einfache Textumbruch-Hilfe.
func DrawWrappedText(screen *ebiten.Image, x, y, maxWidth int, text string, textColor color.Color) {
	_ = textColor
	const approxCharWidth = 6
	maxChars := maxWidth / approxCharWidth
	if maxChars < 10 {
		maxChars = 10
	}

	words := strings.Fields(text)
	line := ""
	lineY := y

	for _, word := range words {
		candidate := word
		if line != "" {
			candidate = line + " " + word
		}

		if len(candidate) > maxChars && line != "" {
			ebitenutil.DebugPrintAt(screen, line, x, lineY)
			lineY += 18
			line = word
		} else {
			line = candidate
		}
	}

	if line != "" {
		ebitenutil.DebugPrintAt(screen, line, x, lineY)
	}
}
