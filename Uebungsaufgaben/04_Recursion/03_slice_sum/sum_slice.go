package slicesum

// Übung 3 – Rekursion mit Slices
// Schwierigkeit: ★★☆☆☆
//
// Lernziele:
//   - Nicht nur Zahlen können rekursiv verkleinert werden
//   - Ein Slice kann in "erstes Element" und "Rest" zerlegt werden
//   - Den Abbruchfall aus der Datenstruktur ableiten
//
// Beispiel:
//
//     [4, 7, 2]
//       4 + SumSlice([7, 2])
//             7 + SumSlice([2])
//                   2 + SumSlice([])
//
// Ein leeres Slice ist hier der Abbruchfall.
//
// Hilfreiche Go-Syntax:
//
//     numbers[0]   // erstes Element
//     numbers[1:]  // alles außer dem ersten Element

func SumSlice(numbers []int) int {
	// TODO 1:
	// Was soll bei einem leeren Slice zurückgegeben werden?

	// TODO 2:
	// Addiere das erste Element zum Ergebnis des restlichen Slices.

	return 0
}

// Experiment:
//
// Schreibe die einzelnen Funktionsaufrufe für SumSlice([]int{42}) auf Papier.
// Was würde passieren, wenn der Abbruchfall fehlt?
