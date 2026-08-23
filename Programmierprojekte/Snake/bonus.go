package main

// ============================================================
// FREIWILLIGE BONUS-AUFGABEN
// ============================================================
//
// Diese Funktionen werden für das Grundspiel nicht benötigt.
// Wenn du mit den Pflichtaufgaben fertig bist, kannst du sie als
// Ausgangspunkt für Erweiterungen verwenden.
// ============================================================

// CountFreeCells soll zählen, wie viele Felder des Spielfelds NICHT
// von der Schlange belegt sind.
func CountFreeCells(snake Snake, width, height int) int {
	// BONUS TODO
	return 0
}

// WrapPoint soll eine Position auf die gegenüberliegende Seite des
// Spielfelds setzen, wenn sie den Rand verlässt.
//
// Beispiel bei width = 10:
//
//	X = -1  -> X = 9
//	X = 10  -> X = 0
//
// Damit kann später ein Spielmodus ohne tödliche Wände gebaut werden.
func WrapPoint(point Point, width, height int) Point {
	// BONUS TODO
	return point
}
