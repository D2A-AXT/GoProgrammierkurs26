package main

// ============================================================
// SNAKE - STUDENTENCODE
// ============================================================
//
// In dieser Datei befinden sich die eigentlichen Pflichtaufgaben.
// Die grafische Oberfläche, Tastatureingabe und Game-Loop sind
// bereits fertig und befinden sich in ui.go und engine/.
//
// Ziel:
//   Implementiere die TODOs, bis alle Tests grün sind und Snake
//   vollständig spielbar ist.
//
// Tests starten:
//   go test ./...
//
// Programm starten:
//   go run .
//
// Tipp:
//   Arbeite Aufgabe für Aufgabe und starte nach jeder Änderung
//   die Tests erneut.
// ============================================================

// Point beschreibt eine Position auf dem Spielfeld.
type Point struct {
	X int
	Y int
}

// Direction beschreibt eine Bewegungsrichtung.
// DX verändert die X-Position, DY die Y-Position.
type Direction struct {
	DX int
	DY int
}

var (
	Up    = Direction{DX: 0, DY: -1}
	Down  = Direction{DX: 0, DY: 1}
	Left  = Direction{DX: -1, DY: 0}
	Right = Direction{DX: 1, DY: 0}
)

// Snake enthält die Schlange.
// Body[0] ist immer der Kopf.
// Die folgenden Einträge sind die Körpersegmente.
type Snake struct {
	Body      []Point
	Direction Direction
}

// NewSnake erzeugt eine neue Schlange.
// Diese Funktion ist bereits fertig.
func NewSnake(start Point, length int, direction Direction) Snake {
	body := make([]Point, length)

	for i := 0; i < length; i++ {
		body[i] = Point{
			X: start.X - direction.DX*i,
			Y: start.Y - direction.DY*i,
		}
	}

	return Snake{
		Body:      body,
		Direction: direction,
	}
}

// Head gibt die Position des Kopfes zurück.
// Diese Hilfsfunktion ist bereits fertig.
func (s Snake) Head() Point {
	if len(s.Body) == 0 {
		return Point{}
	}
	return s.Body[0]
}

// Contains prüft, ob ein Punkt von der Schlange belegt ist.
// Diese Hilfsfunktion ist bereits fertig.
func (s Snake) Contains(point Point) bool {
	for _, part := range s.Body {
		if part == point {
			return true
		}
	}
	return false
}

// ============================================================
// AUFGABE 1 - Richtung ändern
// ============================================================
//
// ChangeDirection soll die Bewegungsrichtung der Schlange ändern.
//
// Eine Schlange darf sich aber NICHT direkt um 180° drehen:
//
//	Rechts -> Links  ist verboten
//	Oben   -> Unten  ist verboten
//
// Warum ein Pointer?
//
//	Wir wollen Direction innerhalb der vorhandenen Snake verändern.
//
// Tipp:
//
//	Zwei Richtungen sind genau dann entgegengesetzt, wenn sich ihre
//	DX-Werte und ihre DY-Werte jeweils zu 0 addieren.
func ChangeDirection(snake *Snake, newDirection Direction) {
	// TODO: Implementiere diese Funktion.
}

// ============================================================
// AUFGABE 2 - Schlange bewegen
// ============================================================
//
// Bewege die Schlange genau ein Feld in ihre aktuelle Richtung.
//
// Beispiel vor der Bewegung nach rechts:
//
//	H X X
//
// danach:
//
//	H X X
//
// H = Kopf
// X = Körper
//
// Die Länge der Schlange soll sich dabei NICHT verändern.
//
// Mögliche Vorgehensweise:
//  1. Berechne die neue Kopfposition.
//  2. Verschiebe die Körperteile von hinten nach vorne.
//  3. Schreibe den neuen Kopf nach Body[0].
func MoveSnake(snake *Snake) {
	// TODO: Implementiere diese Funktion.
}

// ============================================================
// AUFGABE 3 - Wandkollision erkennen
// ============================================================
//
// Das Spielfeld besitzt die Koordinaten:
//
//	X: 0 bis width-1
//	Y: 0 bis height-1
//
// Gib true zurück, wenn sich der KOPF außerhalb dieses Bereichs
// befindet.
func HasWallCollision(snake Snake, width, height int) bool {
	// TODO: Implementiere diese Funktion.
	return false
}

// ============================================================
// AUFGABE 4 - Selbstkollision erkennen
// ============================================================
//
// Gib true zurück, wenn der Kopf dieselbe Position wie eines der
// übrigen Körpersegmente besitzt.
//
// Wichtig:
//
//	Body[0] ist der Kopf selbst und darf nicht mit sich selbst
//	verglichen werden.
func HasSelfCollision(snake Snake) bool {
	// TODO: Implementiere diese Funktion.
	return false
}

// ============================================================
// AUFGABE 5 - Futter erkennen
// ============================================================
//
// Die Schlange frisst das Futter, wenn ihr Kopf auf derselben
// Position liegt wie das Futter.
func HasEatenFood(snake Snake, food Point) bool {
	// TODO: Implementiere diese Funktion.
	return false
}

// ============================================================
// AUFGABE 6 - Schlange wachsen lassen
// ============================================================
//
// Verlängere die Schlange um genau ein Segment.
//
// Dafür kann das letzte Körpersegment ein weiteres Mal an den Slice
// angehängt werden. Beim nächsten MoveSnake-Aufruf wird daraus ein
// sichtbares neues Segment.
//
// Tipp:
//
//	append(...) fügt einem Slice ein neues Element hinzu.
func GrowSnake(snake *Snake) {
	// TODO: Implementiere diese Funktion.
}
