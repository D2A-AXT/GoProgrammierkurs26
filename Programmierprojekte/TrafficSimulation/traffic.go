package main

// Direction beschreibt die Fahrtrichtung eines Fahrzeugs.
type Direction int

const (
	East Direction = iota
	South
)

// LightState beschreibt den Zustand einer Ampel.
type LightState int

const (
	Red LightState = iota
	Yellow
	Green
)

// TrafficPhase beschreibt den Zustand der gesamten Kreuzung.
//
// Die Kreuzung schaltet in dieser Reihenfolge:
//
// EastGreen
//
//	↓
//
// EastYellow
//
//	↓
//
// SouthGreen
//
//	↓
//
// SouthYellow
//
//	↓
//
// EastGreen ...
type TrafficPhase int

const (
	EastGreen TrafficPhase = iota
	EastYellow
	SouthGreen
	SouthYellow
)

// Car beschreibt ein Fahrzeug.
type Car struct {
	ID        int
	X         float64
	Y         float64
	Direction Direction
	Speed     float64
	Length    float64
	Width     float64
}

// NextPhase gibt die nächste Ampelphase zurück.
//
// AUFGABE 1:
// Implementiere den oben beschriebenen Zyklus.
func NextPhase(phase TrafficPhase) TrafficPhase {
	// TODO
	return phase
}

// LightForDirection gibt an, welche Ampelfarbe eine Fahrtrichtung
// in einer bestimmten Phase sieht.
//
// Regeln:
//
// EastGreen:
//
//	East  -> Green
//	South -> Red
//
// EastYellow:
//
//	East  -> Yellow
//	South -> Red
//
// SouthGreen:
//
//	East  -> Red
//	South -> Green
//
// SouthYellow:
//
//	East  -> Red
//	South -> Yellow
//
// AUFGABE 2:
func LightForDirection(phase TrafficPhase, direction Direction) LightState {
	// TODO
	return Red
}

// DistanceToStopLine berechnet den Abstand eines Fahrzeugs zur Stopplinie.
//
// Für East fährt das Fahrzeug nach rechts.
// stopLine ist deshalb eine X-Koordinate.
//
// Für South fährt das Fahrzeug nach unten.
// stopLine ist deshalb eine Y-Koordinate.
//
// Ist das Fahrzeug bereits hinter der Stopplinie, darf das Ergebnis
// negativ sein.
//
// AUFGABE 3:
func DistanceToStopLine(car Car, stopLine float64) float64 {
	// TODO
	return 0
}

// HasSafeDistance prüft den Abstand zu einem vorausfahrenden Fahrzeug.
//
// Regeln:
// - other muss in dieselbe Richtung fahren.
// - other muss vor car fahren.
// - Zwischen den Fahrzeugen soll mindestens minGap Abstand bleiben.
// - Gibt es keinen Konflikt, gib true zurück.
//
// Für East:
//
//	Vergleiche die X-Positionen.
//
// Für South:
//
//	Vergleiche die Y-Positionen.
//
// AUFGABE 4:
func HasSafeDistance(car Car, other Car, minGap float64) bool {
	// TODO
	return true
}

// CanCarMove entscheidet, ob ein Fahrzeug im nächsten Simulationsschritt
// fahren darf.
//
// Regeln:
//
// 1. Wenn ein vorausfahrendes Fahrzeug zu nah ist -> false.
//
//  2. Ist das Fahrzeug bereits HINTER der Stopplinie -> true.
//     Es soll die Kreuzung vollständig verlassen.
//
// 3. Vor der Stopplinie:
//   - Green  -> fahren
//   - Yellow -> nur fahren, wenn das Fahrzeug die Stopplinie
//     mit dem nächsten Schritt erreichen/überfahren würde.
//     Sonst anhalten.
//   - Red    -> anhalten.
//
// Hinweis:
// movementPerStep ist die Strecke, die das Fahrzeug im nächsten Schritt
// zurücklegen würde.
//
// AUFGABE 5:
func CanCarMove(
	car Car,
	cars []Car,
	light LightState,
	stopLine float64,
	minGap float64,
	movementPerStep float64,
) bool {
	// TODO
	return false
}

// MoveCar bewegt ein Fahrzeug genau einen Schritt.
//
// East:
//
//	X erhöhen
//
// South:
//
//	Y erhöhen
//
// AUFGABE 6:
// Verwende einen Pointer, damit das Originalfahrzeug verändert wird.
func MoveCar(car *Car, distance float64) {
	// TODO
}

// RemoveExitedCars entfernt Fahrzeuge, die das sichtbare Spielfeld
// vollständig verlassen haben.
//
// Ein Fahrzeug gilt als ausgefahren, wenn:
//
// East:
//
//	car.X > maxX
//
// South:
//
//	car.Y > maxY
//
// AUFGABE 7:
// Erzeuge einen neuen Slice mit allen Fahrzeugen, die noch sichtbar sind.
func RemoveExitedCars(cars []Car, maxX float64, maxY float64) []Car {
	// TODO
	return cars
}

// CountWaitingCars zählt alle Fahrzeuge, die VOR ihrer Stopplinie stehen
// und deren Ampel gerade Rot ist.
//
// AUFGABE 8:
// Diese Funktion wird in der Statusanzeige verwendet.
func CountWaitingCars(
	cars []Car,
	phase TrafficPhase,
	eastStopLine float64,
	southStopLine float64,
) int {
	// TODO
	return 0
}
