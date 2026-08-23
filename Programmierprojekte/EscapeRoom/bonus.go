package main

// BONUS-AUFGABEN
//
// Diese Funktionen werden vom Hauptprogramm nicht benötigt.

// RemoveItem entfernt einen Gegenstand aus dem Inventar.
//
// BONUS 1:
// Nutze FindItemIndex und Slice-Operationen.
func RemoveItem(player *Player, itemID string) bool {
	// TODO
	return false
}

// CountLockedExits zählt alle verschlossenen Ausgänge,
// für die der Spieler den benötigten Gegenstand NICHT besitzt.
//
// BONUS 2:
func CountLockedExits(player Player, room Room) int {
	// TODO
	return 0
}

// CanReachRoom prüft sehr vereinfacht, ob ein Zielraum direkt
// vom aktuellen Raum aus erreichbar ist.
//
// BONUS 3:
// Nur direkte Ausgänge betrachten.
// Keine Rekursion, keine Wegsuche.
func CanReachRoom(player Player, rooms []Room, targetRoom string) bool {
	// TODO
	return false
}
