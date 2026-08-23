# Pong

Dieses Projekt ist eine Abschlussaufgabe für einen Go-Einsteigerkurs. Fenster, Grafik, Tastatureingabe, Spielstand und Game-Loop sind bereits fertig. Die eigentliche Spiellogik wird in **`game.go`** implementiert.

Das Spiel ist klassisches **2-Spieler-Pong**:

- Spieler 1: `W` / `S`
- Spieler 2: Pfeiltasten
- erster Spieler mit 5 Punkten gewinnt

## Lernziele

Du verwendest dabei insbesondere:

- Datentypen und Variablen
- Funktionen
- Bedingungen
- Structs
- Pointer
- `float64`
- einfache 2D-Koordinaten und Geschwindigkeiten

---

## 1. Projekt vorbereiten

Öffne **diesen Ordner** in VS Code.

Danach kann das Programm gestartet werden:

```bash
go run .
```

---

## 2. Tests starten

```bash
go test ./...
```

Oder nutze in VS Code die Go-Test-Schaltflächen.

Zu Beginn schlagen mehrere Tests fehl. Implementiere die Aufgaben nacheinander, bis alle Tests grün sind.

---

# Aufgaben

Alle Pflichtaufgaben befinden sich in **`game.go`**.

## Aufgabe 1 – `MovePaddle`

Bewege einen Schläger nach oben oder unten und verhindere, dass er das Spielfeld verlässt.

**Themen:** Pointer, Structs, Bedingungen, `float64`

## Aufgabe 2 – `MoveBall`

Bewege den Ball entsprechend seiner X- und Y-Geschwindigkeit.

**Themen:** Pointer, Structs, Variablen

## Aufgabe 3 – `BounceOffHorizontalWalls`

Erkenne Kollisionen mit der oberen und unteren Spielfeldkante und lasse den Ball abprallen.

**Themen:** Bedingungen, Pointer, Vorzeichen

## Aufgabe 4 – `HasPaddleCollision`

Prüfe, ob sich Ball und Schläger überlappen. Der Ball wird dafür vereinfacht als kleines Rechteck behandelt.

**Themen:** Bedingungen, Koordinaten, Structs

## Aufgabe 5 – `BounceFromPaddle`

Lasse den Ball vom Schläger abprallen. Je nachdem, wo der Ball den Schläger trifft, verändert sich auch seine vertikale Flugrichtung.

**Themen:** Pointer, Berechnungen, Structs

## Aufgabe 6 – `DetectScore`

Erkenne, ob der Ball das Spielfeld vollständig links oder rechts verlassen hat und welcher Spieler dadurch einen Punkt erhält.

**Themen:** Bedingungen, eigene Datentypen

---

# Bedienung

| Taste | Funktion |
|---|---|
| `W` / `S` | Spieler 1 bewegen |
| `↑` / `↓` | Spieler 2 bewegen |
| `SPACE` | Spiel starten / pausieren |
| `R` | Match zurücksetzen |

Nach jedem Punkt pausiert das Spiel. Mit `SPACE` beginnt der nächste Ballwechsel.

---

# Freiwillige Bonusaufgaben

In **`bonus.go`** befinden sich zwei zusätzliche Aufgaben:

- `IncreaseBallSpeed` – Ball nach jedem Schlägertreffer schneller machen
- `MoveComputerPaddle` – einfache Computersteuerung für einen Schläger

Weitere Erweiterungsideen:

- Ein-Spieler-Modus gegen den Computer
- verschiedene Schwierigkeitsgrade
- Paddle wird bei höherem Spielstand kleiner
- Ball wird mit der Zeit schneller
- anderer Abprallwinkel abhängig vom Trefferpunkt
- Power-Ups
- Best-of-3-Matches
- Startmenü

---

# Projektstruktur

```text
Pong_Workshop/
├── README.md
├── go.mod
├── main.go          # fertiger Programmeinstieg
├── ui.go            # fertige Ebitengine-Oberfläche und Game-Loop
├── game.go          # <-- HIER ARBEITEN
├── game_test.go     # Tests für die Pflichtaufgaben
├── bonus.go         # freiwillige Zusatzaufgaben
└── engine/
    └── drawing.go   # fertige Grafik-Helfer
```

Die Dateien `ui.go` und `engine/` dürfen natürlich angesehen werden. Für die Aufgaben müssen sie aber nicht verstanden oder verändert werden.
