package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"workshop/gameoflife/engine"
)

const (
	ScreenWidth  = 980
	ScreenHeight = 650

	WorldRows = 34
	WorldCols = 48

	CellSize = 14
	GridX    = 24
	GridY    = 92

	GenerationDelayTicks = 8
)

var (
	backgroundColor = color.RGBA{R: 22, G: 25, B: 31, A: 255}
	panelColor      = color.RGBA{R: 31, G: 35, B: 43, A: 255}
	aliveColor      = color.RGBA{R: 102, G: 205, B: 170, A: 255}
	deadColor       = color.RGBA{R: 44, G: 49, B: 60, A: 255}
	gridColor       = color.RGBA{R: 63, G: 69, B: 82, A: 255}
	textColor       = color.RGBA{R: 235, G: 238, B: 243, A: 255}
	mutedTextColor  = color.RGBA{R: 167, G: 174, B: 188, A: 255}
	accentColor     = color.RGBA{R: 120, G: 160, B: 255, A: 255}
)

type Game struct {
	World        World
	Running      bool
	Generation   int
	updateTicks  int
	initialWorld World
}

func NewGame() *Game {
	initial := InitialWorld(WorldRows, WorldCols)

	return &Game{
		World:        cloneWorld(initial),
		initialWorld: initial,
	}
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.Running = !g.Running
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyN) {
		g.step()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.World = cloneWorld(g.initialWorld)
		g.Generation = 0
		g.Running = false
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		g.World = NewWorld(WorldRows, WorldCols)
		g.Generation = 0
		g.Running = false
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		row, col, ok := engine.CellAt(x, y, GridX, GridY, CellSize, WorldRows, WorldCols)
		if ok {
			ToggleCell(&g.World, row, col)
			g.Running = false
		}
	}

	if g.Running {
		g.updateTicks++
		if g.updateTicks >= GenerationDelayTicks {
			g.step()
			g.updateTicks = 0
		}
	}

	return nil
}

func (g *Game) step() {
	g.World = NextGeneration(g.World)
	g.Generation++
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)

	ebitenutil.DebugPrintAt(screen, "CONWAY'S GAME OF LIFE", GridX, 24)
	ebitenutil.DebugPrintAt(screen, "Programming Workshop - Go", GridX, 46)

	engine.DrawGrid(
		screen,
		GridX,
		GridY,
		CellSize,
		WorldRows,
		WorldCols,
		func(row, col int) bool { return g.World.Cells[row][col] },
		aliveColor,
		deadColor,
		gridColor,
	)

	panelX := GridX + WorldCols*CellSize + 28
	panelY := GridY
	panelW := ScreenWidth - panelX - 24
	panelH := WorldRows * CellSize
	engine.DrawPanel(screen, panelX, panelY, panelW, panelH, panelColor)

	status := "PAUSIERT"
	statusColor := mutedTextColor
	if g.Running {
		status = "LAEUFT"
		statusColor = aliveColor
	}

	engine.DrawLabel(screen, panelX+18, panelY+20, "STATUS", mutedTextColor)
	engine.DrawLabel(screen, panelX+18, panelY+42, status, statusColor)

	engine.DrawLabel(screen, panelX+18, panelY+82, "GENERATION", mutedTextColor)
	engine.DrawLabel(screen, panelX+18, panelY+104, fmt.Sprintf("%d", g.Generation), textColor)

	engine.DrawLabel(screen, panelX+18, panelY+150, "STEUERUNG", accentColor)
	engine.DrawLabel(screen, panelX+18, panelY+178, "SPACE  Start / Pause", textColor)
	engine.DrawLabel(screen, panelX+18, panelY+200, "N      Einzelschritt", textColor)
	engine.DrawLabel(screen, panelX+18, panelY+222, "R      Muster laden", textColor)
	engine.DrawLabel(screen, panelX+18, panelY+244, "C      Leeren", textColor)
	engine.DrawLabel(screen, panelX+18, panelY+266, "Maus   Zelle aendern", textColor)

	engine.DrawLabel(screen, panelX+18, panelY+320, "TIPP", accentColor)
	engine.DrawWrappedText(
		screen,
		panelX+18,
		panelY+344,
		panelW-36,
		"Wenn die Simulation sofort ausstirbt, sind die TODOs in game.go noch nicht vollstaendig.",
		mutedTextColor,
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func cloneWorld(world World) World {
	clone := NewWorld(world.Rows(), world.Cols())
	for row := range world.Cells {
		copy(clone.Cells[row], world.Cells[row])
	}
	return clone
}
