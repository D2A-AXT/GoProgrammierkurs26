package main

// ============================================================
// CONWAY'S GAME OF LIFE - STUDENTENCODE
// ============================================================
//
// In dieser Datei befinden sich die eigentlichen Aufgaben.
// Die grafische Oberfläche ist bereits fertig und befindet sich
// in ui.go und im Ordner engine/.
//
// Ziel:
//   Implementiere die TODOs, bis alle Tests grün sind.
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

// World beschreibt das Spielfeld.
// Cells[row][col] ist true, wenn die Zelle lebt.
type World struct {
	Cells [][]bool
}

// NewWorld erzeugt eine leere Welt mit rows Zeilen und cols Spalten.
// Diese Funktion ist bereits fertig.
func NewWorld(rows, cols int) World {
	cells := make([][]bool, rows)
	for row := range cells {
		cells[row] = make([]bool, cols)
	}

	return World{Cells: cells}
}

// Rows gibt die Anzahl der Zeilen zurück.
// Diese Hilfsfunktion ist bereits fertig.
func (w World) Rows() int {
	return len(w.Cells)
}

// Cols gibt die Anzahl der Spalten zurück.
// Diese Hilfsfunktion ist bereits fertig.
func (w World) Cols() int {
	if len(w.Cells) == 0 {
		return 0
	}
	return len(w.Cells[0])
}

// InBounds prüft, ob eine Position innerhalb der Welt liegt.
// Diese Hilfsfunktion ist bereits fertig.
func (w World) InBounds(row, col int) bool {
	return row >= 0 && row < w.Rows() && col >= 0 && col < w.Cols()
}

// ============================================================
// AUFGABE 1 - Eine Zelle umschalten
// ============================================================
//
// ToggleCell soll den Zustand einer Zelle umdrehen:
//
//	tot     -> lebendig
//	lebendig -> tot
//
// Wenn die Position außerhalb des Spielfelds liegt, soll nichts
// passieren.
//
// Warum ein Pointer?
//
//	Die Funktion soll die übergebene Welt verändern.
//
// Beispiel:
//
//	world.Cells[2][3] == false
//	ToggleCell(&world, 2, 3)
//	world.Cells[2][3] == true
func ToggleCell(world *World, row, col int) {
	// TODO: Implementiere diese Funktion.
}

// ============================================================
// AUFGABE 2 - Lebende Nachbarn zählen
// ============================================================
//
// Eine Zelle besitzt maximal acht Nachbarn:
//
//	X X X
//	X O X
//	X X X
//
// O ist die betrachtete Zelle.
// X sind ihre Nachbarn.
//
// Zellen außerhalb des Spielfelds zählen nicht als Nachbarn.
// Die betrachtete Zelle selbst darf NICHT mitgezählt werden.
//
// Tipp:
//
//	Zwei verschachtelte Schleifen von -1 bis +1 können hilfreich
//	sein. Mit world.InBounds(...) kannst du den Rand prüfen.
func CountNeighbours(world World, row, col int) int {
	// TODO: Implementiere diese Funktion.
	return 0
}

// ============================================================
// AUFGABE 3 - Regeln von Conway implementieren
// ============================================================
//
// Eine lebende Zelle:
//   - stirbt bei weniger als 2 lebenden Nachbarn
//   - überlebt bei 2 oder 3 lebenden Nachbarn
//   - stirbt bei mehr als 3 lebenden Nachbarn
//
// Eine tote Zelle:
//   - wird bei genau 3 lebenden Nachbarn lebendig
//   - bleibt ansonsten tot
func WillBeAlive(currentlyAlive bool, neighbours int) bool {
	// TODO: Implementiere die Regeln.
	return false
}

// ============================================================
// AUFGABE 4 - Nächste Generation berechnen
// ============================================================
//
// Erzeuge eine NEUE Welt und berechne für jede Zelle ihren Zustand
// in der nächsten Generation.
//
// Wichtig:
//
//	Verändere die alte Welt nicht während der Berechnung!
//	Sonst würden bereits veränderte Zellen die Berechnung ihrer
//	Nachbarn beeinflussen.
//
// Vorgehen für jede Zelle:
//  1. Nachbarn zählen
//  2. neuen Zustand mit WillBeAlive(...) bestimmen
//  3. Zustand in die neue Welt schreiben
func NextGeneration(world World) World {
	next := NewWorld(world.Rows(), world.Cols())

	// TODO: Berechne hier die nächste Generation.

	return next
}
