package palindrome

// Übung 4 – Palindrom erkennen
// Schwierigkeit: ★★☆☆☆
//
// Lernziele:
//   - Ein Problem selbst in ein kleineres Teilproblem zerlegen
//   - Einen geeigneten Abbruchfall bestimmen
//   - Bedingungen mit Rekursion kombinieren
//
// Ein Palindrom liest sich vorwärts und rückwärts gleich:
//
//     anna
//     otto
//     racecar
//
// Betrachte:
//
//     racecar
//     r     r
//      aceca
//
// Das erste und letzte Zeichen müssen gleich sein.
// Danach bleibt wieder das gleiche, aber kleinere Problem.
//
// Strategie:
//
//     1. Was ist der kleinste Fall?
//     2. Kann ich diesen direkt lösen?
//     3. Wie mache ich das Problem kleiner?
//     4. Was mache ich mit dem rekursiven Ergebnis?

// Für diese Übung betrachten wir nur einfache ASCII-Zeichen.
func IsPalindrome(text string) bool {
	// TODO:
	//
	// Überlege zuerst:
	//     Ist "" ein Palindrom?
	//     Ist "a" ein Palindrom?
	//
	// Hilfreich:
	//     text[0]
	//     text[len(text)-1]
	//     text[1 : len(text)-1]

	return false
}

// Experiment:
//
// Zeichne die Aufrufe für IsPalindrome("anna") und IsPalindrome("hallo").
// Bei welchem Aufruf kann "hallo" bereits sicher abgebrochen werden?
