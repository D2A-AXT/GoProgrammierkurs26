package addresses

// Übung 5 – Die erste Adresse
// Schwierigkeit: ★★☆☆☆
//
// Ab hier kommen Pointer dazu.
//
// Lernziele:
//   - &variable liefert die Adresse einer Variable
//   - *T ist der Typ "Pointer auf T"
//   - Ein Pointer kann auf eine lokale Variable zeigen

// PointerTo soll einen Pointer auf den übergebenen Wert zurückgeben.
//
// Das wirkt zunächst ungewöhnlich:
// `value` ist lokal in dieser Funktion. Trotzdem darf seine Adresse zurückgegeben
// werden. Go stellt sicher, dass der Speicher lange genug lebt.
func PointerTo(value int) *int {
	// TODO: return &value
	return nil
}

// SameAddress soll true liefern, wenn beide Pointer dieselbe Adresse enthalten.
func SameAddress(a *int, b *int) bool {
	// TODO: Pointer können direkt mit == verglichen werden.
	return false
}

// Experiment 1 – Adresse anzeigen:
//
//     value := 42
//     p := &value
//     fmt.Printf("value = %d\n", value)
//     fmt.Printf("Adresse = %p\n", p)
//
// `%p` formatiert eine Pointer-Adresse.
//
// Experiment 2 – Escape Analysis:
// Führe im Terminal in diesem Ordner aus:
//
//     go build -gcflags="-m"
//
// Suche nach einer Meldung wie "moved to heap: value".
// Der Compiler entscheidet per Escape Analysis, ob ein Wert auf dem Stack bleiben
// kann oder auf den Heap muss. Man sollte Stack/Heap in Go normalerweise nicht
// manuell erzwingen.
