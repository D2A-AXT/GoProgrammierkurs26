package main

// Position beschreibt eine Position im Labyrinth.
type Position struct {
	X int
	Y int
}

// Direction beschreibt eine Bewegungsrichtung.
type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

// Robot enthält den aktuellen Zustand des Roboters.
//
// Direction ist die Richtung, in die der Roboter aktuell schaut.
type Robot struct {
	Position  Position
	Direction Direction
}

// World enthält das Labyrinth.
//
// '#' = Wand
// ' ' = freies Feld
// 'S' = Start
// 'G' = Ziel
type World struct {
	Cells [][]rune
}

// DirectionVector wandelt eine Richtung in eine Positionsänderung um.
//
// Beispiele:
// Up    -> {X: 0,  Y: -1}
// Right -> {X: 1,  Y: 0}
//
// AUFGABE 1:
// Implementiere die vier Richtungen.
func DirectionVector(direction Direction) Position {
	// TODO
	return Position{}
}

// IsInside prüft, ob sich eine Position innerhalb des Spielfelds befindet.
//
// AUFGABE 2:
// Verwende die Breite und Höhe der World.
func IsInside(world World, position Position) bool {
	// TODO
	return false
}

// IsWall prüft, ob sich an einer Position eine Wand befindet.
//
// Positionen außerhalb des Spielfelds sollen ebenfalls als Wand gelten.
//
// AUFGABE 3:
// Nutze IsInside, bevor du auf world.Cells zugreifst.
func IsWall(world World, position Position) bool {
	// TODO
	return true
}

// CanMove prüft, ob der Roboter einen Schritt in die gewünschte Richtung
// machen darf.
//
// AUFGABE 4:
// Berechne zuerst die Zielposition und prüfe anschließend, ob dort eine Wand ist.
func CanMove(world World, robot Robot, direction Direction) bool {
	// TODO
	return false
}

// MoveRobot bewegt den Roboter um genau ein Feld.
//
// Wichtig:
// - Der Roboter soll immer in die gewünschte Richtung schauen.
// - Bei einer Wand soll er sich NICHT bewegen.
// - Die Funktion verändert den übergebenen Robot direkt.
//
// AUFGABE 5:
// Hier kommen Structs und Pointer zusammen.
func MoveRobot(world World, robot *Robot, direction Direction) {
	// TODO
}

// HasReachedGoal prüft, ob der Roboter das Ziel erreicht hat.
//
// AUFGABE 6:
func HasReachedGoal(robot Robot, goal Position) bool {
	// TODO
	return false
}

// TurnRight gibt die Richtung rechts vom Roboter zurück.
//
// Beispiel:
// Up -> Right -> Down -> Left -> Up
func TurnRight(direction Direction) Direction {
	return (direction + 1) % 4
}

// TurnLeft gibt die Richtung links vom Roboter zurück.
func TurnLeft(direction Direction) Direction {
	return (direction + 3) % 4
}

// TurnBack dreht den Roboter um 180 Grad.
func TurnBack(direction Direction) Direction {
	return (direction + 2) % 4
}

// ChooseDirection entscheidet, wohin der Roboter im Automatikmodus fährt.
//
// Verwende die Rechte-Hand-Regel:
// 1. Wenn rechts frei ist: rechts abbiegen
// 2. Sonst wenn geradeaus frei ist: geradeaus
// 3. Sonst wenn links frei ist: links abbiegen
// 4. Sonst: umdrehen
//
// AUFGABE 7:
// Diese Aufgabe verbindet die bisherigen Funktionen.
// Sie benötigt KEINE Rekursion und KEINE Wegsuche.
func ChooseDirection(world World, robot Robot) Direction {
	// TODO
	return robot.Direction
}
