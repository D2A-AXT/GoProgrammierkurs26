package rectangle

// Übung 2 – Structs und Methoden
// Schwierigkeit: ★★☆☆☆
//
// Lernziele:
//   - Methoden auf Structs
//   - Value Receiver
//   - Berechnungen mit Struct-Feldern

type Rectangle struct {
	Width  float64
	Height float64
}

// Area gibt die Fläche zurück.
//
// Kleine Lektion:
// (r Rectangle) ist ein sogenannter Value Receiver.
// Die Methode bekommt also einen Rectangle-WERT.
func (r Rectangle) Area() float64 {
	// TODO: Width * Height
	return 0
}

// Perimeter gibt den Umfang zurück.
func (r Rectangle) Perimeter() float64 {
	// TODO: 2*Width + 2*Height
	return 0
}

// Scaled soll ein NEUES Rechteck zurückgeben, dessen Seiten mit factor
// multipliziert wurden. Das ursprüngliche Rectangle darf nicht verändert werden.
func (r Rectangle) Scaled(factor float64) Rectangle {
	// TODO: Verändere die Kopie r und gib sie zurück.
	return Rectangle{}
}
