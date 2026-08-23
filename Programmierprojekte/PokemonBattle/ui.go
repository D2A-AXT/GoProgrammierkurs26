package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	backgroundColor = color.RGBA{R: 25, G: 29, B: 38, A: 255}
	panelColor      = color.RGBA{R: 41, G: 47, B: 59, A: 255}
	textColor       = color.RGBA{R: 235, G: 238, B: 244, A: 255}
	mutedColor      = color.RGBA{R: 170, G: 177, B: 190, A: 255}
	playerColor     = color.RGBA{R: 88, G: 168, B: 255, A: 255}
	enemyColor      = color.RGBA{R: 238, G: 102, B: 88, A: 255}
	healthColor     = color.RGBA{R: 86, G: 196, B: 112, A: 255}
	healthBack      = color.RGBA{R: 75, G: 79, B: 88, A: 255}
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)

	drawTitle(screen)
	drawMonsterCard(screen, g.Player, 55, 110, playerColor, false)
	drawMonsterCard(screen, g.Enemy, 505, 110, enemyColor, true)
	drawVersus(screen)
	drawBattlePanel(screen, g)
}

func drawTitle(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(
		screen,
		"MONSTER BATTLE",
		370,
		35,
	)
	ebitenutil.DebugPrintAt(
		screen,
		"Go Workshop",
		405,
		60,
	)
}

func drawMonsterCard(
	screen *ebiten.Image,
	monster Monster,
	x int,
	y int,
	accent color.Color,
	enemy bool,
) {
	const cardW = 340
	const cardH = 270

	ebitenutil.DrawRect(screen, float64(x), float64(y), cardW, cardH, panelColor)

	// Monster symbol
	cx := float64(x + cardW/2)
	cy := float64(y + 88)

	ebitenutil.DrawCircle(screen, cx, cy, 50, accent)
	ebitenutil.DrawCircle(screen, cx-17, cy-8, 7, color.White)
	ebitenutil.DrawCircle(screen, cx+17, cy-8, 7, color.White)

	if enemy {
		ebitenutil.DrawRect(screen, cx-24, cy+18, 48, 5, color.White)
	} else {
		ebitenutil.DrawRect(screen, cx-18, cy+16, 36, 5, color.White)
	}

	ebitenutil.DebugPrintAt(
		screen,
		monster.Name,
		x+18,
		y+158,
	)

	drawHealthBar(screen, monster, x+18, y+185, cardW-36)

	stats := fmt.Sprintf(
		"HP %d/%d   ATK %d   DEF %d\nHeiltränke: %d",
		monster.HP,
		monster.MaxHP,
		monster.Attack,
		monster.Defense,
		monster.Potions,
	)

	ebitenutil.DebugPrintAt(screen, stats, x+18, y+215)
}

func drawHealthBar(screen *ebiten.Image, monster Monster, x, y, width int) {
	const height = 14

	ebitenutil.DrawRect(
		screen,
		float64(x),
		float64(y),
		float64(width),
		height,
		healthBack,
	)

	ratio := 0.0
	if monster.MaxHP > 0 {
		ratio = float64(monster.HP) / float64(monster.MaxHP)
	}
	if ratio < 0 {
		ratio = 0
	}
	if ratio > 1 {
		ratio = 1
	}

	ebitenutil.DrawRect(
		screen,
		float64(x),
		float64(y),
		float64(width)*ratio,
		height,
		healthColor,
	)
}

func drawVersus(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "VS", 442, 235)
}

func drawBattlePanel(screen *ebiten.Image, g *Game) {
	x := 55
	y := 410
	w := 790
	h := 170

	ebitenutil.DrawRect(
		screen,
		float64(x),
		float64(y),
		float64(w),
		float64(h),
		panelColor,
	)

	ebitenutil.DebugPrintAt(screen, g.Message, x+18, y+18)
	ebitenutil.DebugPrintAt(screen, g.SecondLine, x+18, y+40)

	if g.State != PlayerTurn {
		if g.State == EnemyTurn {
			ebitenutil.DebugPrintAt(screen, "Gegner ist am Zug...", x+18, y+78)
		}
		return
	}

	ebitenutil.DebugPrintAt(screen, "DEINE AKTIONEN", x+18, y+74)

	for i, move := range g.Player.Moves {
		line := fmt.Sprintf(
			"%d - %-14s Stärke %d",
			i+1,
			move.Name,
			move.Power,
		)
		ebitenutil.DebugPrintAt(screen, line, x+18+(i%2)*270, y+98+(i/2)*24)
	}

	potionText := fmt.Sprintf("H - Heiltrank (+%d HP)", PotionHeal)
	ebitenutil.DebugPrintAt(screen, potionText, x+545, y+98)

	ebitenutil.DebugPrintAt(
		screen,
		"R - Kampf neu starten",
		x+545,
		y+122,
	)

	_ = textColor
	_ = mutedColor
}
