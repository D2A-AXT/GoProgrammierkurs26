package gerade_zahlen_filtern

// Gerade Zahlen filtern
//
// Ergänze die TODO-Stellen.
// Die Funktion wird automatisch durch die zugehörige Testdatei geprüft.

func FilterEven(numbers []int) []int {
	var evenNumbers []int = []int{}
	//var num int

	// TODO: Füge nur gerade Zahlen hinzu.

	for _, num := range numbers {
		if num%2 == 0 {
			evenNumbers = append(evenNumbers, num)
		}
	}

	return evenNumbers
}
