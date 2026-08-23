# Conway's Game of Life

Dieses Projekt ist eine Abschlussaufgabe für einen Go-Einsteigerkurs. Die grafische Oberfläche ist bereits fertig. Die eigentliche Spiellogik wird in **`game.go`** implementiert.

## Lernziele

Du verwendest dabei insbesondere:

- Datentypen und Variablen
- Funktionen
- Bedingungen
- Schleifen
- Slices / zweidimensionale Slices
- Structs
- Pointer

---

## 1. Projekt starten

Öffne **diesen Ordner** in VS Code.

Im Terminal:

```bash
go mod download
go run .
```

Beim ersten Start lädt Go die Bibliothek **Ebitengine** automatisch herunter.

Das Programm kompiliert bereits, obwohl die Aufgaben noch nicht gelöst sind. Die Simulation funktioniert zu Beginn aber noch nicht korrekt.

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

## Aufgabe 1 – `ToggleCell`

Schalte eine Zelle zwischen **lebendig** und **tot** um.

Dabei wird erstmals ein Pointer verwendet, weil die bestehende Welt verändert werden soll.

## Aufgabe 2 – `CountNeighbours`

Zähle die lebenden Nachbarzellen einer Position.

Eine Zelle kann höchstens acht Nachbarn besitzen. Außerhalb des Spielfelds existieren keine Zellen.

## Aufgabe 3 – `WillBeAlive`

Implementiere die Regeln von Conway's Game of Life.

### Lebende Zelle

- weniger als 2 Nachbarn → stirbt
- 2 oder 3 Nachbarn → überlebt
- mehr als 3 Nachbarn → stirbt

### Tote Zelle

- genau 3 Nachbarn → wird lebendig
- sonst → bleibt tot

## Aufgabe 4 – `NextGeneration`

Berechne aus einer vorhandenen Welt die nächste Generation.

**Wichtig:** Erzeuge dafür eine neue Welt. Verändere die alte Welt nicht während der Berechnung.

---

# Bedienung

| Taste | Funktion |
|---|---|
| `SPACE` | Simulation starten / pausieren |
| `N` | genau eine Generation berechnen |
| `R` | Startmuster zurücksetzen |
| `C` | Spielfeld leeren |
| Linke Maustaste | Zelle umschalten |

---

# Freiwillige Bonusaufgaben

Wenn du früher fertig bist, findest du in **`bonus.go`** weitere TODOs:

- `CountAlive` – lebende Zellen zählen
- `WorldsEqual` – zwei Welten vergleichen
- `ClearWorld` – Welt über einen Pointer leeren

Weitere mögliche Erweiterungen:

- Zufälliges Startmuster erzeugen
- Generationen pro Sekunde einstellbar machen
- Simulation automatisch stoppen, wenn sich nichts mehr verändert
- eigene bekannte Game-of-Life-Muster ergänzen
- andere Regeln ausprobieren

---

# Projektstruktur

```text
GameOfLife_Workshop/
├── README.md
├── go.mod
├── main.go          # fertiger Programmeinstieg
├── ui.go            # fertige Ebitengine-Oberfläche
├── patterns.go      # fertige Startmuster
├── game.go          # <-- HIER ARBEITEN
├── game_test.go     # Tests für die Pflichtaufgaben
├── bonus.go         # freiwillige Zusatzaufgaben
└── engine/
    ├── drawing.go   # wiederverwendbare Grafik-Helfer
    ├── input.go     # wiederverwendbare Eingabe-Helfer
    └── input_test.go
```

Die Dateien `ui.go` und `engine/` dürfen natürlich angesehen werden. Für die Aufgaben müssen sie aber nicht verstanden oder verändert werden.

---
