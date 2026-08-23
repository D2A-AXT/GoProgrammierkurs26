package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	WindowWidth  = 980
	WindowHeight = 720

	RoadCenterX = 490.0
	RoadCenterY = 330.0
	RoadWidth   = 150.0

	EastLaneY  = RoadCenterY + 35
	SouthLaneX = RoadCenterX - 35

	EastStopLine  = RoadCenterX - RoadWidth/2 - 15
	SouthStopLine = RoadCenterY - RoadWidth/2 - 15

	CarLength = 34.0
	CarWidth  = 20.0
	MinGap    = 16.0

	SimulationStep = 0.6
)

type Game struct {
	Cars      []Car
	Phase     TrafficPhase
	Paused    bool
	NextCarID int

	lastPhaseChange time.Time
	lastSpawn       time.Time

	Message string
}

func NewGame() *Game {
	g := &Game{}
	g.Reset()
	return g
}

func (g *Game) Reset() {
	g.Cars = []Car{
		newEastCar(1, 70),
		newEastCar(2, 20),
		newSouthCar(3, 70),
	}
	g.Phase = EastGreen
	g.Paused = false
	g.NextCarID = 4
	g.lastPhaseChange = time.Now()
	g.lastSpawn = time.Now()
	g.Message = "Simulation läuft."
}

func newEastCar(id int, x float64) Car {
	return Car{
		ID:        id,
		X:         x,
		Y:         EastLaneY,
		Direction: East,
		Speed:     2.2,
		Length:    CarLength,
		Width:     CarWidth,
	}
}

func newSouthCar(id int, y float64) Car {
	return Car{
		ID:        id,
		X:         SouthLaneX,
		Y:         y,
		Direction: South,
		Speed:     2.0,
		Length:    CarLength,
		Width:     CarWidth,
	}
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.Reset()
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.Paused = !g.Paused
		if g.Paused {
			g.Message = "Simulation pausiert."
		} else {
			g.Message = "Simulation läuft."
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyN) {
		g.Phase = NextPhase(g.Phase)
		g.lastPhaseChange = time.Now()
		g.Message = "Ampelphase manuell weitergeschaltet."
	}

	if g.Paused {
		return nil
	}

	g.updateTrafficLights()
	g.spawnCars()
	g.moveCars()

	g.Cars = RemoveExitedCars(
		g.Cars,
		WindowWidth+80,
		WindowHeight+80,
	)

	// Schutz für den Fall, dass Aufgabe 7 noch nicht implementiert ist:
	// So wächst die Liste beim Ausprobieren nicht unbegrenzt.
	if len(g.Cars) > 60 {
		g.Cars = append([]Car(nil), g.Cars[len(g.Cars)-60:]...)
	}

	return nil
}

func (g *Game) updateTrafficLights() {
	duration := 5 * time.Second
	if g.Phase == EastYellow || g.Phase == SouthYellow {
		duration = 1500 * time.Millisecond
	}

	if time.Since(g.lastPhaseChange) >= duration {
		g.Phase = NextPhase(g.Phase)
		g.lastPhaseChange = time.Now()
	}
}

func (g *Game) spawnCars() {
	if time.Since(g.lastSpawn) < 1600*time.Millisecond {
		return
	}
	g.lastSpawn = time.Now()

	if rand.Intn(2) == 0 {
		if g.canSpawn(East) {
			g.Cars = append(g.Cars, newEastCar(g.NextCarID, -40))
			g.NextCarID++
		}
	} else {
		if g.canSpawn(South) {
			g.Cars = append(g.Cars, newSouthCar(g.NextCarID, -40))
			g.NextCarID++
		}
	}
}

func (g *Game) canSpawn(direction Direction) bool {
	for _, car := range g.Cars {
		if car.Direction != direction {
			continue
		}

		if direction == East && car.X < 80 {
			return false
		}
		if direction == South && car.Y < 80 {
			return false
		}
	}

	return true
}

func (g *Game) moveCars() {
	// Wir arbeiten hier mit einer Momentaufnahme.
	// So ist die Entscheidung eines Fahrzeugs nicht davon abhängig,
	// ob ein anderes Fahrzeug in dieser Schleife schon bewegt wurde.
	snapshot := append([]Car(nil), g.Cars...)

	for i := range g.Cars {
		car := &g.Cars[i]

		light := LightForDirection(g.Phase, car.Direction)

		stopLine := EastStopLine
		if car.Direction == South {
			stopLine = SouthStopLine
		}

		distance := car.Speed * SimulationStep

		if CanCarMove(
			*car,
			snapshot,
			light,
			stopLine,
			MinGap,
			distance,
		) {
			MoveCar(car, distance)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return WindowWidth, WindowHeight
}

func phaseName(phase TrafficPhase) string {
	switch phase {
	case EastGreen:
		return "Ost-West: GRÜN"
	case EastYellow:
		return "Ost-West: GELB"
	case SouthGreen:
		return "Nord-Süd: GRÜN"
	case SouthYellow:
		return "Nord-Süd: GELB"
	default:
		return fmt.Sprintf("Phase %d", phase)
	}
}
