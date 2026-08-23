package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"workshop/snake/engine"
)

const (
	ScreenWidth  = 1040
	ScreenHeight = 660

	BoardCols = 32
	BoardRows = 24

	CellSize = 20
	BoardX   = 24
	BoardY   = 92

	MoveDelayTicks = 7
)

var (
	backgroundColor = color.RGBA{R: 22, G: 25, B: 31, A: 255}
	panelColor      = color.RGBA{R: 31, G: 35, B: 43, A: 255}
	boardColor      = color.RGBA{R: 38, G: 43, B: 52, A: 255}
	gridColor       = color.RGBA{R: 53, G: 59, B: 70, A: 255}
	snakeColor      = color.RGBA{R: 97, G: 201, B: 126, A: 255}
	headColor       = color.RGBA{R: 148, G: 230, B: 166, A: 255}
	foodColor       = color.RGBA{R: 235, G: 91, B: 91, A: 255}
	textColor       = color.RGBA{R: 235, G: 238, B: 243, A: 255}
	mutedTextColor  = color.RGBA{R: 167, G: 174, B: 188, A: 255}
	accentColor     = color.RGBA{R: 120, G: 160, B: 255, A: 255}
)

type Game struct {
	Snake     Snake
	Food      Point
	Score     int
	Running   bool
	GameOver  bool
	moveTicks int
	rng       *rand.Rand
}

func NewGame() *Game {
	game := &Game{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	game.reset()
	return game
}

func (g *Game) reset() {
	g.Snake = NewSnake(
		Point{X: BoardCols / 2, Y: BoardRows / 2},
		4,
		Right,
	)
	g.Score = 0
	g.Running = false
	g.GameOver = false
	g.moveTicks = 0
	g.Food = randomFreePoint(g.Snake, BoardCols, BoardRows, g.rng)
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.reset()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && !g.GameOver {
		g.Running = !g.Running
	}

	// Pfeiltasten und WASD werden unterstützt.
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		ChangeDirection(&g.Snake, Up)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		ChangeDirection(&g.Snake, Down)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
		ChangeDirection(&g.Snake, Left)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
		ChangeDirection(&g.Snake, Right)
	}

	if g.Running && !g.GameOver {
		g.moveTicks++
		if g.moveTicks >= MoveDelayTicks {
			g.step()
			g.moveTicks = 0
		}
	}

	return nil
}

func (g *Game) step() {
	MoveSnake(&g.Snake)

	if HasWallCollision(g.Snake, BoardCols, BoardRows) || HasSelfCollision(g.Snake) {
		g.GameOver = true
		g.Running = false
		return
	}

	if HasEatenFood(g.Snake, g.Food) {
		GrowSnake(&g.Snake)
		g.Score++
		g.Food = randomFreePoint(g.Snake, BoardCols, BoardRows, g.rng)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)

	ebitenutil.DebugPrintAt(screen, "SNAKE", BoardX, 24)
	ebitenutil.DebugPrintAt(screen, "Programming Workshop - Go", BoardX, 46)

	engine.DrawBoard(
		screen,
		BoardX,
		BoardY,
		CellSize,
		BoardRows,
		BoardCols,
		boardColor,
		gridColor,
	)

	engine.DrawCell(screen, BoardX, BoardY, CellSize, g.Food.X, g.Food.Y, foodColor)

	for i := len(g.Snake.Body) - 1; i >= 0; i-- {
		part := g.Snake.Body[i]
		partColor := snakeColor
		if i == 0 {
			partColor = headColor
		}
		engine.DrawCell(screen, BoardX, BoardY, CellSize, part.X, part.Y, partColor)
	}

	panelX := BoardX + BoardCols*CellSize + 28
	panelY := BoardY
	panelW := ScreenWidth - panelX - 24
	panelH := BoardRows * CellSize
	engine.DrawPanel(screen, panelX, panelY, panelW, panelH, panelColor)

	status := "PAUSIERT"
	if g.Running {
		status = "LAEUFT"
	}
	if g.GameOver {
		status = "GAME OVER"
	}

	engine.DrawLabel(screen, panelX+18, panelY+20, "STATUS", mutedTextColor)
	engine.DrawLabel(screen, panelX+18, panelY+42, status, textColor)

	engine.DrawLabel(screen, panelX+18, panelY+82, "PUNKTE", mutedTextColor)
	engine.DrawLabel(screen, panelX+18, panelY+104, fmt.Sprintf("%d", g.Score), textColor)

	engine.DrawLabel(screen, panelX+18, panelY+144, "LAENGE", mutedTextColor)
	engine.DrawLabel(screen, panelX+18, panelY+166, fmt.Sprintf("%d", len(g.Snake.Body)), textColor)

	engine.DrawLabel(screen, panelX+18, panelY+210, "STEUERUNG", accentColor)
	engine.DrawLabel(screen, panelX+18, panelY+238, "SPACE      Start / Pause", textColor)
	engine.DrawLabel(screen, panelX+18, panelY+260, "Pfeile/WASD Richtung", textColor)
	engine.DrawLabel(screen, panelX+18, panelY+282, "R          Neustart", textColor)

	engine.DrawLabel(screen, panelX+18, panelY+330, "TIPP", accentColor)
	engine.DrawWrappedText(
		screen,
		panelX+18,
		panelY+354,
		panelW-36,
		"Wenn sich die Schlange noch nicht bewegt, beginne mit den TODOs in game.go.",
		mutedTextColor,
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func randomFreePoint(snake Snake, width, height int, rng *rand.Rand) Point {
	freeCells := make([]Point, 0, width*height-len(snake.Body))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			point := Point{X: x, Y: y}
			if !snake.Contains(point) {
				freeCells = append(freeCells, point)
			}
		}
	}

	if len(freeCells) == 0 {
		return Point{}
	}

	return freeCells[rng.Intn(len(freeCells))]
}
