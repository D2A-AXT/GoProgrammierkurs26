package factorial

// Übung 2 – Fakultät und der Call Stack
// Schwierigkeit: ★☆☆☆☆
//
// Lernziele:
//   - Eine mathematische Definition rekursiv umsetzen
//   - Verstehen, dass Funktionsaufrufe aufeinander warten
//   - Unterschied zwischen Aufruf und Rückgabe beobachten
//
// Die Fakultät ist definiert als:
//
//     5! = 5 * 4 * 3 * 2 * 1
//
// Rekursiv:
//
//     n! = n * (n-1)!
//
// Abbruchfall:
//
//     0! = 1
//
// Beim Aufruf von Factorial(4) entsteht auf dem Call Stack ungefähr:
//
//     Factorial(4)
//         Factorial(3)
//             Factorial(2)
//                 Factorial(1)
//                     Factorial(0)
//
// Danach kommen die Ergebnisse in umgekehrter Reihenfolge zurück.

func Factorial(n int) int {
	// TODO:
	// Implementiere Abbruchfall und rekursiven Fall.
	return 0
}

// Experiment:
//
// Importiere testweise "fmt" und gib direkt beim Eintritt in die Funktion aus:
//
//     fmt.Println("Start:", n)
//
// Gib außerdem kurz vor der Rückgabe aus:
//
//     fmt.Println("Ende:", n)
//
// Dafür musst du das rekursive Ergebnis eventuell zuerst in einer Variablen speichern.
// Warum ist die Reihenfolge beim Zurückkehren umgekehrt?
