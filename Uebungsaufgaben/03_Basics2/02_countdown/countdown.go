package countdown

// Countdown
//
// Ergänze die TODO-Stellen.
// Die Funktion wird automatisch durch die zugehörige Testdatei geprüft.

func Countdown(start int) []int {
	var result []int
	var i int

	// TODO: Zähle von start bis 1 rückwärts.

	for i = start; i > 0; i-- {
		result = append(result, i)
	}

	return result
}
