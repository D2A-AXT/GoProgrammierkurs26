package main

// BONUS-AUFGABEN
//
// Diese Funktionen werden vom Hauptprogramm nicht benötigt.

// AverageSpeed berechnet die durchschnittliche Geschwindigkeit
// aller Fahrzeuge.
//
// BONUS 1:
// Gibt bei leerem Slice 0 zurück.
func AverageSpeed(cars []Car) float64 {
	// TODO
	return 0
}

// CountCarsByDirection zählt Fahrzeuge einer bestimmten Richtung.
//
// BONUS 2:
func CountCarsByDirection(cars []Car, direction Direction) int {
	// TODO
	return 0
}

// LongestQueue berechnet, wie viele Fahrzeuge maximal in einer
// der beiden Fahrtrichtungen warten.
//
// BONUS 3:
// Nutze LightForDirection und DistanceToStopLine.
//
// Ein Fahrzeug zählt als wartend, wenn:
// - es vor der Stopplinie steht
// - seine Ampel Rot ist
func LongestQueue(
	cars []Car,
	phase TrafficPhase,
	eastStopLine float64,
	southStopLine float64,
) int {
	// TODO
	return 0
}
