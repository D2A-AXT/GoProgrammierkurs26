package main

// ============================================================
// PONG - STUDENTENCODE
// ============================================================
//
// In dieser Datei befinden sich die eigentlichen Pflichtaufgaben.
// Fenster, Tastatureingabe, Zeichnen, Spielstand und Game-Loop sind
// bereits fertig und befinden sich in ui.go und engine/.
//
// Ziel:
//   Implementiere die TODOs, bis alle Tests grün sind und Pong
//   vollständig spielbar ist.
//
// Tests starten:
//   go test ./...
//
// Programm starten:
//   go run .
//
// Rekursion wird nicht benötigt.
// ============================================================

// Vector beschreibt eine Position oder Bewegungsrichtung in 2D.
type Vector struct {
	X float64
	Y float64
}

// Paddle beschreibt einen Schläger.
// Position ist die MITTE des Schlägers.
type Paddle struct {
	Position Vector
	Width    float64
	Height   float64
	Speed    float64
}

// Ball beschreibt den Ball.
// Position ist die MITTE des Balls.
type Ball struct {
	Position Vector
	Velocity Vector
	Radius   float64
}

// ScoringPlayer beschreibt, ob und wer einen Punkt erhalten hat.
type ScoringPlayer int

const (
	NoScore ScoringPlayer = iota
	LeftPlayer
	RightPlayer
)

// ============================================================
// AUFGABE 1 - Schläger bewegen
// ============================================================
//
// MovePaddle bewegt den Schläger nach oben oder unten.
//
// direction:
//
//	-1 = nach oben
//	 0 = stehen bleiben
//	+1 = nach unten
//
// Die Position soll um direction * paddle.Speed verändert werden.
// Der Schläger darf das Spielfeld dabei NICHT verlassen.
//
// Wichtig:
//
//	paddle.Position.Y beschreibt die MITTE des Schlägers.
//	Deshalb musst du beim Begrenzen die halbe Höhe berücksichtigen.
//
// Warum ein Pointer?
//
//	Wir wollen die Position des vorhandenen Paddle verändern.
func MovePaddle(paddle *Paddle, direction float64, fieldHeight float64) {
	// TODO: Implementiere diese Funktion.
}

// ============================================================
// AUFGABE 2 - Ball bewegen
// ============================================================
//
// MoveBall bewegt den Ball um genau einen Simulationsschritt.
//
// Neue Position = alte Position + Geschwindigkeit
//
// Beispiel:
//
//	Position = (100, 200)
//	Velocity = (  5,  -2)
//
// danach:
//
//	Position = (105, 198)
func MoveBall(ball *Ball) {
	// TODO: Implementiere diese Funktion.
}

// ============================================================
// AUFGABE 3 - Kollision mit Ober-/Unterkante
// ============================================================
//
// BounceOffHorizontalWalls prüft, ob der Ball die obere oder untere
// Spielfeldkante berührt bzw. überschreitet.
//
// Bei einer Kollision soll:
//  1. der Ball wieder vollständig ins Spielfeld gesetzt werden
//  2. die Y-Geschwindigkeit ihr Vorzeichen wechseln
//
// Gib true zurück, wenn eine Kollision stattgefunden hat.
//
// Das Spielfeld geht in Y-Richtung von 0 bis fieldHeight.
func BounceOffHorizontalWalls(ball *Ball, fieldHeight float64) bool {
	// TODO: Implementiere diese Funktion.
	return false
}

// ============================================================
// AUFGABE 4 - Schlägerkollision erkennen
// ============================================================
//
// HasPaddleCollision gibt true zurück, wenn sich Ball und Schläger
// überlappen.
//
// Für diese Übung behandeln wir den Ball zur Kollisionsprüfung wie
// ein kleines Quadrat mit der Größe 2 * Radius.
//
// Bestimme dafür jeweils:
//   - linke / rechte Kante des Balls
//   - obere / untere Kante des Balls
//   - linke / rechte Kante des Paddles
//   - obere / untere Kante des Paddles
//
// Zwei Rechtecke überlappen sich, wenn sie sich sowohl horizontal
// als auch vertikal überschneiden.
func HasPaddleCollision(ball Ball, paddle Paddle) bool {
	// TODO: Implementiere diese Funktion.
	return false
}

// ============================================================
// AUFGABE 5 - Ball vom Schläger abprallen lassen
// ============================================================
//
// BounceFromPaddle verändert die Geschwindigkeit des Balls nach
// einer Schlägerkollision.
//
// Pflichtteil:
//   - X-Geschwindigkeit umkehren
//
// Zusätzlich soll die Y-Geschwindigkeit davon abhängen, WO der Ball
// den Schläger trifft:
//   - Treffer oberhalb der Mitte -> Ball fliegt stärker nach oben
//   - Treffer in der Mitte      -> kaum Änderung
//   - Treffer unterhalb         -> Ball fliegt stärker nach unten
//
// Vorgegebene Formel:
//
//	relativeHit = (ball.Position.Y - paddle.Position.Y) / (paddle.Height / 2)
//	ball.Velocity.Y += relativeHit * 2.0
//
// Damit bleibt die Mathematik überschaubar und das Spiel fühlt sich
// trotzdem deutlich besser an.
func BounceFromPaddle(ball *Ball, paddle Paddle) {
	// TODO: Implementiere diese Funktion.
}

// ============================================================
// AUFGABE 6 - Punkt erkennen
// ============================================================
//
// DetectScore erkennt, ob der Ball das Spielfeld links oder rechts
// vollständig verlassen hat.
//
// Wenn der Ball LINKS herausfliegt, erhält der rechte Spieler den
// Punkt. Wenn er RECHTS herausfliegt, erhält der linke Spieler den
// Punkt.
//
// Gib ansonsten NoScore zurück.
//
// Das Spielfeld geht in X-Richtung von 0 bis fieldWidth.
func DetectScore(ball Ball, fieldWidth float64) ScoringPlayer {
	// TODO: Implementiere diese Funktion.
	return NoScore
}
