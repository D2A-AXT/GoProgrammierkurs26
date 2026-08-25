package text_formatieren

// Text mit fmt formatieren
//
// Ergänze die TODO-Stellen.
// Die Funktion wird automatisch durch die zugehörige Testdatei geprüft.

import "fmt"

func FormatCourse(course string, participants int) string {
	// TODO: Erzeuge den erwarteten Text mit fmt.Sprintf.
	x := fmt.Sprintf("%s hat %d Teilnehmer.", course, participants)
	return x
}
