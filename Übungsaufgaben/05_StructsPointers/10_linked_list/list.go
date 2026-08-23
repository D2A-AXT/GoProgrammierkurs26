package linkedlist

// Übung 10 – Verkettete Liste
// Schwierigkeit: ★★★★★
//
// Lernziele:
//   - Pointer als Verbindungen zwischen Structs
//   - nil als Ende einer Datenstruktur
//   - Eine kleine dynamische Datenstruktur selbst implementieren
//
// Ein Node sieht im Speicher konzeptionell so aus:
//
//   +---------+---------+      +---------+---------+
//   | Value 3 | Next ---+----->| Value 8 | Next --|----> nil
//   +---------+---------+      +---------+---------+
//
// `Next *Node` enthält also nicht den nächsten Node selbst, sondern dessen Adresse.

type Node struct {
	Value int
	Next  *Node
}

// Length zählt die Nodes ab head.
func Length(head *Node) int {
	// TODO:
	// current := head
	// for current != nil { ...; current = current.Next }
	return 0
}

// Sum addiert alle Werte der Liste.
func Sum(head *Node) int {
	// TODO
	return 0
}

// Find gibt einen Pointer auf den ERSTEN Node mit dem gesuchten Wert zurück.
// Wenn der Wert nicht existiert, soll nil zurückgegeben werden.
func Find(head *Node, value int) *Node {
	// TODO
	return nil
}

// Append hängt einen neuen Node ans Ende an und gibt den (möglicherweise neuen)
// Listenanfang zurück.
//
// Sonderfall:
// Ist head == nil, besteht die neue Liste nur aus dem neuen Node.
func Append(head *Node, value int) *Node {
	// TODO
	return nil
}

// Zusatzaufgabe (ohne Test):
// Implementiere anschließend:
//
//     func Prepend(head *Node, value int) *Node
//
// Diese Funktion ist überraschend kurz.
