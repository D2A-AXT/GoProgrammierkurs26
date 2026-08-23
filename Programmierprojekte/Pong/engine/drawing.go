package engine

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// DrawRect zeichnet ein gefülltes Rechteck.
func DrawRect(screen *ebiten.Image, x, y, width, height float64, clr color.Color) {
	ebitenutil.DrawRect(screen, x, y, width, height, clr)
}

// DrawCenteredRect zeichnet ein Rechteck um einen Mittelpunkt.
func DrawCenteredRect(screen *ebiten.Image, centerX, centerY, width, height float64, clr color.Color) {
	ebitenutil.DrawRect(
		screen,
		centerX-width/2,
		centerY-height/2,
		width,
		height,
		clr,
	)
}

// DrawRectOutline zeichnet einen einfachen Rechteckrahmen.
func DrawRectOutline(screen *ebiten.Image, x, y, width, height, thickness float64, clr color.Color) {
	ebitenutil.DrawRect(screen, x, y, width, thickness, clr)
	ebitenutil.DrawRect(screen, x, y+height-thickness, width, thickness, clr)
	ebitenutil.DrawRect(screen, x, y, thickness, height, clr)
	ebitenutil.DrawRect(screen, x+width-thickness, y, thickness, height, clr)
}

// DrawDashedVerticalLine zeichnet eine gestrichelte vertikale Linie.
func DrawDashedVerticalLine(screen *ebiten.Image, x, y, height, dashLength, gapLength float64, clr color.Color) {
	for currentY := y; currentY < y+height; currentY += dashLength + gapLength {
		remaining := y + height - currentY
		length := dashLength
		if remaining < length {
			length = remaining
		}
		ebitenutil.DrawRect(screen, x, currentY, 2, length, clr)
	}
}
