package memory

// Übung 4 – Speichergröße, Padding und Scope
// Schwierigkeit: ★★★☆☆
//
// Lernziele:
//   - unsafe.Sizeof zum Beobachten der Speichergröße
//   - Structs können wegen Alignment/Padding größer sein als die Summe ihrer Felder
//   - lokale Variablen besitzen einen begrenzten Scope
//
// `unsafe` heißt nicht, dass diese Übung gefährlich ist. Das Package umgeht aber
// einige Sicherheitsgarantien von Go und sollte in normalem Anwendungscode nur
// mit gutem Grund eingesetzt werden.

type BadLayout struct {
	Enabled bool
	Count   int64
	Visible bool
}

type BetterLayout struct {
	Count   int64
	Enabled bool
	Visible bool
}

type SizeInfo struct {
	BoolSize         uintptr
	Int64Size        uintptr
	BadLayoutSize    uintptr
	BetterLayoutSize uintptr
}

// GetSizes soll die Größen der Datentypen/Structs zurückgeben.
func GetSizes() SizeInfo {
	// TODO: Ergänze oben `import "unsafe"` und verwende unsafe.Sizeof(...)
	// Beispiele:
	//   unsafe.Sizeof(bool(false))
	//   unsafe.Sizeof(int64(0))
	//   unsafe.Sizeof(BadLayout{})
	return SizeInfo{}
}

// MakeLabel demonstriert lokalen Scope.
func MakeLabel(name string) string {
	prefix := "Student: " // prefix existiert nur innerhalb dieser Funktion.

	// TODO: Gib prefix + name zurück.
	return prefix // Starterwert: kompiliert, ist aber noch nicht korrekt.
}

// Experiment 1:
// Schreibe in GetSizes vorübergehend weitere Sizeof-Aufrufe für int, int32,
// float64, string und einen beliebigen Struct.
//
// Experiment 2:
// Tausche die Reihenfolge der Felder von BadLayout. Ändert sich die Größe?
// Warum kann ein Struct Padding-Bytes enthalten?
//
// Experiment 3 – Scope:
// Versuche unterhalb von MakeLabel auf `prefix` zuzugreifen.
// Der Code kompiliert nicht. Eine lokale Variable ist nur in ihrem Block sichtbar.
