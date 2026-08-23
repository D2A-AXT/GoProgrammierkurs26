package main

// Item beschreibt einen Gegenstand.
type Item struct {
	ID          string
	Name        string
	Description string
}

// Exit beschreibt einen Ausgang aus einem Raum.
//
// TargetRoom ist die ID des Zielraums.
// RequiredItem ist leer, wenn der Ausgang nicht verschlossen ist.
type Exit struct {
	Name         string
	TargetRoom   string
	RequiredItem string
	LockedText   string
}

// Room beschreibt einen Raum des Escape Rooms.
type Room struct {
	ID          string
	Name        string
	Description string
	Items       []Item
	Exits       []Exit
}

// Player beschreibt den aktuellen Zustand des Spielers.
type Player struct {
	CurrentRoom string
	Inventory   []Item
}

// FindRoom sucht einen Raum anhand seiner ID.
//
// AUFGABE 1:
// Durchlaufe den Slice rooms.
// Wenn kein Raum gefunden wird, gib nil zurück.
func FindRoom(rooms []Room, id string) *Room {
	// TODO
	return nil
}

// FindItemIndex sucht einen Gegenstand anhand seiner ID.
//
// Rückgabe:
//
//	0..n = Index
//	-1   = nicht gefunden
//
// AUFGABE 2:
func FindItemIndex(items []Item, id string) int {
	// TODO
	return -1
}

// HasItem prüft, ob der Spieler einen Gegenstand besitzt.
//
// AUFGABE 3:
// Nutze möglichst FindItemIndex.
func HasItem(player Player, itemID string) bool {
	// TODO
	return false
}

// TakeItem nimmt einen Gegenstand aus einem Raum auf.
//
// Regeln:
// - Existiert das Item nicht im Raum, gib false zurück.
// - Entferne es aus room.Items.
// - Füge es player.Inventory hinzu.
// - Gib true zurück.
//
// AUFGABE 4:
// Hier übst du Slices, append und Pointer.
func TakeItem(player *Player, room *Room, itemID string) bool {
	// TODO
	return false
}

// CanUseExit prüft, ob ein Ausgang benutzt werden darf.
//
// Regeln:
// - Ist RequiredItem leer, ist der Ausgang frei.
// - Sonst muss der Spieler den Gegenstand im Inventar besitzen.
//
// AUFGABE 5:
func CanUseExit(player Player, exit Exit) bool {
	// TODO
	return false
}

// FindExit sucht einen Ausgang anhand seines Namens.
//
// Der Name soll exakt verglichen werden.
//
// AUFGABE 6:
func FindExit(room Room, exitName string) *Exit {
	// TODO
	return nil
}

// MovePlayer versucht, einen Ausgang zu benutzen.
//
// Rückgabe:
//
//	moved       = true, wenn der Raum gewechselt wurde
//	message     = Text für die Oberfläche
//
// Regeln:
// - Ausgang existiert nicht -> Fehlermeldung
// - Ausgang verschlossen -> LockedText zurückgeben
// - Ausgang frei -> CurrentRoom verändern
//
// AUFGABE 7:
func MovePlayer(player *Player, rooms []Room, exitName string) (bool, string) {
	// TODO
	return false, ""
}

// IsEscaped prüft, ob der Spieler den Escape Room verlassen hat.
//
// Der Ausgangsraum besitzt die ID:
//
//	"outside"
//
// AUFGABE 8:
func IsEscaped(player Player) bool {
	// TODO
	return false
}
