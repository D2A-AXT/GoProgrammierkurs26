package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	bgColor       = color.RGBA{R: 22, G: 25, B: 31, A: 255}
	panelColor    = color.RGBA{R: 37, G: 42, B: 52, A: 255}
	selectedColor = color.RGBA{R: 60, G: 92, B: 132, A: 255}
	itemColor     = color.RGBA{R: 206, G: 164, B: 72, A: 255}
	exitColor     = color.RGBA{R: 90, G: 165, B: 225, A: 255}
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(bgColor)

	room := g.currentRoom()
	if room == nil {
		ebitenutil.DebugPrintAt(screen, "Ungültiger Raum.", 30, 30)
		return
	}

	drawRoomPanel(screen, room, g)
	drawInventory(screen, g.Player)
	drawMessage(screen, g.Message)
	drawHelp(screen)
}

func drawRoomPanel(screen *ebiten.Image, room *Room, g *Game) {
	const (
		x = 35
		y = 35
		w = 610
		h = 500
	)

	ebitenutil.DrawRect(screen, x, y, w, h, panelColor)

	ebitenutil.DebugPrintAt(screen, room.Name, x+20, y+20)
	ebitenutil.DebugPrintAt(screen, room.Description, x+20, y+48)

	ebitenutil.DebugPrintAt(screen, "GEGENSTÄNDE", x+20, y+105)

	itemY := y + 135
	for i, item := range room.Items {
		selected := g.SelectedItem == i
		drawChoice(screen, x+20, itemY+i*46, 560, item.Name, "E: aufnehmen", itemColor, selected)
	}

	exitStartY := itemY + max(1, len(room.Items))*46 + 35
	ebitenutil.DebugPrintAt(screen, "AUSGÄNGE", x+20, exitStartY)

	for i, exit := range room.Exits {
		selected := g.SelectedExit == i
		label := exit.Name

		if exit.RequiredItem != "" {
			label += "  [verschlossen]"
			if HasItem(g.Player, exit.RequiredItem) {
				label = exit.Name + "  [Schlüssel vorhanden]"
			}
		}

		drawChoice(
			screen,
			x+20,
			exitStartY+30+i*46,
			560,
			label,
			"SPACE: benutzen",
			exitColor,
			selected,
		)
	}
}

func drawChoice(
	screen *ebiten.Image,
	x, y, w int,
	title string,
	action string,
	accent color.Color,
	selected bool,
) {
	background := color.RGBA{R: 48, G: 54, B: 65, A: 255}
	if selected {
		background = selectedColor
	}

	ebitenutil.DrawRect(
		screen,
		float64(x),
		float64(y),
		float64(w),
		38,
		background,
	)

	ebitenutil.DrawRect(
		screen,
		float64(x),
		float64(y),
		6,
		38,
		accent,
	)

	ebitenutil.DebugPrintAt(screen, title, x+16, y+8)
	ebitenutil.DebugPrintAt(screen, action, x+w-130, y+8)
}

func drawInventory(screen *ebiten.Image, player Player) {
	const (
		x = 680
		y = 35
		w = 265
		h = 500
	)

	ebitenutil.DrawRect(screen, x, y, w, h, panelColor)
	ebitenutil.DebugPrintAt(screen, "INVENTAR", x+18, y+20)

	if len(player.Inventory) == 0 {
		ebitenutil.DebugPrintAt(screen, "(leer)", x+18, y+55)
		return
	}

	for i, item := range player.Inventory {
		ebitenutil.DrawRect(
			screen,
			float64(x+18),
			float64(y+52+i*52),
			float64(w-36),
			42,
			color.RGBA{R: 48, G: 54, B: 65, A: 255},
		)

		ebitenutil.DebugPrintAt(screen, item.Name, x+30, y+65+i*52)
	}
}

func drawMessage(screen *ebiten.Image, message string) {
	const (
		x = 35
		y = 560
		w = 910
		h = 45
	)

	ebitenutil.DrawRect(screen, x, y, w, h, panelColor)
	ebitenutil.DebugPrintAt(screen, message, x+15, y+15)
}

func drawHelp(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(
		screen,
		"↑/↓ oder W/S: auswählen    E: Gegenstand aufnehmen    SPACE: Ausgang benutzen    R: Reset",
		35,
		625,
	)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func formatInventory(player Player) string {
	result := ""
	for i, item := range player.Inventory {
		if i > 0 {
			result += ", "
		}
		result += item.Name
	}
	return fmt.Sprintf("[%s]", result)
}
