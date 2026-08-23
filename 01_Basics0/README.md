# Basics 0 – Variablen, Datentypen und Packages

**Themen:** Variablen, Datentypen, Zuweisung, Type Inference, Imports und Standard-Packages.

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

### 01. Hello, World!

**Schwierigkeit:** 1/5 – Einstieg

Erstelle deine erste kleine Go-Funktion. `HelloWorld()` soll exakt den Text
`Hello, World!` zurückgeben. Es gibt keine Parameter und keine zusätzliche Logik.
Die Aufgabe dient nur dazu, den Aufbau einer Funktion, `return` und den Test-Workflow kennenzulernen.

**Dateien:** `hello_world.go` und `hello_world_test.go`

---

### 02. Persönliche Begrüßung

**Schwierigkeit:** 1/5 – sehr leicht

Erweitere die erste Aufgabe um einen Parameter. `Greeting(name string)` soll
einen Namen entgegennehmen und daraus `Hallo, <Name>!` erzeugen. Damit übst du zum ersten Mal
einen Funktionsparameter und das Zusammensetzen von Strings.

**Dateien:** `persoenliche_begruessung.go` und `persoenliche_begruessung_test.go`

---

### 03. Variablen kennenlernen

**Schwierigkeit:** 2/5 – leicht

Lege in `StudentProfile()` vier Variablen mit expliziter Typangabe an:
Name (`string`), Alter (`int`), Größe (`float64`) und Studentenstatus (`bool`).
Gib anschließend alle vier Werte zurück. Die Aufgabe konzentriert sich bewusst nur auf Deklaration und Datentypen.

**Dateien:** `variablen_kennenlernen.go` und `variablen_kennenlernen_test.go`

---

### 04. Werte verändern

**Schwierigkeit:** 2/5 – leicht

In `UpdateScore()` existieren bereits zwei Variablen. Ändere deren Werte mit
einer normalen Zuweisung (`=`). Aus `status = "gestartet"` soll `"fertig"` werden und
aus `points = 0` soll `50` werden. So wird der Unterschied zwischen Deklaration und späterer Zuweisung sichtbar.

**Dateien:** `werte_aendern.go` und `werte_aendern_test.go`

---

### 05. Go erkennt den Typ

**Schwierigkeit:** 2/5 – leicht

Erstelle in `CourseData()` mehrere Variablen ausschließlich mit der Kurzform `:=`.
Go soll die Datentypen also selbst aus den zugewiesenen Werten ableiten. Zurückgegeben werden
Kursname, Teilnehmerzahl und die Information, ob der Kurs aktiv ist.

**Dateien:** `typableitung.go` und `typableitung_test.go`

---

### 06. Text mit fmt formatieren

**Schwierigkeit:** 3/5 – mittel

Verwende erstmals ein Standard-Package. `FormatCourse(course string, participants int)`
soll mit `fmt.Sprintf` einen formatierten Text erzeugen. Damit übst du Import, Funktionsaufruf
aus einem Package und unterschiedliche Platzhalter für Text und ganze Zahlen.

**Dateien:** `text_formatieren.go` und `text_formatieren_test.go`

---

### 07. Text mit strings bearbeiten

**Schwierigkeit:** 3/5 – mittel

Implementiere `ChangeCase(text string)`. Die Funktion soll denselben Text einmal
komplett groß- und einmal komplett kleingeschrieben zurückgeben. Dafür wird das Package `strings`
verwendet. Die Aufgabe führt zwei typische Funktionen aus einem Standard-Package ein.

**Dateien:** `text_bearbeiten.go` und `text_bearbeiten_test.go`

---

### 08. Status aus mehreren Packages

**Schwierigkeit:** 4/5 – anspruchsvoller

Kombiniere jetzt mehrere bereits bekannte Bausteine. `BuildStatus(name string, number int)`
soll den Namen mit `strings.ToUpper` großschreiben, die Zahl mit `strconv.Itoa` in Text umwandeln
und daraus mit `fmt.Sprintf` den erwarteten Status erzeugen. Die Aufgabe prüft das Zusammenspiel mehrerer Imports.

**Dateien:** `status_aus_packages.go` und `status_aus_packages_test.go`

---

## Lösungshilfen

Nutze diesen Abschnitt erst, wenn du nach einem eigenen Lösungsversuch festhängst.

### 01. Hello, World!

- Die Funktion gibt einen `string` zurück.
- Hinter `return` kann direkt ein Text in Anführungszeichen stehen.

### 02. Persönliche Begrüßung

- Strings lassen sich mit `+` verketten.
- Verwende den Parameter `name` zwischen Begrüßung und Ausrufezeichen.

### 03. Variablen kennenlernen

- Beispiel: `var age int = 20`.
- Mehrere Werte können mit `return a, b, c, d` zurückgegeben werden.

### 04. Werte verändern

- Die Variablen existieren bereits – deshalb `=` statt `:=` verwenden.
- Der Datentyp einer bestehenden Variable bleibt erhalten.

### 05. Go erkennt den Typ

- Beispiel: `course := "Go Grundlagen"`.
- Text wird zu `string`, ganze Zahlen zu `int`, `true/false` zu `bool`.

### 06. Text mit fmt formatieren

- Importiere `fmt`.
- `%s` steht für einen String, `%d` für einen Integer.
- `fmt.Sprintf(...)` gibt den erzeugten String zurück.

### 07. Text mit strings bearbeiten

- `strings.ToUpper(text)` erzeugt Großbuchstaben.
- `strings.ToLower(text)` erzeugt Kleinbuchstaben.

### 08. Status aus mehreren Packages

- Du benötigst `fmt`, `strings` und `strconv`.
- Erzeuge zuerst zwei Zwischenergebnisse und formatiere danach den finalen String.

## Wenn alle Stricke reißen und die Sonne nicht mehr scheint:

- Du findest einen Ordner mit Musterlösungen in dem alle aufgaben gelöst sind
- bitte beachten, dass man am besten lernt wenn man mehrfach auf die nase fliegt und wieder aufsteht!