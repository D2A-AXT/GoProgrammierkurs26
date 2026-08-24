package main

import (
	"fmt"
	"math/rand"
)

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
	exercise := 11

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

	name := "Max"
	fmt.Println("Hallo", name+"!")
	apples := 3
	fmt.Println("Du hast", apples, "Äpfel.")
	apples = apples + 2
	fmt.Println("Nach dem Einkauf hast du", apples, "Äpfel.")
	apples--
	fmt.Println("Du hattest Hunger. Es sind", apples, "Äpfel übrig.")

	// fmt.Println("Aufgabe 1: Sortiert die Code-Schnipsel aus den Kommentaren!")
}

// ============================================================
// AUFGABE 2 – WAS KOMMT RAUS?
// ============================================================
//
// Erst raten, DANN ausführen.
//
// Fragen:
// 1. Was wird ausgegeben?
// 2. Was passiert, wenn x auf 10 geändert wird? => Ausgabe 12
// 3. Was ist der Unterschied zwischen
//      fmt.Println(x + y) => Variablen werden verrechnet
//    und
//      fmt.Println("x + y") = Nur der String wird ausgegeben
//
// BONUS:
// Fügt eine Ausgabe für x*y und x-y hinzu.

func aufgabe02() {
	x := 5
	y := 2

	fmt.Println("x + y:") //Ausgabe: "x + y"
	fmt.Println(x + y)    // Ausgabe: "7"
	fmt.Println("x - y:") //Ausgabe: "x + y"
	fmt.Println(x - y)    // Ausgabe: "7"
	fmt.Println("x * y:") //Ausgabe: "x + y"
	fmt.Println(x * y)    // Ausgabe: "7"
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
	name := "Rieko"
	alter := 19
	var stadt string = "Heidelberg"

	fmt.Println("Hallo", name)
	fmt.Println("Du bist", alter, "Jahre alt.")
	fmt.Println("Nächstes Jahr bist du", alter+1, "Jahre alt.")
	fmt.Println("In 10 Jahren bist du", alter+10, "Jahre alt.")
	fmt.Println("Du befindest dich in", stadt)

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
// - Welchen Wert hat apples am Anfang? => 5
// - Was passiert in der zweiten Zeile? => zwei werden zu apples addiert
// - Was wird ausgegeben? => 7
// - Was könnte := bedeuten? => Shortcut mit initialisierung und ohne expliziten Typ
// - Was könnte = bedeuten? => Assignment
//
// Experiment:
// Ändert +2 zu:
// - +10 => 15
// - -1 => 4
// - *2 => 10
//
// BONUS:
// Legt zusätzlich eine Variable "bananas" an.
// Gebt am Ende die Gesamtzahl aller Früchte aus.

func aufgabe04() {
	apples := 5
	apples = apples + 2
	var bananas int = 4

	fmt.Println("Äpfel:", apples)
	fmt.Println("Bananen:", bananas)
	fmt.Println("Freche Früchtchen:", apples+bananas)

}

// ============================================================s
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
// fmt.Println("Hallo Welt!)=>
// fmt.Println("Hallo Welt!")
//
// Fehler 2:
// fmt.Prinln("Hallo") =>
// fmt.Println("Hallo")
//
// Fehler 3:
// fmt.Println(unbekannt) => Variable belegen
//
// Fehler 4:
// x := 5
// x := 10 => Doppelpunkt weg
// fmt.Println(x)
//
// Fehler 5:
// fmt.Println("Hallo")
// fmt.Println("Welt" => Klammer dazu
//
// BONUS:
// Baut selbst einen kleinen Fehler ein
// und lasst euren Sitznachbarn herausfinden, was kaputt ist.

func aufgabe05() {

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
//     x := 3      => 3
//     x = x + 2   => 5
//     x = x * 4   => 20
//     x = x - 5   => 15
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
	x = x * 9
	x = x - 3

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
// - Wann wird "volljährig" ausgegeben? Wenn größer oder gleich als 18
// - Was bedeutet >= vermutlich? größer oder gleich
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
	alter := 15

	if alter >= 18 {
		fmt.Println("Du bist volljährig.")
	} else if alter >= 16 {
		fmt.Println("Du bist noch minderjährig.")
	} else {
		fmt.Println("Du bist noch sehr minderjährig.")
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
	} else if number == 5 {
		fmt.Println("Die Zahl ist genau fünf")
	} else {
		fmt.Println("Die Zahl ist klein!")
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
// - Wie oft wird etwas ausgegeben? => 5 mal
// - Welche Zahlen erscheinen? => 0,1,2,3,4
//
// Probiert danach:
// - i < 10 => 0 bis 9
// - i < 3  => 0 bis 2
// - i += 2 statt i++ => jede zweite zahl
// - Start bei i := 1 => 1 bis 5
//
// BONUS:
// Lasst nur die Zahlen
// 10, 20, 30, 40, 50
// ausgeben.

func aufgabe09() {
	for i := 10; i <= 50; i = i + 10 {
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

	name := "Rieko"
	alter := 19
	fmt.Println("Hallo!")
	fmt.Printf("Ich heiße %s.\n", name)
	fmt.Printf("Ich bin %d Jahre alt.\n", alter)
	fmt.Printf("In 10 Jahren bin ich %d.\n", alter+10)

	//fmt.Println("Code-Lego: Baut euer eigenes Programm!")
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

var charmap map[int8]string = map[int8]string{
	0: "~", // water
	1: "~", // ship
	2: "x", // miss
	3: "#", // hit
}

var ships [10]int8 = [10]int8{5, 4, 4, 3, 3, 3, 2, 2, 2, 2}

const fieldsize int8 = 10

func aufgabe11() {
	fmt.Println("Freie Mini-Challenge!")

	var input_x int8 = 0
	var input_y int8 = 0

	var spielfeld [fieldsize][fieldsize]int8

	populate(&spielfeld)

	for isShipRemaining(&spielfeld) {

		printfield(&spielfeld)

		fmt.Print("Please input your next guess as two numbers (x and y) as two numbers seperated by a space:\n>>> ")
		fmt.Scan(&input_x, &input_y)

		fmt.Println("X:", input_x, "Y:", input_y)
		fmt.Print("\n\n\n")

		hit(&spielfeld, input_x, input_y)

	}

	fmt.Println("You Won.\nCongratulations!")
	printfield(&spielfeld)

}

func populate(field *[fieldsize][fieldsize]int8) {
	// Reset
	for x := range fieldsize {
		for y := range fieldsize {
			field[x][y] = 0
		}
	}

	placeShips(field, 0)

}

func printfield(field *[fieldsize][fieldsize]int8) {
	fmt.Print(" #")
	for x := range fieldsize {
		fmt.Print(" ")
		fmt.Print(x)
	}
	fmt.Print("\n")

	for y := range fieldsize {
		fmt.Printf(" %d", y)
		for x := range fieldsize {
			fmt.Printf(" %s", charmap[field[x][y]])
		}
		fmt.Print("\n")

	}
}

func hit(field *[fieldsize][fieldsize]int8, x int8, y int8) {
	switch field[x][y] {
	case 0:
		field[x][y] = 2
	case 1:
		field[x][y] = 3
	}
}

func isShipRemaining(field *[fieldsize][fieldsize]int8) bool {
	var shipRemaining bool = false
	for y := range fieldsize {
		for x := range fieldsize {
			if field[x][y] == 1 {
				shipRemaining = true
			}
		}
	}
	return shipRemaining
}

type ShipPos struct {
	x   int8
	y   int8
	hor bool
}

func placeShips(field *[fieldsize][fieldsize]int8, index int) bool {

	if index >= len(ships) {
		return true
	}

	var size int8 = ships[index]
	var x int8
	var y int8

	free_positions := make([]ShipPos, 0)
	for y = range fieldsize {
		for x = range fieldsize {
			// fmt.Println("checking", x, y)
			if doesShipFit(field, ShipPos{x, y, false}, size) {
				free_positions = append(free_positions, ShipPos{x, y, false})
				// fmt.Println("Horizontal fits")

			}
			if doesShipFit(field, ShipPos{x, y, true}, size) {
				free_positions = append(free_positions, ShipPos{x, y, true})
				// fmt.Println("Vertical fits")

			}
		}
	}

	// fmt.Println(free_positions)
	// fmt.Println(len(free_positions))

	var pos_index int
	var my_pos ShipPos
	var okay bool

	for len(free_positions) != 0 {

		pos_index = rand.Intn(len(free_positions))

		my_pos = free_positions[pos_index]

		free_positions = append(free_positions[:pos_index], free_positions[pos_index+1:]...)

		drawShip(field, my_pos, size, true)

		// printfield(field)
		// fmt.Println()
		// fmt.Scanln()

		okay = placeShips(field, index+1)

		if okay {
			return true
		}

		drawShip(field, my_pos, size, false)

	}

	return false
}

func doesShipFit(field *[fieldsize][fieldsize]int8, pos ShipPos, size int8) bool {
	var x int8
	var y int8
	var lower_x int8
	var lower_y int8
	var upper_x int8
	var upper_y int8

	if pos.hor {
		// Check size
		if pos.x+size > fieldsize {
			return false
		}

		lower_x = max(pos.x-1, 0)
		lower_y = max(pos.y-1, 0)
		upper_x = min(pos.x+size, fieldsize-1) //not +1 cause size is one more than must be added to find end of ship
		upper_y = min(pos.y+1, fieldsize-1)

	} else {

		// Check size
		if pos.y+size > fieldsize {
			return false
		}

		lower_x = max(pos.x-1, 0)
		lower_y = max(pos.y-1, 0)
		upper_x = min(pos.x+1, fieldsize-1)
		upper_y = min(pos.y+size, fieldsize-1) //not +1 cause size is one more than must be added to find end of ship
	}
	// fmt.Println(lower_x, lower_y, upper_x, upper_y)

	//check if space is free
	for x = lower_x; x <= upper_x; x++ {
		for y = lower_y; y <= upper_y; y++ {
			if field[x][y] != 0 {
				return false
			}
		}
	}
	return true
}

func drawShip(field *[fieldsize][fieldsize]int8, pos ShipPos, size int8, exist bool) {

	var value int8

	if exist {
		value = 1
	} else {
		value = 0
	}

	// fmt.Println("Drawing Ship", pos, "with len", size)

	if pos.hor {
		for x := pos.x; x < pos.x+size; x++ {
			field[x][pos.y] = value

			// fmt.Println("Drawing at x", x)

		}
	} else {
		for y := pos.y; y < pos.y+size; y++ {
			// fmt.Println("Drawing at y", y)
			field[pos.x][y] = value
		}
	}
}
