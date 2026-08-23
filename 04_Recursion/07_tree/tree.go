package tree

// Übung 7 – Rekursive Datenstrukturen
// Schwierigkeit: ★★★☆☆
//
// Lernziele:
//   - Eine rekursive Datenstruktur erkennen
//   - Rekursion zur Traversierung eines Baumes verwenden
//   - Structs, Pointer, Slices und Rekursion kombinieren
//
// Beispiel:
//
//             10
//            /  \
//           5    8
//          / \
//         2   3
//
// Ein Node enthält einen Wert und weitere Nodes.
// Ein Teilbaum ist also selbst wieder ein Baum.
// Genau deshalb passt Rekursion hier besonders gut.

type Node struct {
	Value    int
	Children []*Node
}

// SumTree berechnet die Summe aller Werte im Baum.
// Für den Baum oben ist das Ergebnis 28.
func SumTree(node *Node) int {
	// TODO 1:
	// Ein nil-Pointer bedeutet: Hier existiert kein Node.
	// Was sollte dann die Summe sein?

	// TODO 2:
	// Beginne mit dem Wert des aktuellen Nodes.

	// TODO 3:
	// Gehe alle Children durch.
	// Jedes Child ist selbst wieder ein Baum.

	return 0
}

// Experiment:
//
// Importiere testweise "fmt" und ergänze:
//
//     fmt.Println(node.Value)
//
// Wann werden die Nodes ausgegeben?
// Was ändert sich, wenn du erst NACH der Verarbeitung der Kinder ausgibst?
