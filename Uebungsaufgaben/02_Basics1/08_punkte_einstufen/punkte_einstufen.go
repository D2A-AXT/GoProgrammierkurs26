package punkte_einstufen

// Punkte einstufen
//
// Ergänze die TODO-Stellen.
// Die Funktion wird automatisch durch die zugehörige Testdatei geprüft.

func ClassifyPoints(points int) string {
	// TODO: Klassifiziere die Punkte.
	if points >= 90 {
		return "sehr gut"
	} else if points >= 75 {
		return "gut"
	} else if points >= 50 {
		return "bestanden"
	} else {
		return "nicht bestanden"
	}
}
