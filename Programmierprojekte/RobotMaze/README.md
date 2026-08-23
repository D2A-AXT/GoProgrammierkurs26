# Roboter im Labyrinth

## Ziel

In diesem Projekt steuerst du einen kleinen Roboter durch ein Labyrinth.

Die Grafik, Tastatureingabe und der automatische Ablauf sind bereits vorbereitet.
Deine Aufgabe besteht darin, die eigentliche **Logik des Roboters** zu implementieren.

Das Projekt vertieft insbesondere:

- Datentypen und Variablen
- Funktionen
- Bedingungen
- Structs
- Pointer
- einfache Berechnungen mit Koordinaten
- das Zerlegen eines Problems in kleine Funktionen

---

## Starten

Diesen Ordner in VS Code öffnen

Danach:

```bash
go run .
```

Tests starten:

```bash
go test ./...
```

---

## Bedienung

### Manueller Modus

- `W` / `↑` – nach oben
- `D` / `→` – nach rechts
- `S` / `↓` – nach unten
- `A` / `←` – nach links

### Automatik

- `SPACE` – Automatik ein-/ausschalten
- `R` – Neustart

---

# Deine Aufgaben

Alle Pflichtaufgaben befinden sich in:

```text
game.go
```

Suche dort nach:

```go
// TODO
```

---

## Aufgabe 1 – Richtungen in Koordinaten umwandeln

Implementiere:

```go
func DirectionVector(direction Direction) Position
```

Eine Bewegung nach rechts bedeutet zum Beispiel:

```text
X + 1
Y + 0
```

Nach oben:

```text
X + 0
Y - 1
```

Diese kleine Hilfsfunktion wird später von mehreren anderen Funktionen verwendet.

---

## Aufgabe 2 – Liegt eine Position im Spielfeld?

Implementiere:

```go
func IsInside(world World, position Position) bool
```

Eine Position darf weder links/oberhalb des Spielfelds noch rechts/unterhalb
der vorhandenen Zellen liegen.

Achte besonders auf negative Werte.

---

## Aufgabe 3 – Wände erkennen

Implementiere:

```go
func IsWall(world World, position Position) bool
```

Im Labyrinth werden Wände durch `'#'` dargestellt.

Positionen außerhalb des Spielfelds sollen ebenfalls als Wand gelten.

---

## Aufgabe 4 – Darf der Roboter fahren?

Implementiere:

```go
func CanMove(world World, robot Robot, direction Direction) bool
```

Berechne zuerst die Position, auf der der Roboter nach einem Schritt landen würde.

Prüfe anschließend:

> Ist dieses Feld frei?

---

## Aufgabe 5 – Roboter bewegen

Implementiere:

```go
func MoveRobot(world World, robot *Robot, direction Direction)
```

Hier wird ein **Pointer** verwendet:

```go
robot *Robot
```

Dadurch kannst du den ursprünglichen Robot verändern.

Regeln:

1. Der Roboter schaut anschließend immer in die gewünschte Richtung.
2. Ist das Feld frei, bewegt er sich um genau ein Feld.
3. Ist dort eine Wand, bleibt seine Position unverändert.

---

## Aufgabe 6 – Ziel erkennen

Implementiere:

```go
func HasReachedGoal(robot Robot, goal Position) bool
```

Der Roboter hat gewonnen, wenn seine Position genau der Zielposition entspricht.

---

## Aufgabe 7 – Automatische Navigation

Jetzt soll der Roboter selbst durch das Labyrinth fahren.

Implementiere:

```go
func ChooseDirection(world World, robot Robot) Direction
```

Verwende die **Rechte-Hand-Regel**.

Stelle dir vor, der Roboter legt seine rechte Hand permanent an die Wand.

Die Reihenfolge lautet:

1. rechts, wenn dort frei ist
2. sonst geradeaus
3. sonst links
4. sonst umdrehen

Die Hilfsfunktionen sind bereits vorhanden:

```go
TurnRight(...)
TurnLeft(...)
TurnBack(...)
```

Du musst also keine komplizierte Wegsuche programmieren.

---

# Tests

Zu jeder Pflichtaufgabe gibt es Tests.

Starte sie mit:

```bash
go test ./...
```

Zu Beginn schlagen mehrere Tests absichtlich fehl.

Du kannst auch nur einen einzelnen Test ausführen:

```bash
go test -run TestDirectionVector
```

oder beispielsweise:

```bash
go test -run TestMoveRobot
```

---

# Empfohlene Reihenfolge

Arbeite die Aufgaben möglichst in dieser Reihenfolge ab:

```text
1. DirectionVector
        ↓
2. IsInside
        ↓
3. IsWall
        ↓
4. CanMove
        ↓
5. MoveRobot
        ↓
6. HasReachedGoal
        ↓
7. ChooseDirection
```

Nach Aufgabe 5 kannst du den Roboter bereits **manuell durch das Labyrinth fahren**.

Die automatische Navigation kommt erst ganz am Ende.

---

# Bonusaufgaben

Wenn du mit den Pflichtaufgaben fertig bist, findest du weitere Aufgaben in:

```text
bonus.go
```

Dort kannst du beispielsweise:

- freie Nachbarfelder zählen
- die Manhattan-Distanz berechnen
- eine eigene Navigationsstrategie ausprobieren

Interessant ist besonders der Vergleich:

> Funktioniert eine Strategie, die immer versucht näher zum Ziel zu kommen,
> genauso zuverlässig wie die Rechte-Hand-Regel?

---

## Tipp

Versuche nicht, mehrere Aufgaben gleichzeitig zu lösen.

Implementiere eine Funktion, führe die Tests aus und mache erst dann weiter.

So kannst du Fehler leichter finden.
