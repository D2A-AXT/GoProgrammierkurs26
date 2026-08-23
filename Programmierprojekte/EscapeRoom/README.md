# Projekt: Escape Room

## Ziel

In diesem Projekt programmierst du die Logik eines kleinen grafischen
Escape-Room-Adventures.

Die Oberfläche, Räume und Story sind bereits vorbereitet.
Du implementierst das Backend:

- Räume finden
- Gegenstände suchen
- Inventar verwalten
- Gegenstände aufnehmen
- verschlossene Türen prüfen
- Räume wechseln
- Sieg erkennen

---

# Das übst du

Besonders wichtig sind:

- Structs
- Slices
- Pointer
- Funktionen
- Bedingungen
- Schleifen
- `append`
- Elemente aus Slices entfernen
- mehrere zusammenarbeitende Datenstrukturen

---

# Starten

Diesen Ordner in VS Code öffnen

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
↑ / ↓ oder W / S    Auswahl bewegen
E                    Gegenstand aufnehmen
SPACE                Ausgang benutzen
R                    Neustart
```

Ziel:

> Finde die benötigten Gegenstände und verlasse das Gebäude.

---

# Wo arbeite ich?

Alle Pflichtaufgaben befinden sich in:

```text
adventure.go
```

Suche nach:

```go
// TODO
```

Die Dateien

```text
game.go
ui.go
world.go
```

sind bereits vorbereitet.

---

# Datenmodell

Ein Gegenstand:

```go
type Item struct {
    ID          string
    Name        string
    Description string
}
```

Ein Ausgang:

```go
type Exit struct {
    Name         string
    TargetRoom   string
    RequiredItem string
    LockedText   string
}
```

Ein Raum:

```go
type Room struct {
    ID          string
    Name        string
    Description string
    Items       []Item
    Exits       []Exit
}
```

Der Spieler:

```go
type Player struct {
    CurrentRoom string
    Inventory   []Item
}
```

---

# Aufgabe 1 – Raum suchen

Implementiere:

```go
func FindRoom(
    rooms []Room,
    id string,
) *Room
```

Durchlaufe alle Räume und suche nach der ID.

Wenn nichts gefunden wird:

```go
return nil
```

Achte darauf, einen Pointer auf den echten Raum im Slice zurückzugeben.

---

# Aufgabe 2 – Gegenstand suchen

Implementiere:

```go
func FindItemIndex(
    items []Item,
    id string,
) int
```

Rückgabe:

```text
0..n    gefunden
-1      nicht gefunden
```

---

# Aufgabe 3 – Inventar prüfen

Implementiere:

```go
func HasItem(
    player Player,
    itemID string,
) bool
```

Prüfe, ob sich der Gegenstand im Inventar befindet.

Nutze möglichst:

```go
FindItemIndex(...)
```

---

# Aufgabe 4 – Gegenstand aufnehmen

Implementiere:

```go
func TakeItem(
    player *Player,
    room *Room,
    itemID string,
) bool
```

Dabei müssen zwei Dinge passieren:

### 1. Gegenstand aus dem Raum entfernen

Beispiel:

```go
items = append(
    items[:index],
    items[index+1:]...,
)
```

### 2. Gegenstand ins Inventar legen

```go
player.Inventory = append(
    player.Inventory,
    item,
)
```

Diese Aufgabe ist eine gute Slice-Übung.

---

# Aufgabe 5 – Verschlossene Tür prüfen

Implementiere:

```go
func CanUseExit(
    player Player,
    exit Exit,
) bool
```

Wenn:

```go
exit.RequiredItem == ""
```

ist die Tür immer frei.

Sonst muss der Spieler den Gegenstand besitzen.

---

# Aufgabe 6 – Ausgang suchen

Implementiere:

```go
func FindExit(
    room Room,
    exitName string,
) *Exit
```

Durchlaufe:

```go
room.Exits
```

und suche nach dem passenden Namen.

---

# Aufgabe 7 – Raum wechseln

Implementiere:

```go
func MovePlayer(
    player *Player,
    rooms []Room,
    exitName string,
) (bool, string)
```

Die Funktion gibt **zwei Werte** zurück.

Beispiel:

```go
moved, message := MovePlayer(...)
```

Regeln:

### Ausgang existiert nicht

```text
moved = false
```

und eine passende Fehlermeldung.

### Ausgang verschlossen

```text
moved = false
message = exit.LockedText
```

### Ausgang frei

```go
player.CurrentRoom = exit.TargetRoom
```

und:

```text
moved = true
```

---

# Aufgabe 8 – Entkommen?

Implementiere:

```go
func IsEscaped(player Player) bool
```

Der Escape Room ist geschafft, wenn:

```go
player.CurrentRoom == "outside"
```

---

# Empfohlene Reihenfolge

```text
FindRoom
   ↓
FindItemIndex
   ↓
HasItem
   ↓
TakeItem
   ↓
CanUseExit
   ↓
FindExit
   ↓
MovePlayer
   ↓
IsEscaped
```

Nach jeder Aufgabe:

```bash
go test ./...
```

---

# Einzelne Tests

Zum Beispiel:

```bash
go test -run TestTakeItem
```

oder:

```bash
go test -run TestMovePlayer
```

---

# Bonusaufgaben

In:

```text
bonus.go
```

findest du zusätzliche Aufgaben:

- Gegenstände wieder aus dem Inventar entfernen
- verschlossene Ausgänge zählen
- prüfen, ob ein Raum direkt erreichbar ist

Keine Bonusaufgabe benötigt Rekursion.

---

# Architektur

Die eigentliche Adventure-Logik ist vollständig unabhängig von Ebitengine:

```text
ui.go
  ↓
game.go
  ↓
adventure.go
  ↓
Room / Item / Exit / Player
```

Dadurch lässt sich die Logik mit normalen Unit Tests prüfen.

Das ist auch bei größeren Programmen ein wichtiges Prinzip:
Die Benutzeroberfläche sollte nicht selbst die Fachlogik enthalten.
