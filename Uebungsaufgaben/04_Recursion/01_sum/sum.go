package sum

// Übung 1 – Erste Rekursion: Summe von 1 bis n
// Schwierigkeit: ★☆☆☆☆
//
// Lernziele:
//   - Eine rekursive Funktion ruft sich selbst auf
//   - Jede Rekursion braucht einen Abbruchfall
//   - Jeder rekursive Aufruf muss sich dem Abbruchfall nähern
//
// Aus der Vorlesung:
//
// Eine rekursive Funktion besteht normalerweise aus:
//
//     1. Abbruchfall
//     2. Rekursivem Fall
//
// Beispiel:
//
//     SumTo(4)
//     = 4 + SumTo(3)
//     = 4 + 3 + SumTo(2)
//     = 4 + 3 + 2 + SumTo(1)
//     = 4 + 3 + 2 + 1
//     = 10
//
// Der Parameter wird bei jedem Aufruf kleiner.
// Dadurch erreichen wir irgendwann den Abbruchfall.

// SumTo berechnet die Summe aller Zahlen von 1 bis n.
//
// Für n <= 0 soll 0 zurückgegeben werden.
func SumTo(n int) int {

	// TODO 1:
	// Implementiere zuerst den Abbruchfall.
	//
	// Frage:
	// Was ist das Ergebnis von SumTo(0)?

	// TODO 2:
	// Implementiere danach den rekursiven Fall.
	//
	// Hinweis:
	//
	//     SumTo(n) = n + SumTo(n - 1)

	return 0
}

// Experiment:
//
// Füge am Anfang von SumTo testweise folgende Ausgabe ein:
//
//     fmt.Println("SumTo aufgerufen mit:", n)
//
// Rufe danach SumTo(5) auf.
//
// In welcher Reihenfolge werden die Funktionen aufgerufen?
//
// Erwartung:
//
//     SumTo(5)
//     SumTo(4)
//     SumTo(3)
//     SumTo(2)
//     SumTo(1)
//     SumTo(0)
//
// Entferne danach die Ausgabe wieder.