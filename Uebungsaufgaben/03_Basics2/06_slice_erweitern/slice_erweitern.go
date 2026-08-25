package slice_erweitern

// Slice erweitern
//
// Ergänze die TODO-Stellen.
// Die Funktion wird automatisch durch die zugehörige Testdatei geprüft.

func ExtendNames(names []string, first string, second string) []string {
	// TODO: Füge first und second hinzu.

	names = append(names, first, second)

	return names
}
