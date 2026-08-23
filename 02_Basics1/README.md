# Basics 1 – Operationen, Bedingungen und Funktionen

**Themen:** Arithmetik, Vergleich, logische Operatoren, Funktionen sowie if / else.

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

### 01. Zwei Zahlen addieren

**Schwierigkeit:** 1/5 – sehr leicht

Implementiere `Add(a int, b int) int`. Die Funktion bekommt zwei ganze Zahlen und soll
deren Summe zurückgeben. Nach Basics 0 ist das die bewusst einfachste Aufgabe: ein Operator,
eine Berechnung und ein Rückgabewert.

**Dateien:** `zwei_zahlen_addieren.go` und `zwei_zahlen_addieren_test.go`

---

### 02. Grundrechenarten

**Schwierigkeit:** 2/5 – leicht

Erweitere die Addition um mehrere arithmetische Operatoren. `BasicOperations(a, b)`
soll Summe, Differenz und Produkt zurückgeben. Dadurch übst du mehrere Berechnungen, aber noch
ohne Bedingungen oder zusätzliche Logik.

**Dateien:** `grundrechenarten.go` und `grundrechenarten_test.go`

---

### 03. Nachricht ausgeben

**Schwierigkeit:** 2/5 – leicht

Schreibe eine Funktion ohne Rückgabewert. `PrintMessage(message string)` soll den
übergebenen Text mit `fmt.Println` ausgeben. Hier steht nicht die Berechnung im Mittelpunkt,
sondern der Unterschied zwischen einer Funktion mit und ohne Rückgabewert.

**Dateien:** `nachricht_ausgeben.go` und `nachricht_ausgeben_test.go`

---

### 04. Zahlen vergleichen

**Schwierigkeit:** 2/5 – leicht

Implementiere `IsGreater(a int, b int) bool`. Die Funktion soll `true` zurückgeben,
wenn `a` größer als `b` ist, ansonsten `false`. Hier kommt erstmals ein Vergleichsoperator ins Spiel.

**Dateien:** `zahlen_vergleichen.go` und `zahlen_vergleichen_test.go`

---

### 05. Gerade oder ungerade

**Schwierigkeit:** 3/5 – mittel

Implementiere `IsEven(number int) bool`. Jetzt werden zwei Operatoren kombiniert:
Mit Modulo wird der Rest einer Division bestimmt und anschließend mit `0` verglichen.

**Dateien:** `gerade_oder_ungerade.go` und `gerade_oder_ungerade_test.go`

---

### 06. Zutritt prüfen

**Schwierigkeit:** 3/5 – mittel

Implementiere `CanEnter(age int, hasID bool) bool`. Zutritt ist nur erlaubt, wenn
die Person mindestens 18 Jahre alt **und** ein Ausweisdokument vorhanden ist. Die Aufgabe kombiniert
einen Vergleich mit einem logischen Operator.

**Dateien:** `zutritt_pruefen.go` und `zutritt_pruefen_test.go`

---

### 07. Temperatur bewerten

**Schwierigkeit:** 4/5 – anspruchsvoller

Implementiere `TemperatureLabel(temp int) string`. Bei Temperaturen ab 25 soll
`"warm"` zurückgegeben werden, darunter `"kühl"`. Hier wird erstmals bewusst eine
`if/else`-Entscheidung verlangt.

**Dateien:** `temperatur_bewerten.go` und `temperatur_bewerten_test.go`

---

### 08. Punkte einstufen

**Schwierigkeit:** 5/5 – Abschlussaufgabe

Implementiere `ClassifyPoints(points int) string`. Ab 90 Punkten gilt `"sehr gut"`,
ab 75 `"gut"`, ab 50 `"bestanden"`, darunter `"nicht bestanden"`. Diese Aufgabe ist bewusst
die schwierigste in Basics 1, weil mehrere Vergleichsgrenzen in der richtigen Reihenfolge
mit `if / else if / else` kombiniert werden müssen.

**Dateien:** `punkte_einstufen.go` und `punkte_einstufen_test.go`

---

## Lösungshilfen

Nutze diesen Abschnitt erst, wenn du nach einem eigenen Lösungsversuch festhängst.

### 01. Zwei Zahlen addieren

- Der Additionsoperator ist `+`.
- Das Ergebnis kann direkt hinter `return` stehen.

### 02. Grundrechenarten

- Verwende `+`, `-` und `*`.
- Mehrere Rückgabewerte werden durch Kommas getrennt.

### 03. Nachricht ausgeben

- Die Funktion benötigt keinen Rückgabetyp.
- Verwende `fmt.Println(message)`.

### 04. Zahlen vergleichen

- Der Operator für 'größer als' ist `>`.
- Ein Vergleich liefert bereits direkt einen booleschen Wert.

### 05. Gerade oder ungerade

- `number % 2` liefert den Rest bei Division durch 2.
- Eine gerade Zahl hat dabei den Rest `0`.

### 06. Zutritt prüfen

- `age >= 18` prüft das Alter.
- Für UND wird in Go `&&` verwendet.

### 07. Temperatur bewerten

- Die Bedingung lautet sinngemäß `temp >= 25`.
- Nutze `if` für den ersten und `else` für den zweiten Fall.

### 08. Punkte einstufen

- Prüfe die höchste Grenze zuerst.
- Sobald ein Zweig zutrifft, werden spätere Zweige nicht mehr ausgeführt.
- Teste besonders 90, 75, 50 und Werte knapp darunter.

## Wenn alle Stricke reißen und die Sonne nicht mehr scheint:

- Du findest einen Ordner mit Musterlösungen in dem alle aufgaben gelöst sind
- bitte beachten, dass man am besten lernt wenn man mehrfach auf die nase fliegt und wieder aufsteht!