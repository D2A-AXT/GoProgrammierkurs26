package binarysearch

// Übung 6 – Rekursive Binärsuche
// Schwierigkeit: ★★★☆☆
//
// Lernziele:
//   - Ein Problem nicht nur um 1, sondern stark verkleinern
//   - Divide-and-Conquer verstehen
//   - Zusammenhang zwischen Rekursion und Laufzeit erkennen
//
// Voraussetzung: Das Slice ist SORTIERT.
//
// Beispiel:
//
//     [2, 4, 7, 10, 15, 20, 31]
//
// Gesucht: 20
// Mitte: 10
//
// Weil 20 > 10 ist, kann die linke Hälfte sofort verworfen werden.
// Bei jedem Schritt bleibt ungefähr nur die Hälfte übrig.

func BinarySearch(numbers []int, target int) int {
	return search(numbers, target, 0)
}

// offset gibt an, an welcher Position das aktuelle Teilslice
// im ursprünglichen Slice begonnen hat.
func search(numbers []int, target int, offset int) int {
	// TODO 1:
	// Abbruchfall: Was passiert, wenn numbers leer ist?

	// TODO 2:
	// Bestimme die Mitte:
	//     middle := len(numbers) / 2

	// TODO 3:
	// Vergleiche numbers[middle] mit target.

	// TODO 4:
	// Suche nur noch links ODER rechts weiter.
	// Bei der rechten Hälfte muss der offset angepasst werden.

	return -1
}

// Experiment:
//
// Wie viele Schritte benötigt die Binärsuche maximal ungefähr bei
// 8, 16 und 1024 Elementen?
//
// Schreibe die Folge auf:
//     1024 -> 512 -> 256 -> ...
//
// Vergleiche das mit einer linearen Suche.
