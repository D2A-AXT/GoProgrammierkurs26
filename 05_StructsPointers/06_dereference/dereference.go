package dereference

// Übung 6 – Dereferenzieren und verändern
// Schwierigkeit: ★★☆☆☆
//
// Lernziele:
//   - *pointer liest den Wert an einer Adresse
//   - *pointer = ... verändert den Wert an dieser Adresse
//   - Eine Funktion kann dadurch Daten des Aufrufers verändern

// Read soll den Wert lesen, auf den p zeigt.
func Read(p *int) int {
	// TODO: return *p
	return 0
}

// Write soll den Wert an der Adresse p auf newValue setzen.
func Write(p *int, newValue int) {
	// TODO: *p = newValue
}

// Add soll delta zum Wert addieren, auf den p zeigt.
func Add(p *int, delta int) {
	// TODO
}

// Merksatz:
// Go ist trotzdem IMMER Pass by Value.
// Beim Aufruf Write(&x, 5) wird der Pointer (die Adresse) als Wert kopiert.
// Beide Pointer-Werte zeigen danach aber auf denselben Speicher.
