package pointerreceivers

// Übung 8 – Value Receiver vs. Pointer Receiver
// Schwierigkeit: ★★★☆☆
//
// Lernziele:
//   - Methoden mit (c Counter) arbeiten auf einer Kopie
//   - Methoden mit (c *Counter) können den Originalwert verändern
//   - Entscheidung, wann ein Pointer Receiver sinnvoll ist

type Counter struct {
	Value int
}

// Added verwendet absichtlich einen VALUE RECEIVER.
// Die Methode soll eine veränderte KOPIE zurückgeben.
func (c Counter) Added(amount int) Counter {
	// TODO
	return Counter{}
}

// Add verwendet absichtlich einen POINTER RECEIVER.
// Die Methode soll den ORIGINALEN Counter verändern.
func (c *Counter) Add(amount int) {
	// TODO
}

// Reset setzt den ORIGINALEN Counter auf 0.
func (c *Counter) Reset() {
	// TODO
}
