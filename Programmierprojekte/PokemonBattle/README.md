# Projekt: Monster Battle

## Ziel

In diesem Projekt programmierst du die Kampfmechanik eines kleinen
rundenbasierten Monster-Spiels.

Die Oberfläche, Tastatureingabe und der allgemeine Spielablauf sind bereits
fertig. Du implementierst das **Backend des Kampfsystems**.

Das Projekt ist bewusst an bekannte Monster-Kampfspiele angelehnt, verwendet
aber eigene Figuren und keine externen Grafiken.

---

## Das übst du

Besonders wichtig sind:

- Structs
- Slices
- Pointer
- Funktionen
- Bedingungen
- Schleifen
- mehrere zusammenarbeitende Funktionen

---

# Starten

Diesen Ordner in VS Code öffnen, anschließend

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

Im Kampf hast du drei Attacken und Heiltränke.

```text
1  erste Attacke
2  zweite Attacke
3  dritte Attacke
H  Heiltrank
R  Kampf neu starten
```

Der Gegner führt seinen Zug automatisch aus.

---

# Wo arbeite ich?

Alle Pflichtaufgaben befinden sich in:

```text
battle.go
```

Suche nach:

```go
// TODO
```

Die Benutzeroberfläche befindet sich hauptsächlich in:

```text
ui.go
game.go
```

Diese Dateien musst du für die Pflichtaufgaben **nicht verändern**.

---

# Datenmodell

Ein Monster wird durch ein Struct beschrieben:

```go
type Monster struct {
    Name     string
    HP       int
    MaxHP    int
    Attack   int
    Defense  int
    Potions  int
    Moves    []Move
}
```

Eine Attacke ist ebenfalls ein Struct:

```go
type Move struct {
    Name  string
    Power int
}
```

Beachte besonders:

```go
Moves []Move
```

Ein Monster besitzt also einen **Slice aus Attacken**.

---

# Aufgabe 1 – Besiegt?

Implementiere:

```go
func IsDefeated(monster Monster) bool
```

Ein Monster ist besiegt, wenn:

```text
HP <= 0
```

---

# Aufgabe 2 – Schaden berechnen

Implementiere:

```go
func CalculateDamage(
    attacker Monster,
    defender Monster,
    move Move,
) int
```

Verwende die vereinfachte Formel:

```text
Schaden = Angriff + Attackenstärke - Verteidigung
```

Beispiel:

```text
Attack       = 20
Move.Power   = 10
Defense      = 8

Schaden      = 22
```

Der Schaden darf niemals kleiner als `1` sein.

---

# Aufgabe 3 – Angreifen

Implementiere:

```go
func PerformAttack(
    attacker Monster,
    defender *Monster,
    move Move,
) int
```

Jetzt wird der Schaden tatsächlich von den HP abgezogen.

Achte darauf:

```text
HP dürfen niemals kleiner als 0 werden.
```

Hier wird bewusst ein Pointer verwendet:

```go
defender *Monster
```

Dadurch kannst du das ursprüngliche Monster verändern.

---

# Aufgabe 4 – Heilen

Implementiere:

```go
func Heal(monster *Monster, amount int) int
```

Beispiel:

```text
Vorher:   70 / 100 HP
Heilung:  50

Nachher: 100 / 100 HP
```

Die Funktion soll in diesem Fall `30` zurückgeben,
weil tatsächlich nur 30 HP geheilt wurden.

---

# Aufgabe 5 – Heiltrank

Implementiere:

```go
func UsePotion(monster *Monster, healAmount int) int
```

Regeln:

- kein Trank vorhanden → nichts passiert
- volle HP → kein Trank wird verbraucht
- ansonsten wird geheilt
- `Potions` wird um 1 reduziert

Benutze dabei möglichst deine bereits geschriebene Funktion:

```go
Heal(...)
```

---

# Aufgabe 6 – Attackenindex prüfen

Implementiere:

```go
func IsValidMove(monster Monster, moveIndex int) bool
```

Ein gültiger Index muss innerhalb des Slices liegen.

Für drei Attacken gelten beispielsweise:

```text
0  gültig
1  gültig
2  gültig

3  ungültig
-1 ungültig
```

---

# Aufgabe 7 – Gegnerentscheidung

Implementiere:

```go
func ChooseEnemyAction(monster Monster) int
```

Der Gegner soll eine kleine Entscheidungslogik bekommen.

Rückgabewerte:

```text
-1       Heiltrank
0..n     Index einer Attacke
```

Regeln:

### Wenig HP

Wenn das Monster höchstens ein Drittel seiner maximalen HP besitzt
und noch einen Trank hat:

```text
→ heilen
```

### Sonst

Suche im Slice:

```go
monster.Moves
```

die Attacke mit der größten `Power`.

Hier brauchst du eine Schleife.

---

# Empfohlene Reihenfolge

```text
IsDefeated
    ↓
CalculateDamage
    ↓
PerformAttack
    ↓
Heal
    ↓
UsePotion
    ↓
IsValidMove
    ↓
ChooseEnemyAction
```

Teste nach jeder Aufgabe:

```bash
go test ./...
```

Du kannst auch einzelne Tests starten:

```bash
go test -run TestCalculateDamage
```

oder:

```bash
go test -run TestUsePotion
```

---

# Bonusaufgaben

In:

```text
bonus.go
```

findest du zusätzliche Aufgaben.

Unter anderem:

- durchschnittliche Attackenstärke berechnen
- stärkste Attacke als `Move` zurückgeben
- eine komplette Kampfrunde ohne GUI simulieren

Diese Aufgaben sind nicht notwendig, um das Spiel fertigzustellen.

---

# Tipp

Die Oberfläche ist nur die Darstellung.

Die eigentliche Logik steckt in kleinen normalen Go-Funktionen:

```text
Eingabe
  ↓
game.go
  ↓
battle.go
  ↓
Monster wird verändert
  ↓
ui.go zeichnet den neuen Zustand
```

Dadurch kannst du fast die gesamte Kampfmechanik mit normalen Unit Tests prüfen.
