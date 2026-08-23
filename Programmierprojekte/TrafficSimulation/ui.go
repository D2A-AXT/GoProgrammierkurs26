package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	backgroundColor = color.RGBA{R: 35, G: 105, B: 57, A: 255}
	roadColor       = color.RGBA{R: 55, G: 58, B: 63, A: 255}
	lineColor       = color.RGBA{R: 225, G: 225, B: 205, A: 255}
	panelColor      = color.RGBA{R: 27, G: 31, B: 39, A: 255}
	carEastColor    = color.RGBA{R: 80, G: 160, B: 235, A: 255}
	carSouthColor   = color.RGBA{R: 225, G: 115, B: 78, A: 255}
	redLight        = color.RGBA{R: 225, G: 72, B: 72, A: 255}
	yellowLight     = color.RGBA{R: 235, G: 193, B: 64, A: 255}
	greenLight      = color.RGBA{R: 73, G: 198, B: 105, A: 255}
	offLight        = color.RGBA{R: 76, G: 79, B: 84, A: 255}
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)

	drawRoads(screen)
	drawStopLines(screen)
	drawTrafficLights(screen, g.Phase)
	drawCars(screen, g.Cars)
	drawStatusPanel(screen, g)
}

func drawRoads(screen *ebiten.Image) {
	// Horizontale Straße
	ebitenutil.DrawRect(
		screen,
		0,
		RoadCenterY-RoadWidth/2,
		WindowWidth,
		RoadWidth,
		roadColor,
	)

	// Vertikale Straße
	ebitenutil.DrawRect(
		screen,
		RoadCenterX-RoadWidth/2,
		0,
		RoadWidth,
		WindowHeight-110,
		roadColor,
	)

	// Fahrbahnmarkierungen
	for x := 0.0; x < WindowWidth; x += 45 {
		ebitenutil.DrawRect(
			screen,
			x,
			RoadCenterY-2,
			24,
			4,
			lineColor,
		)
	}

	for y := 0.0; y < WindowHeight-110; y += 45 {
		ebitenutil.DrawRect(
			screen,
			RoadCenterX-2,
			y,
			4,
			24,
			lineColor,
		)
	}
}

func drawStopLines(screen *ebiten.Image) {
	ebitenutil.DrawRect(
		screen,
		EastStopLine,
		EastLaneY-CarWidth,
		5,
		CarWidth*2,
		lineColor,
	)

	ebitenutil.DrawRect(
		screen,
		SouthLaneX-CarWidth,
		SouthStopLine,
		CarWidth*2,
		5,
		lineColor,
	)
}

func drawTrafficLights(screen *ebiten.Image, phase TrafficPhase) {
	eastState := LightForDirection(phase, East)
	southState := LightForDirection(phase, South)

	drawTrafficLight(
		screen,
		int(EastStopLine-55),
		int(EastLaneY-70),
		eastState,
		"Ost",
	)

	drawTrafficLight(
		screen,
		int(SouthLaneX+45),
		int(SouthStopLine-60),
		southState,
		"Süd",
	)
}

func drawTrafficLight(
	screen *ebiten.Image,
	x, y int,
	state LightState,
	label string,
) {
	ebitenutil.DrawRect(
		screen,
		float64(x),
		float64(y),
		36,
		86,
		color.RGBA{R: 25, G: 27, B: 30, A: 255},
	)

	colors := []color.Color{offLight, offLight, offLight}

	switch state {
	case Red:
		colors[0] = redLight
	case Yellow:
		colors[1] = yellowLight
	case Green:
		colors[2] = greenLight
	}

	for i, c := range colors {
		ebitenutil.DrawCircle(
			screen,
			float64(x+18),
			float64(y+15+i*27),
			9,
			c,
		)
	}

	ebitenutil.DebugPrintAt(screen, label, x+3, y+92)
}

func drawCars(screen *ebiten.Image, cars []Car) {
	for _, car := range cars {
		c := carEastColor
		width := car.Length
		height := car.Width

		if car.Direction == South {
			c = carSouthColor
			width = car.Width
			height = car.Length
		}

		ebitenutil.DrawRect(
			screen,
			car.X-width/2,
			car.Y-height/2,
			width,
			height,
			c,
		)

		// Windschutzscheibe
		if car.Direction == East {
			ebitenutil.DrawRect(
				screen,
				car.X+width/2-9,
				car.Y-height/2+3,
				5,
				height-6,
				color.RGBA{R: 190, G: 220, B: 235, A: 255},
			)
		} else {
			ebitenutil.DrawRect(
				screen,
				car.X-width/2+3,
				car.Y+height/2-9,
				width-6,
				5,
				color.RGBA{R: 190, G: 220, B: 235, A: 255},
			)
		}
	}
}

func drawStatusPanel(screen *ebiten.Image, g *Game) {
	y := WindowHeight - 110

	ebitenutil.DrawRect(
		screen,
		0,
		float64(y),
		WindowWidth,
		110,
		panelColor,
	)

	waiting := CountWaitingCars(
		g.Cars,
		g.Phase,
		EastStopLine,
		SouthStopLine,
	)

	state := "LÄUFT"
	if g.Paused {
		state = "PAUSE"
	}

	line1 := fmt.Sprintf(
		"Verkehrssimulation   |   %s   |   Fahrzeuge: %d   |   Wartend: %d   |   %s",
		phaseName(g.Phase),
		len(g.Cars),
		waiting,
		state,
	)

	ebitenutil.DebugPrintAt(screen, line1, 22, y+18)
	ebitenutil.DebugPrintAt(screen, g.Message, 22, y+46)
	ebitenutil.DebugPrintAt(
		screen,
		"SPACE: Pause/Weiter    N: Ampelphase manuell weiter    R: Reset",
		22,
		y+77,
	)
}
