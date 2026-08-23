package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	CellSize       = 28
	StatusHeight   = 90
	WindowWidth    = 21 * CellSize
	WindowHeight   = 19*CellSize + StatusHeight
	AutoStepPeriod = 130 * time.Millisecond
)

type Game struct {
	World World
	Robot Robot
	Start Position
	Goal  Position

	Automatic bool
	Won       bool
	Steps     int

	lastAutoStep time.Time
}

func NewGame() *Game {
	world, start, goal := LoadWorld()

	return &Game{
		World: world,
		Robot: Robot{
			Position:  start,
			Direction: Right,
		},
		Start:        start,
		Goal:         goal,
		lastAutoStep: time.Now(),
	}
}

func (g *Game) Reset() {
	g.Robot.Position = g.Start
	g.Robot.Direction = Right
	g.Automatic = false
	g.Won = false
	g.Steps = 0
	g.lastAutoStep = time.Now()
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.Reset()
	}

	if g.Won {
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.Automatic = !g.Automatic
		g.lastAutoStep = time.Now()
	}

	if g.Automatic {
		g.updateAutomatic()
	} else {
		g.updateManual()
	}

	if HasReachedGoal(g.Robot, g.Goal) {
		g.Won = true
		g.Automatic = false
	}

	return nil
}

func (g *Game) updateManual() {
	var direction Direction
	move := false

	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW):
		direction = Up
		move = true
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || inpututil.IsKeyJustPressed(ebiten.KeyD):
		direction = Right
		move = true
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS):
		direction = Down
		move = true
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA):
		direction = Left
		move = true
	}

	if !move {
		return
	}

	oldPosition := g.Robot.Position
	MoveRobot(g.World, &g.Robot, direction)

	if g.Robot.Position != oldPosition {
		g.Steps++
	}
}

func (g *Game) updateAutomatic() {
	if time.Since(g.lastAutoStep) < AutoStepPeriod {
		return
	}
	g.lastAutoStep = time.Now()

	direction := ChooseDirection(g.World, g.Robot)
	oldPosition := g.Robot.Position

	MoveRobot(g.World, &g.Robot, direction)

	if g.Robot.Position != oldPosition {
		g.Steps++
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 20, G: 22, B: 28, A: 255})

	g.drawMaze(screen)
	g.drawGoal(screen)
	g.drawRobot(screen)
	g.drawStatus(screen)
}

func (g *Game) drawMaze(screen *ebiten.Image) {
	for y, row := range g.World.Cells {
		for x, cell := range row {
			px := x * CellSize
			py := y * CellSize

			if cell == '#' {
				ebitenutil.DrawRect(
					screen,
					float64(px+1),
					float64(py+1),
					float64(CellSize-2),
					float64(CellSize-2),
					color.RGBA{R: 65, G: 72, B: 86, A: 255},
				)
			} else {
				ebitenutil.DrawRect(
					screen,
					float64(px+1),
					float64(py+1),
					float64(CellSize-2),
					float64(CellSize-2),
					color.RGBA{R: 29, G: 33, B: 41, A: 255},
				)
			}
		}
	}
}

func (g *Game) drawGoal(screen *ebiten.Image) {
	x := g.Goal.X*CellSize + 5
	y := g.Goal.Y*CellSize + 5

	ebitenutil.DrawRect(
		screen,
		float64(x),
		float64(y),
		float64(CellSize-10),
		float64(CellSize-10),
		color.RGBA{R: 90, G: 190, B: 110, A: 255},
	)
}

func (g *Game) drawRobot(screen *ebiten.Image) {
	centerX := float64(g.Robot.Position.X*CellSize + CellSize/2)
	centerY := float64(g.Robot.Position.Y*CellSize + CellSize/2)

	vector.FillCircle(
		screen,
		float32(centerX),
		float32(centerY),
		float32(CellSize)/2-5,
		color.RGBA{R: 90, G: 165, B: 245, A: 255},
		true,
	)

	// Kleine Richtungsanzeige
	offset := DirectionVector(g.Robot.Direction)
	dirX := centerX + float64(offset.X*7)
	dirY := centerY + float64(offset.Y*7)

	vector.FillCircle(
		screen,
		float32(dirX),
		float32(dirY),
		3,
		color.White,
		true,
	)
}

func (g *Game) drawStatus(screen *ebiten.Image) {
	top := len(g.World.Cells)*CellSize + 12

	mode := "MANUELL"
	if g.Automatic {
		mode = "AUTOMATIK"
	}

	status := fmt.Sprintf(
		"Modus: %s   Schritte: %d\nPfeiltasten/WASD: bewegen   SPACE: Automatik an/aus   R: Reset",
		mode,
		g.Steps,
	)

	if g.Won {
		status = fmt.Sprintf(
			"ZIEL ERREICHT! Schritte: %d\nR: Neustart",
			g.Steps,
		)
	}

	ebitenutil.DebugPrintAt(screen, status, 12, top)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return WindowWidth, WindowHeight
}
