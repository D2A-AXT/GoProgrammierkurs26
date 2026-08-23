package fibonacci

// Übung 5 – Fibonacci und Laufzeit
// Schwierigkeit: ★★☆☆☆
//
// Lernziele:
//   - Rekursion kann einfach aussehen und trotzdem ineffizient sein
//   - Mehrere rekursive Aufrufe erzeugen einen Aufrufbaum
//   - Laufzeit praktisch messen
//
// Fibonacci:
//
//     0, 1, 1, 2, 3, 5, 8, 13, 21, ...
//
// Definition:
//
//     Fib(0) = 0
//     Fib(1) = 1
//     Fib(n) = Fib(n-1) + Fib(n-2)
//
// Problem: Teilprobleme werden mehrfach berechnet.
//
//             Fib(5)
//            /      \
//        Fib(4)     Fib(3)
//        /   \       /   \
//    Fib(3) Fib(2) Fib(2) Fib(1)
//
// Naive rekursive Fibonacci-Berechnung wächst exponentiell.

func Fib(n int) int {
	// TODO:
	// Implementiere die Definition von oben direkt.
	return 0
}

// Experiment 1 – Laufzeit:
//
// Starte:
//
//     go test -bench=.
//
// Ändere im Benchmark Fib(30) nacheinander zu Fib(35) und Fib(40).
// Beobachte, wie stark die Laufzeit wächst.
//
// Experiment 2 – Anzahl der Aufrufe:
//
// Lege testweise eine globale Variable an:
//
//     var calls int
//
// Erhöhe sie bei jedem Aufruf mit calls++.
// Vergleiche Fib(10), Fib(20) und Fib(30).
//
// Frage:
// Warum ist Fib(30) viel langsamer als SumTo(30), obwohl beide n verkleinern?
