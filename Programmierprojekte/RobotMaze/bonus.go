package main

// BONUS-AUFGABEN
//
// Diese Funktionen werden vom Hauptprogramm nicht benötigt.
// Sie sind für Teilnehmer gedacht, die mit den Pflichtaufgaben früher fertig sind.

// CountOpenNeighbours zählt, wie viele direkte Nachbarfelder frei sind.
//
// Bonus 1:
// Nutze CanMove für alle vier Richtungen.
func CountOpenNeighbours(world World, robot Robot) int {
	// TODO
	return 0
}

// ManhattanDistance gibt die Manhattan-Distanz zwischen zwei Positionen zurück.
//
// Beispiel:
// a = (2, 3)
// b = (5, 7)
// Ergebnis = |5-2| + |7-3| = 7
//
// Bonus 2:
// Implementiere die Funktion ohne math.Abs.
// Hinweis: Du kannst die Differenzen selbst positiv machen.
func ManhattanDistance(a Position, b Position) int {
	// TODO
	return 0
}

// ChooseDirectionTowardsGoal ist eine alternative einfache Strategie.
//
// Bonus 3:
// Probiere zuerst freie Richtungen aus, die die Manhattan-Distanz
// zum Ziel verkleinern.
//
// Achtung:
// Diese Strategie garantiert NICHT, jedes Labyrinth zu lösen.
// Genau das macht sie als Experiment interessant.
func ChooseDirectionTowardsGoal(world World, robot Robot, goal Position) Direction {
	// TODO
	return robot.Direction
}
