# Snake

Dieses Projekt ist eine Abschlussaufgabe für einen Go-Einsteigerkurs. Das Fenster, die Tastatureingabe und die Game-Loop sind bereits fertig. Die eigentliche Spiellogik wird in **`game.go`** implementiert.

## Lernziele

Du verwendest dabei insbesondere:

- Datentypen und Variablen
- Funktionen
- Bedingungen
- Schleifen
- Slices
- Structs
- Pointer

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

## Aufgabe 1 – `ChangeDirection`

Ändere die Bewegungsrichtung der Schlange.

Dabei darf sich die Schlange nicht direkt um 180° drehen. Eine Schlange, die nach rechts läuft, darf also nicht sofort nach links laufen.

**Themen:** Structs, Pointer, Bedingungen

## Aufgabe 2 – `MoveSnake`

Bewege die Schlange genau ein Feld in ihre aktuelle Richtung.

Dabei soll der Kopf eine neue Position erhalten und jedes Körpersegment die vorherige Position seines Vorgängers übernehmen.

**Themen:** Pointer, Slices, Schleifen

## Aufgabe 3 – `HasWallCollision`

Prüfe, ob der Kopf das Spielfeld verlassen hat.

**Themen:** Bedingungen, Structs

## Aufgabe 4 – `HasSelfCollision`

Prüfe, ob der Kopf auf einem anderen Körpersegment liegt.

**Themen:** Slices, Schleifen, Bedingungen

## Aufgabe 5 – `HasEatenFood`

Prüfe, ob der Kopf auf derselben Position wie das Futter liegt.

**Themen:** Structs, Bedingungen

## Aufgabe 6 – `GrowSnake`

Verlängere die Schlange um ein Körpersegment.

**Themen:** Slices, Pointer, `append`

---

# Bedienung

| Taste | Funktion |
|---|---|
| `SPACE` | Spiel starten / pausieren |
| Pfeiltasten | Richtung ändern |
| `WASD` | Richtung ändern |
| `R` | Spiel neu starten |

---

# Freiwillige Bonusaufgaben

In **`bonus.go`** befinden sich zwei zusätzliche Funktionen:

- `CountFreeCells` – freie Spielfelder zählen
- `WrapPoint` – Spielfeldränder miteinander verbinden

Weitere Erweiterungsideen:

- Schlange mit wachsendem Score schneller werden lassen
- Hindernisse hinzufügen
- verschiedene Schwierigkeitsgrade
- Futter mit unterschiedlicher Punktzahl
- zweites Spezialfutter mit Zeitlimit
- Highscore innerhalb der laufenden Anwendung
- Spielmodus ohne tödliche Wände mit `WrapPoint`

---

# Projektstruktur

```text
Snake_Workshop/
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
