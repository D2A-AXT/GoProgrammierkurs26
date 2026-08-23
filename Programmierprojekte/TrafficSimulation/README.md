# Projekt: Verkehrssimulation

## Ziel

In diesem Projekt programmierst du die Logik einer kleinen Verkehrssimulation.

Auf dem Bildschirm befindet sich eine Kreuzung mit:

- zwei Verkehrsrichtungen
- zwei Ampeln
- automatisch erzeugten Fahrzeugen
- Stopplinien
- Sicherheitsabständen
- wechselnden Ampelphasen

Die komplette Grafik und der Simulations-Takt sind bereits vorbereitet.

Du implementierst die **Verkehrslogik**.

---

# Das übst du

Besonders wichtig sind:

- Structs
- Pointer
- Slices
- Funktionen
- Bedingungen
- Schleifen
- Enums mit `iota`
- Zustände / Zustandswechsel
- einfache Simulationen

---

# Starten

Öffne diesen Ordner in VS Code.

Programm starten:

```bash
go run .
```

Tests starten:

```bash
go test ./...
```

---

# Bedienung

```text
SPACE    Simulation pausieren / fortsetzen
N        Ampelphase manuell weiterschalten
R        Simulation zurücksetzen
```

Die Fahrzeuge werden automatisch erzeugt.

---

# Wo arbeite ich?

Alle Pflichtaufgaben befinden sich in:

```text
traffic.go
```

Suche dort nach:

```go
// TODO
```

Die Dateien:

```text
game.go
ui.go
main.go
```

enthalten das fertige Framework.

---

# Datenmodell

## Fahrtrichtung

```go
type Direction int

const (
    East Direction = iota
    South
)
```

## Ampelfarbe

```go
type LightState int

const (
    Red LightState = iota
    Yellow
    Green
)
```

## Ampelphase

```go
type TrafficPhase int

const (
    EastGreen TrafficPhase = iota
    EastYellow
    SouthGreen
    SouthYellow
)
```

Die gesamte Kreuzung befindet sich also immer in genau **einem Zustand**.

---

# Aufgabe 1 – Nächste Ampelphase

Implementiere:

```go
func NextPhase(
    phase TrafficPhase,
) TrafficPhase
```

Die Reihenfolge lautet:

```text
EastGreen
    ↓
EastYellow
    ↓
SouthGreen
    ↓
SouthYellow
    ↓
EastGreen
```

Das ist ein kleiner **Zustandsautomat**.

---

# Aufgabe 2 – Welche Farbe sieht ein Fahrzeug?

Implementiere:

```go
func LightForDirection(
    phase TrafficPhase,
    direction Direction,
) LightState
```

Beispiel:

```text
Phase: EastGreen

East  -> Green
South -> Red
```

In:

```text
SouthYellow
```

gilt:

```text
East  -> Red
South -> Yellow
```

---

# Aufgabe 3 – Abstand zur Stopplinie

Implementiere:

```go
func DistanceToStopLine(
    car Car,
    stopLine float64,
) float64
```

Ein Auto in Richtung `East` fährt entlang der X-Achse:

```text
distance = stopLine - car.X
```

Ein Auto in Richtung `South` fährt entlang der Y-Achse:

```text
distance = stopLine - car.Y
```

Ist das Fahrzeug bereits hinter der Linie, darf das Ergebnis negativ sein.

---

# Aufgabe 4 – Sicherheitsabstand

Implementiere:

```go
func HasSafeDistance(
    car Car,
    other Car,
    minGap float64,
) bool
```

Nur Fahrzeuge in derselben Richtung interessieren uns.

Beispiel:

```text
car        other
 ███       ███
 100       150
```

`other` fährt voraus.

Berücksichtige außerdem die Länge der Fahrzeuge.

Für zwei gleich lange Fahrzeuge kannst du dir vorstellen:

```text
Abstand zwischen Mittelpunkten
    -
halbe Länge car
    -
halbe Länge other
```

Dieser freie Abstand muss mindestens `minGap` sein.

---

# Aufgabe 5 – Darf das Fahrzeug fahren?

Implementiere:

```go
func CanCarMove(
    car Car,
    cars []Car,
    light LightState,
    stopLine float64,
    minGap float64,
    movementPerStep float64,
) bool
```

Hier kommen mehrere vorherige Funktionen zusammen.

## Regel 1 – Fahrzeug davor

Ist ein anderes Fahrzeug zu nah:

```text
STOP
```

## Regel 2 – Kreuzung bereits betreten

Ist das Fahrzeug bereits hinter der Stopplinie:

```text
WEITERFAHREN
```

Es soll nicht mitten in der Kreuzung stehen bleiben.

## Regel 3 – Grün

```text
FAHREN
```

## Regel 4 – Rot

```text
STOP
```

## Regel 5 – Gelb

Ist das Fahrzeug noch weit genug entfernt:

```text
STOP
```

Würde es aber im nächsten Simulationsschritt die Stopplinie erreichen oder
überfahren:

```text
WEITERFAHREN
```

---

# Aufgabe 6 – Fahrzeug bewegen

Implementiere:

```go
func MoveCar(
    car *Car,
    distance float64,
)
```

Hier wird ein Pointer verwendet.

Für `East`:

```go
car.X += distance
```

Für `South`:

```go
car.Y += distance
```

---

# Aufgabe 7 – Ausgefahrene Autos entfernen

Implementiere:

```go
func RemoveExitedCars(
    cars []Car,
    maxX float64,
    maxY float64,
) []Car
```

Fahrzeuge, die das sichtbare Feld verlassen haben, sollen aus dem Slice
entfernt werden.

Du kannst dafür einen neuen Slice aufbauen:

```go
remaining := []Car{}
```

und passende Fahrzeuge mit:

```go
remaining = append(
    remaining,
    car,
)
```

hinzufügen.

---

# Aufgabe 8 – Wartende Autos zählen

Implementiere:

```go
func CountWaitingCars(
    cars []Car,
    phase TrafficPhase,
    eastStopLine float64,
    southStopLine float64,
) int
```

Ein Fahrzeug zählt als wartend, wenn:

1. es noch vor seiner Stopplinie steht
2. seine Ampel gerade Rot ist

Der Wert wird unten im Fenster angezeigt.

---

# Empfohlene Reihenfolge

```text
NextPhase
   ↓
LightForDirection
   ↓
DistanceToStopLine
   ↓
HasSafeDistance
   ↓
CanCarMove
   ↓
MoveCar
   ↓
RemoveExitedCars
   ↓
CountWaitingCars
```

Nach jeder Aufgabe:

```bash
go test ./...
```

---

# Einzelne Tests

Zum Beispiel:

```bash
go test -run TestNextPhase
```

oder:

```bash
go test -run TestCanCarMove
```

---

# Bonusaufgaben

In:

```text
bonus.go
```

findest du zusätzliche Aufgaben:

- Durchschnittsgeschwindigkeit berechnen
- Fahrzeuge nach Fahrtrichtung zählen
- längste Warteschlange bestimmen

---

# Was passiert im Hintergrund?

Die Oberfläche arbeitet ungefähr so:

```text
Timer
  ↓
Ampelphase aktualisieren
  ↓
Fahrzeuge erzeugen
  ↓
CanCarMove(...)
  ↓
MoveCar(...)
  ↓
RemoveExitedCars(...)
  ↓
Grafik zeichnen
```

Die Datei `traffic.go` kennt Ebitengine überhaupt nicht.

Damit bleibt die Verkehrslogik unabhängig von der Grafik und lässt sich mit
normalen Unit Tests prüfen.

Das ist auch bei realen Anwendungen ein wichtiges Architekturprinzip.
