package main

import "fmt"

// ============================================================
// GO PLAYGROUND
// ============================================================
//
// Ziel:
// Probiert Dinge aus, verändert Code, führt ihn aus und schaut,
// was passiert.
//
// Starten:
//   go run .
//
// Oben in main() könnt ihr auswählen, welche Aufgabe ihr starten wollt.
//
// Wichtig:
// Fehler sind erlaubt und sogar erwünscht.
// Wenn etwas nicht funktioniert:
//   1. Fehlermeldung lesen
//   2. Vermutung aufstellen
//   3. Etwas ändern
//   4. Noch einmal ausführen
//
// ============================================================

func main() {
	// Ändert diese Zahl, um eine andere Aufgabe zu starten.
	exercise := 1

	switch exercise {
	case 1:
		aufgabe01()
	case 2:
		aufgabe02()
	case 3:
		aufgabe03()
	case 4:
		aufgabe04()
	case 5:
		aufgabe05()
	case 6:
		aufgabe06()
	case 7:
		aufgabe07()
	case 8:
		aufgabe08()
	case 9:
		aufgabe09()
	case 10:
		aufgabe10()
	case 11:
		aufgabe11()
	default:
		fmt.Println("Diese Aufgabe gibt es noch nicht.")
	}
}

// ============================================================
// AUFGABE 1 – CODE-PUZZLE: BRING DEN CODE IN DIE RICHTIGE REIHENFOLGE
// ============================================================
//
// Unten stehen mehrere Code-Schnipsel.
// Sie ergeben zusammen ein kleines Programm.
//
// Eure Aufgabe:
//
// 1. Bringt die Schnipsel in die richtige Reihenfolge.
// 2. Schreibt/kopiert sie unten in die Funktion aufgabe01().
// 3. Speichert die Datei.
// 4. Startet das Programm mit:
//
//      go run .
//
// 5. Prüft, ob folgende Ausgabe entsteht:
//
//      Hallo Max!
//      Du hast 3 Äpfel.
//      Nach dem Einkauf hast du 5 Äpfel.
//
// ------------------------------------------------------------
// CODE-SCHNIPSEL – NOCH NICHT IN DER RICHTIGEN REIHENFOLGE
// ------------------------------------------------------------
//
//     fmt.Println("Du hast", apples, "Äpfel.")
//
//     apples = apples + 2
//
//     name := "Max"
//
//     fmt.Println("Nach dem Einkauf hast du", apples, "Äpfel.")
//
//     apples := 3
//
//     fmt.Println("Hallo", name+"!")
//
// ------------------------------------------------------------
//
// Tipp:
// Variablen müssen zuerst erstellt werden,
// bevor man sie benutzen kann.
//
// BONUS 1:
// Ändert den Namen und die Anzahl der Äpfel.
//
// BONUS 2:
// Kauft statt 2 Äpfeln 5 weitere.
//
// BONUS 3:
// Fügt am Ende hinzu, dass ein Apfel gegessen wird.
// Danach soll die neue Anzahl ausgegeben werden.

func aufgabe01() {
	// TODO:
	// Kopiert die Code-Schnipsel von oben hier hinein
	// und bringt sie in die richtige Reihenfolge.

	fmt.Println("Aufgabe 1: Sortiert die Code-Schnipsel aus den Kommentaren!")
}

// ============================================================
// AUFGABE 2 – WAS KOMMT RAUS?
// ============================================================
//
// Erst raten, DANN ausführen.
//
// Fragen:
// 1. Was wird ausgegeben?
// 2. Was passiert, wenn x auf 10 geändert wird?
// 3. Was ist der Unterschied zwischen
//      fmt.Println(x + y)
//    und
//      fmt.Println("x + y")
//
// BONUS:
// Fügt eine Ausgabe für x*y und x-y hinzu.

func aufgabe02() {
	x := 5
	y := 2

	fmt.Println(x + y)
	fmt.Println("x + y")
}

// ============================================================
// AUFGABE 3 – ÄNDERE EINE SACHE
// ============================================================
//
// Ändert:
// - euren Namen
// - euer Alter
//
// Ergänzt:
// - eine Ausgabe für euer Alter im nächsten Jahr
// - eine Ausgabe für euer Alter in 10 Jahren
//
// BONUS:
// Erstellt eine neue Variable "stadt"
// und gebt auch euren Wohn- oder Studienort aus.

func aufgabe03() {
	name := "Max"
	alter := 18

	fmt.Println("Hallo", name)
	fmt.Println("Du bist", alter, "Jahre alt.")

	// TODO: Alter im nächsten Jahr ausgeben
	// TODO: Alter in 10 Jahren ausgeben
}

// ============================================================
// AUFGABE 4 – CODE-DETEKTIV
// ============================================================
//
// Schaut euch den Code an, ohne ihn zuerst auszuführen.
//
// Fragen:
// - Welchen Wert hat apples am Anfang?
// - Was passiert in der zweiten Zeile?
// - Was wird ausgegeben?
// - Was könnte := bedeuten?
// - Was könnte = bedeuten?
//
// Experiment:
// Ändert +2 zu:
// - +10
// - -1
// - *2
//
// BONUS:
// Legt zusätzlich eine Variable "bananas" an.
// Gebt am Ende die Gesamtzahl aller Früchte aus.

func aufgabe04() {
	apples := 5
	apples = apples + 2

	fmt.Println("Äpfel:", apples)
}

// ============================================================
// AUFGABE 5 – BUG HUNT
// ============================================================
//
// Unten stehen mehrere kaputte Codezeilen als Kommentare.
//
// Nehmt IMMER NUR EINE davon,
// kopiert sie in den aktiven Bereich und versucht:
//
// 1. Programm starten
// 2. Fehlermeldung lesen
// 3. Fehler finden
// 4. Fehler beheben
//
// Fehler 1:
// fmt.Println("Hallo Welt!)
//
// Fehler 2:
// fmt.Prinln("Hallo")
//
// Fehler 3:
// fmt.Println(unbekannt)
//
// Fehler 4:
// x := 5
// x := 10
// fmt.Println(x)
//
// Fehler 5:
// fmt.Println("Hallo")
// fmt.Println("Welt"
//
// BONUS:
// Baut selbst einen kleinen Fehler ein
// und lasst euren Sitznachbarn herausfinden, was kaputt ist.

func aufgabe05() {
	fmt.Println("Bug Hunt!")

	// Kopiert hier jeweils EINEN kaputten Schnipsel hinein.
}

// ============================================================
// AUFGABE 6 – MENSCHLICHER COMPUTER
// ============================================================
//
// Führt das Programm ZEILE FÜR ZEILE im Kopf aus.
//
// Schreibt euch nach jeder Zeile den Wert von x auf:
//
//     x := 3
//     x = x + 2
//     x = x * 4
//     x = x - 5
//
// Was wird am Ende ausgegeben?
//
// Erst danach starten.
//
// BONUS:
// Verändert die Rechenoperationen so,
// dass am Ende genau 42 herauskommt.

func aufgabe06() {
	x := 3
	x = x + 2
	x = x * 4
	x = x - 5

	fmt.Println("x =", x)
}

// ============================================================
// AUFGABE 7 – ENTSCHEIDUNGEN ENTDECKEN
// ============================================================
//
// Ändert "alter" mehrfach und beobachtet die Ausgabe.
//
// Probiert:
// - 10
// - 17
// - 18
// - 25
//
// Fragen:
// - Wann wird "volljährig" ausgegeben?
// - Was bedeutet >= vermutlich?
//
// TODO:
// Ändert den Text in eigene Formulierungen.
//
// BONUS:
// Ergänzt einen weiteren Fall:
// Unter 16 soll "Noch keine 16" ausgegeben werden.
//
// Hinweis:
// Ihr könnt dafür nach "Go else if" suchen oder experimentieren.

func aufgabe07() {
	alter := 18

	if alter >= 18 {
		fmt.Println("Du bist volljährig.")
	} else {
		fmt.Println("Du bist noch nicht volljährig.")
	}
}

// ============================================================
// AUFGABE 8 – ZAHLEN-ORAKEL
// ============================================================
//
// Verändert number und findet heraus,
// wann welcher Text erscheint.
//
// Testet:
// - 2
// - 5
// - 6
// - 100
//
// TODO:
// Sorgt dafür, dass die Zahl 5 eine eigene Ausgabe bekommt.
//
// Gewünschtes Verhalten:
// kleiner als 5  -> "klein"
// genau 5        -> "genau fünf"
// größer als 5   -> "groß"
//
// BONUS:
// Könnt ihr auch prüfen, ob eine Zahl negativ ist?

func aufgabe08() {
	number := 7

	if number > 5 {
		fmt.Println("Die Zahl ist groß!")
	} else {
		fmt.Println("Die Zahl ist klein oder gleich 5!")
	}

	// TODO: Genau 5 getrennt behandeln
}

// ============================================================
// AUFGABE 9 – SCHLEIFEN ENTDECKEN
// ============================================================
//
// Erst raten, dann ausführen.
//
// Fragen:
// - Wie oft wird etwas ausgegeben?
// - Welche Zahlen erscheinen?
//
// Probiert danach:
// - i < 10
// - i < 3
// - i += 2 statt i++
// - Start bei i := 1
//
// BONUS:
// Lasst nur die Zahlen
// 10, 20, 30, 40, 50
// ausgeben.

func aufgabe09() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

// ============================================================
// AUFGABE 10 – CODE-LEGO
// ============================================================
//
// Baut aus diesen Ideen euer eigenes kleines Programm:
//
//     name := "Max"
//     alter := 18
//     fmt.Println("Hallo")
//     fmt.Println(name)
//     fmt.Println(alter)
//     fmt.Println(alter + 10)
//
// Ziel:
// Das Programm soll ungefähr Folgendes ausgeben:
//
//     Hallo!
//     Ich heiße Max.
//     Ich bin 18 Jahre alt.
//     In 10 Jahren bin ich 28.
//
// Ihr dürft die Reihenfolge und Texte selbst wählen.
//
// BONUS:
// Fügt eine if-Abfrage ein.
// Zum Beispiel:
// - volljährig / nicht volljährig
// - Alter größer als 20
// - Alter kleiner als 18

func aufgabe10() {
	// TODO: Baut hier euer Programm.

	fmt.Println("Code-Lego: Baut euer eigenes Programm!")
}

// ============================================================
// AUFGABE 11 – FREIE MINI-CHALLENGE
// ============================================================
//
// Baut ein kleines Programm über euch.
//
// Mindestanforderungen:
//
// - mindestens 3 Variablen
// - mindestens 4 Ausgaben mit fmt.Println
// - mindestens eine Rechnung
// - mindestens eine if-Abfrage
//
// Beispielausgabe:
//
//     Hallo!
//     Ich heiße Lisa.
//     Ich bin 19 Jahre alt.
//     In 10 Jahren bin ich 29.
//     Ich bin volljährig.
//
// Ihr könnt z.B. verwenden:
//
//     name := "Lisa"
//     alter := 19
//     lieblingszahl := 7
//
// BONUS 1:
// Baut eine Schleife ein.
//
// BONUS 2:
// Lasst etwas fünfmal ausgeben.
//
// BONUS 3:
// Erfindet ein Mini-Spiel:
// - Punktestand
// - Zahlen-Orakel
// - Altersprüfung
// - Countdown
// - kleine Rechenmaschine
//
// EXTRA:
// Wenn ihr schon Programmiererfahrung habt,
// versucht eine eigene Funktion zu schreiben und aufzurufen.

func aufgabe11() {
	fmt.Println("Freie Mini-Challenge!")

	// TODO: Euer eigenes Programm beginnt hier.
}
