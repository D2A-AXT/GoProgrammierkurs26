package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	bgColor       = color.RGBA{R: 24, G: 28, B: 35, A: 255}
	panelColor    = color.RGBA{R: 38, G: 44, B: 54, A: 255}
	rowColor      = color.RGBA{R: 47, G: 54, B: 65, A: 255}
	selectedColor = color.RGBA{R: 58, G: 91, B: 130, A: 255}
	headerColor   = color.RGBA{R: 31, G: 36, B: 45, A: 255}
	lowStockColor = color.RGBA{R: 210, G: 91, B: 76, A: 255}
	okStockColor  = color.RGBA{R: 77, G: 164, B: 104, A: 255}
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(bgColor)

	drawHeader(screen)
	drawWarehouseTable(screen, g)
	drawSummary(screen, g)
	drawControls(screen, g)
}

func drawHeader(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "LAGERVERWALTUNG", 38, 28)
	ebitenutil.DebugPrintAt(screen, "Go Workshop", 38, 48)
}

func drawWarehouseTable(screen *ebiten.Image, g *Game) {
	const (
		x     = 38
		y     = 85
		width = 924
		rowH  = 42
	)

	ebitenutil.DrawRect(screen, x, y, width, rowH, headerColor)

	ebitenutil.DebugPrintAt(screen, "ID", x+12, y+14)
	ebitenutil.DebugPrintAt(screen, "ARTIKEL", x+105, y+14)
	ebitenutil.DebugPrintAt(screen, "BESTAND", x+490, y+14)
	ebitenutil.DebugPrintAt(screen, "MIN.", x+590, y+14)
	ebitenutil.DebugPrintAt(screen, "PREIS", x+665, y+14)
	ebitenutil.DebugPrintAt(screen, "STATUS", x+795, y+14)

	for i, product := range g.Warehouse.Products {
		rowY := y + rowH*(i+1)

		background := rowColor
		if i == g.Selected {
			background = selectedColor
		}

		ebitenutil.DrawRect(
			screen,
			float64(x),
			float64(rowY),
			float64(width),
			float64(rowH-2),
			background,
		)

		ebitenutil.DebugPrintAt(screen, product.ID, x+12, rowY+13)
		ebitenutil.DebugPrintAt(screen, product.Name, x+105, rowY+13)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d", product.Quantity), x+490, rowY+13)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d", product.Minimum), x+590, rowY+13)
		ebitenutil.DebugPrintAt(screen, formatCents(product.PriceCents), x+665, rowY+13)

		if IsLowStock(product) {
			ebitenutil.DrawRect(
				screen,
				float64(x+795),
				float64(rowY+9),
				94,
				23,
				lowStockColor,
			)
			ebitenutil.DebugPrintAt(screen, "NACHBESTELLEN", x+801, rowY+14)
		} else {
			ebitenutil.DrawRect(
				screen,
				float64(x+795),
				float64(rowY+9),
				55,
				23,
				okStockColor,
			)
			ebitenutil.DebugPrintAt(screen, "OK", x+813, rowY+14)
		}
	}
}

func drawSummary(screen *ebiten.Image, g *Game) {
	const (
		x = 38
		y = 515
		w = 924
		h = 52
	)

	ebitenutil.DrawRect(screen, x, y, w, h, panelColor)

	totalQuantity := TotalQuantity(g.Warehouse)
	totalValue := InventoryValueCents(g.Warehouse)

	text := fmt.Sprintf(
		"Produkte: %d    Artikel gesamt: %d    Lagerwert: %s",
		len(g.Warehouse.Products),
		totalQuantity,
		formatCents(totalValue),
	)

	ebitenutil.DebugPrintAt(screen, text, x+14, y+18)
}

func drawControls(screen *ebiten.Image, g *Game) {
	ebitenutil.DebugPrintAt(
		screen,
		"↑/↓ oder W/S: Auswahl    E: +10 einlagern    A: -1 auslagern    N: Demo-Artikel anlegen    R: Reset",
		38,
		590,
	)

	ebitenutil.DebugPrintAt(screen, g.Message, 38, 618)
}

func formatCents(cents int) string {
	euros := cents / 100
	rest := cents % 100

	return fmt.Sprintf("%d,%02d EUR", euros, rest)
}
