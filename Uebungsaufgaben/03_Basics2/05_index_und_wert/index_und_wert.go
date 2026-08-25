package index_und_wert

// Index und Wert verwenden
//
// Ergänze die TODO-Stellen.
// Die Funktion wird automatisch durch die zugehörige Testdatei geprüft.

import "fmt"

func LabelValues(values []string) []string {
	var result []string
	var i int
	var val string

	// TODO: Erzeuge für jeden Eintrag "Index: Wert".
	_ = fmt.Sprintf

	for i, val = range values {
		result = append(result, fmt.Sprintf("%d: %s", i, val))
	}

	return result
}
