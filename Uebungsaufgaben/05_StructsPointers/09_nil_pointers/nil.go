package nilpointers

// Übung 9 – nil Pointer sicher behandeln
// Schwierigkeit: ★★★★☆
//
// Lernziele:
//   - Der Zero Value eines Pointers ist nil
//   - Dereferenzieren von nil führt zu einer Panic
//   - Funktionen können nil bewusst als "kein Wert vorhanden" behandeln

type Temperature struct {
	Celsius float64
}

// ReadTemperature soll den Wert und true zurückgeben, wenn t != nil ist.
// Ist t == nil, soll 0 und false zurückgegeben werden.
func ReadTemperature(t *Temperature) (float64, bool) {
	// TODO: Prüfe zuerst t == nil.
	return 0, false
}

// SetTemperature soll false liefern, wenn t nil ist.
// Andernfalls soll Celsius gesetzt und true zurückgegeben werden.
func SetTemperature(t *Temperature, value float64) bool {
	// TODO
	return false
}

// Experiment:
// Entferne die nil-Prüfung testweise und dereferenziere einen nil-Pointer.
// Was passiert?
//
// Beispiel (danach wieder entfernen!):
//     var t *Temperature
//     _ = t.Celsius
//
// Eine Panic ist kein normaler Kontrollfluss. nil sollte vor dem Dereferenzieren
// geprüft werden, wenn nil für die API ein erlaubter Zustand ist.
