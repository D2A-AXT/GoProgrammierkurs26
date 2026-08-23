package main

// ============================================================
// FREIWILLIGE BONUS-AUFGABEN
// ============================================================
//
// Diese Funktionen werden vom Grundspiel nicht benötigt.
// Sie sind für Teilnehmer gedacht, die mit den Pflichtaufgaben
// früher fertig sind.
// ============================================================

// BONUS 1 - Ball beschleunigen
//
// Multipliziere beide Geschwindigkeitskomponenten mit factor.
// Beispiel:
//
//	Velocity = (5, 2)
//	factor   = 1.1
//
// ergibt ungefähr:
//
//	Velocity = (5.5, 2.2)
func IncreaseBallSpeed(ball *Ball, factor float64) {
	// TODO (Bonus)
}

// BONUS 2 - Einfache Computersteuerung
//
// Bewege das Paddle in Richtung der Y-Position des Balls.
// Verwende dafür MovePaddle(...).
//
// Eine einfache Lösung reicht völlig:
//
//	Ball oberhalb -> -1
//	Ball unterhalb -> +1
//	ungefähr gleiche Höhe -> 0
//
// Überlege dir selbst, wie groß der Bereich sein soll, in dem der
// Computer nicht reagiert. So verhinderst du permanentes Zittern.
func MoveComputerPaddle(paddle *Paddle, ball Ball, fieldHeight float64) {
	// TODO (Bonus)
}
