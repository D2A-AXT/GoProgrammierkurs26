package bis_fuenf_zaehlen

// Bis fünf zählen
//
// Ergänze die TODO-Stellen.
// Die Funktion wird automatisch durch die zugehörige Testdatei geprüft.

func CountToFive() []int {
	var result []int
	var x int

	// TODO: Füge mit einer for-Schleife die Zahlen 1 bis 5 hinzu.
	for x = range 5 {
		result = append(result, x+1)
	}
	return result
}
