# Basics 2 – Schleifen und Slices

**Themen:** for-Schleifen, range, Index und Wert, append und Slicing.

## So funktionieren die Tests

Zu jeder Aufgabe gehören zwei Dateien:

- die eigentliche `.go`-Datei – **hier arbeitest du**
- eine passende `_test.go`-Datei – **hier steht der automatische Test**

In VS Code mit installierter **Go-Erweiterung** erscheint über der Testfunktion ein
**Run Test**-Button. Ein Klick darauf führt den Test aus:

- **PASS / grün:** Deine Funktion liefert das erwartete Ergebnis.
- **FAIL / rot:** Das tatsächliche Ergebnis unterscheidet sich vom erwarteten Ergebnis.

Die Tests sind absichtlich lesbar gehalten. Typischerweise findest du darin Werte wie:

```go
input := 5
want := 8
```

`input` ist die Eingabe für deine Funktion. `want` ist das erwartete Ergebnis. Du darfst
diese Werte verändern und danach erneut testen.

Alternativ kannst du im Ordner einer Aufgabe im Terminal ausführen:

```bash
go test
```

Wichtig: Ändere beim normalen Bearbeiten **nur die Aufgabendatei**. Die Testdatei solltest
du zunächst unverändert lassen, damit du zuverlässig prüfen kannst, ob deine Lösung stimmt.

## Aufgaben

### 01. Bis fünf zählen

**Schwierigkeit:** 1/5 – sehr leicht

Implementiere `CountToFive() []int`. Erzeuge mit einer klassischen `for`-Schleife
das Slice `[1, 2, 3, 4, 5]`. Die Aufgabe führt nur Schleifenvariable, Bedingung und Inkrement ein.

**Dateien:** `bis_fuenf_zaehlen.go` und `bis_fuenf_zaehlen_test.go`

---

### 02. Countdown

**Schwierigkeit:** 2/5 – leicht

Implementiere `Countdown(start int) []int`. Die Funktion soll vom übergebenen
Startwert rückwärts bis 1 zählen. Gegenüber Aufgabe 1 ändern sich Startwert, Bedingung und die
Richtung des Zählers.

**Dateien:** `countdown.go` und `countdown_test.go`

---

### 03. Summe bilden

**Schwierigkeit:** 2/5 – leicht

Implementiere `SumTo(n int) int`. Addiere mit einer Schleife alle Zahlen von 1 bis `n`.
Neu ist hier eine zusätzliche Akkumulator-Variable, deren Wert sich in jedem Schleifendurchlauf verändert.

**Dateien:** `summe_bilden.go` und `summe_bilden_test.go`

---

### 04. Werte mit range durchlaufen

**Schwierigkeit:** 3/5 – mittel

Implementiere `CopyValues(numbers []int) []int`. Laufe mit `range` über ein vorhandenes
Slice und übertrage jeden Wert in ein neues Slice. Der Index wird in dieser Aufgabe bewusst nicht benötigt.

**Dateien:** `werte_mit_range.go` und `werte_mit_range_test.go`

---

### 05. Index und Wert verwenden

**Schwierigkeit:** 3/5 – mittel

Implementiere `LabelValues(values []string) []string`. Verwende `range`, diesmal aber
mit Index **und** Wert. Aus `["Apfel", "Banane"]` sollen Texte wie `"0: Apfel"` und `"1: Banane"`
entstehen. Zusätzlich wird `fmt.Sprintf` zum Formatieren verwendet.

**Dateien:** `index_und_wert.go` und `index_und_wert_test.go`

---

### 06. Slice erweitern

**Schwierigkeit:** 3/5 – mittel

Implementiere `ExtendNames(names []string, first string, second string) []string`.
Füge zwei neue Werte am Ende eines bestehenden Slices hinzu. Diese Aufgabe konzentriert sich auf
die Verwendung von `append` ohne zusätzliche Schleifenlogik.

**Dateien:** `slice_erweitern.go` und `slice_erweitern_test.go`

---

### 07. Slice ausschneiden

**Schwierigkeit:** 4/5 – anspruchsvoller

Implementiere `MiddleSection(numbers []int) []int`. Gib nur die Elemente mit den
Indizes 1 bis 3 zurück. Dafür wird die Slicing-Syntax benötigt und besonders wichtig ist,
dass der Endindex bei Go **nicht** enthalten ist.

**Dateien:** `slice_ausschneiden.go` und `slice_ausschneiden_test.go`

---

### 08. Gerade Zahlen filtern

**Schwierigkeit:** 5/5 – Abschlussaufgabe

Implementiere `FilterEven(numbers []int) []int`. Laufe mit `range` über ein Slice,
prüfe jeden Wert mit Modulo und füge nur gerade Zahlen an ein neues Slice an. Diese Abschlussaufgabe
kombiniert Schleife, `range`, `if`, Vergleich, Modulo und `append`.

**Dateien:** `gerade_zahlen_filtern.go` und `gerade_zahlen_filtern_test.go`

---

## Lösungshilfen

Nutze diesen Abschnitt erst, wenn du nach einem eigenen Lösungsversuch festhängst.

### 01. Bis fünf zählen

- Starte mit `i := 1`.
- Die Schleife läuft solange `i <= 5`.
- Füge `i` mit `append` an das Ergebnis an.

### 02. Countdown

- Beginne die Schleife mit `i := start`.
- Verwende `i--` statt `i++`.
- Die 1 soll noch enthalten sein.

### 03. Summe bilden

- Lege vor der Schleife `sum := 0` an.
- In jedem Durchlauf: `sum = sum + i`.

### 04. Werte mit range durchlaufen

- Das Muster lautet `for _, value := range numbers`.
- Der Unterstrich `_` ignoriert den Index.
- Hänge `value` mit `append` an.

### 05. Index und Wert verwenden

- `for index, value := range values` liefert beides.
- Mit `fmt.Sprintf("%d: %s", index, value)` kannst du den Text erzeugen.

### 06. Slice erweitern

- `append` liefert das erweiterte Slice zurück.
- Du kannst mehrere Werte in einem Aufruf anhängen: `append(names, first, second)`.

### 07. Slice ausschneiden

- Gesucht sind die Indizes 1, 2 und 3.
- Dafür lautet der passende Ausschnitt `numbers[1:4]`.

### 08. Gerade Zahlen filtern

- Gerade Zahlen erfüllen `value % 2 == 0`.
- Lege zunächst ein leeres Ergebnis-Slice an.
- Nur im `if`-Zweig wird `append` ausgeführt.
