package main

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"workshop/pong/engine"
)

const (
	WindowWidth  = 960
	WindowHeight = 640

	FieldWidth  = 900.0
	FieldHeight = 500.0

	fieldOffsetX = 30.0
	fieldOffsetY = 90.0

	winningScore = 5
)

var (
	backgroundColor = color.RGBA{R: 18, G: 20, B: 28, A: 255}
	fieldColor      = color.RGBA{R: 28, G: 31, B: 42, A: 255}
	lineColor       = color.RGBA{R: 92, G: 100, B: 120, A: 255}
	playerColor     = color.RGBA{R: 238, G: 241, B: 247, A: 255}
	ballColor       = color.RGBA{R: 255, G: 201, B: 82, A: 255}
	accentColor     = color.RGBA{R: 103, G: 183, B: 255, A: 255}
)

// PongGame enthält den vollständigen Zustand der Oberfläche.
// Diese Datei ist fertig vorgegeben. Für die Pflichtaufgaben muss sie
// nicht verändert werden.
type PongGame struct {
	LeftPaddle  Paddle
	RightPaddle Paddle
	Ball        Ball

	LeftScore  int
	RightScore int

	Running  bool
	GameOver bool
	ServeTo  ScoringPlayer
}

func NewPongGame() *PongGame {
	game := &PongGame{}
	game.ResetMatch()
	return game
}

func (g *PongGame) ResetMatch() {
	g.LeftScore = 0
	g.RightScore = 0
	g.GameOver = false
	g.Running = false

	g.LeftPaddle = Paddle{
		Position: Vector{X: 35, Y: FieldHeight / 2},
		Width:    14,
		Height:   90,
		Speed:    6,
	}

	g.RightPaddle = Paddle{
		Position: Vector{X: FieldWidth - 35, Y: FieldHeight / 2},
		Width:    14,
		Height:   90,
		Speed:    6,
	}

	g.ServeTo = NoScore
	g.ResetBall(randomHorizontalDirection())
}

func (g *PongGame) ResetBall(horizontalDirection float64) {
	verticalDirection := 2.0
	if rand.Intn(2) == 0 {
		verticalDirection = -verticalDirection
	}

	g.Ball = Ball{
		Position: Vector{X: FieldWidth / 2, Y: FieldHeight / 2},
		Velocity: Vector{X: 5 * horizontalDirection, Y: verticalDirection},
		Radius:   9,
	}
}

func randomHorizontalDirection() float64 {
	if rand.Intn(2) == 0 {
		return -1
	}
	return 1
}

func (g *PongGame) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.ResetMatch()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && !g.GameOver {
		g.Running = !g.Running
	}

	// Paddle links: W / S
	leftDirection := 0.0
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		leftDirection -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		leftDirection += 1
	}
	MovePaddle(&g.LeftPaddle, leftDirection, FieldHeight)

	// Paddle rechts: Pfeiltasten
	rightDirection := 0.0
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		rightDirection -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		rightDirection += 1
	}
	MovePaddle(&g.RightPaddle, rightDirection, FieldHeight)

	if !g.Running || g.GameOver {
		return nil
	}

	MoveBall(&g.Ball)
	BounceOffHorizontalWalls(&g.Ball, FieldHeight)

	if g.Ball.Velocity.X < 0 && HasPaddleCollision(g.Ball, g.LeftPaddle) {
		BounceFromPaddle(&g.Ball, g.LeftPaddle)
	}

	if g.Ball.Velocity.X > 0 && HasPaddleCollision(g.Ball, g.RightPaddle) {
		BounceFromPaddle(&g.Ball, g.RightPaddle)
	}

	scoringPlayer := DetectScore(g.Ball, FieldWidth)
	if scoringPlayer != NoScore {
		g.awardPoint(scoringPlayer)
	}

	return nil
}

func (g *PongGame) awardPoint(player ScoringPlayer) {
	if player == LeftPlayer {
		g.LeftScore++
	} else if player == RightPlayer {
		g.RightScore++
	}

	if g.LeftScore >= winningScore || g.RightScore >= winningScore {
		g.GameOver = true
		g.Running = false
		return
	}

	// Nach einem Punkt pausieren wir kurz logisch: Der Ball wird in
	// die Mitte gesetzt und mit SPACE kann weitergespielt werden.
	g.Running = false

	if player == LeftPlayer {
		g.ResetBall(-1)
	} else {
		g.ResetBall(1)
	}
}

func (g *PongGame) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)

	engine.DrawRect(screen, fieldOffsetX, fieldOffsetY, FieldWidth, FieldHeight, fieldColor)
	engine.DrawRectOutline(screen, fieldOffsetX, fieldOffsetY, FieldWidth, FieldHeight, 2, lineColor)
	engine.DrawDashedVerticalLine(screen, fieldOffsetX+FieldWidth/2, fieldOffsetY, FieldHeight, 14, 12, lineColor)

	engine.DrawCenteredRect(
		screen,
		fieldOffsetX+g.LeftPaddle.Position.X,
		fieldOffsetY+g.LeftPaddle.Position.Y,
		g.LeftPaddle.Width,
		g.LeftPaddle.Height,
		playerColor,
	)

	engine.DrawCenteredRect(
		screen,
		fieldOffsetX+g.RightPaddle.Position.X,
		fieldOffsetY+g.RightPaddle.Position.Y,
		g.RightPaddle.Width,
		g.RightPaddle.Height,
		playerColor,
	)

	engine.DrawCenteredRect(
		screen,
		fieldOffsetX+g.Ball.Position.X,
		fieldOffsetY+g.Ball.Position.Y,
		g.Ball.Radius*2,
		g.Ball.Radius*2,
		ballColor,
	)

	ebitenutil.DebugPrintAt(screen, "PONG - GO PROGRAMMING WORKSHOP", 30, 22)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("PLAYER 1   %d     :     %d   PLAYER 2", g.LeftScore, g.RightScore), 340, 46)
	ebitenutil.DebugPrintAt(screen, "W/S: Player 1     Arrow keys: Player 2     SPACE: Start/Pause     R: Reset", 150, 610)

	status := "PAUSED - press SPACE"
	if g.Running {
		status = "PLAYING"
	}
	if g.GameOver {
		winner := "PLAYER 2"
		if g.LeftScore > g.RightScore {
			winner = "PLAYER 1"
		}
		status = fmt.Sprintf("%s WINS! Press R for a new match", winner)
	}

	ebitenutil.DebugPrintAt(screen, status, 30, 68)

	// Kleine visuelle Markierung, wenn das Spiel aktiv ist.
	if g.Running {
		engine.DrawRect(screen, 18, 66, 6, 6, accentColor)
	}
}

func (g *PongGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return WindowWidth, WindowHeight
}
