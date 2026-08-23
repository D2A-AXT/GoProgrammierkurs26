package mergesort

// Übung 8 – Challenge: Merge Sort
// Schwierigkeit: ★★★★☆
//
// Lernziele:
//   - Divide-and-Conquer vollständig anwenden
//   - Zwei rekursive Teilprobleme lösen
//   - Ergebnisse anschließend kombinieren
//
// Merge Sort arbeitet in drei Schritten:
//
//     1. Teilen
//     2. Rekursiv sortieren
//     3. Zusammenführen
//
// Beispiel:
//
//         [8, 3, 6, 2]
//          /         \
//       [8, 3]      [6, 2]
//       /   \        /   \
//     [8]   [3]    [6]   [2]
//
// Ein Slice mit 0 oder 1 Element ist bereits sortiert.
// Danach werden die sortierten Teilergebnisse wieder zusammengeführt.

func MergeSort(numbers []int) []int {
	// TODO 1:
	// Abbruchfall bestimmen.

	// TODO 2:
	// Slice in zwei Hälften teilen.

	// TODO 3:
	// Beide Hälften rekursiv sortieren.

	// TODO 4:
	// Die sortierten Hälften mit merge() verbinden.

	return nil
}

// merge verbindet zwei bereits sortierte Slices.
// Diese Funktion ist bereits fertig.
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// Experiment:
//
// Zeichne für [4, 1, 3, 2] zuerst nur die rekursiven Aufrufe.
// Zeichne danach den Rückweg und die merge()-Operationen.
//
// Frage:
// Warum ist Merge Sort trotz vieler rekursiver Aufrufe wesentlich effizienter
// als naive Fibonacci-Rekursion?
