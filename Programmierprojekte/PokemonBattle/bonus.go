package main

// BONUS-AUFGABEN
//
// Diese Funktionen werden vom Hauptspiel nicht benötigt.
// Sie sind für Teilnehmer gedacht, die früher fertig sind.

// AverageMovePower berechnet die durchschnittliche Stärke aller Attacken.
//
// Gibt 0 zurück, wenn das Monster keine Attacken besitzt.
//
// BONUS 1:
// Hier übst du Schleifen und float64.
func AverageMovePower(monster Monster) float64 {
	// TODO
	return 0
}

// StrongestMove gibt die stärkste Attacke zurück.
//
// BONUS 2:
// Gibt zusätzlich false zurück, wenn keine Attacke vorhanden ist.
//
// Hinweis:
// Mehrere Rückgabewerte:
//
//	move, ok := StrongestMove(monster)
func StrongestMove(monster Monster) (Move, bool) {
	// TODO
	return Move{}, false
}

// SimulateRound führt eine komplette Runde ohne Benutzeroberfläche aus.
//
// BONUS 3:
// 1. player greift enemy an.
// 2. Wenn enemy noch lebt, greift enemy player an.
//
// Diese Funktion eignet sich gut zum Experimentieren mit Pointern.
func SimulateRound(
	player *Monster,
	enemy *Monster,
	playerMove Move,
	enemyMove Move,
) {
	// TODO
}
