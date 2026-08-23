# Programmierkurs 2026 – Übungen

## Hinweise zur Bearbeitung

Die Übungen sind als ein gemeinsames Go-Modul aufgebaut.

Die Übungen zur jeweiligen Lektion finden sich im entsprechenden Ordner. Die Übungen sind so strukturiert, dass sie in der Schwirigkeit steigen. Nicht alle Aufgaben zu jeder Lektion müssen gelöst werden, sollten aber grundsätzlich in der richtigen Reihenfolge bearbeitet werden.

Außerdem enthalten manche Übungen zusätzliche Inhalte, die nicht in der Lektion besprochen wurden.

Öffne in VS Code den Ordner `Übungen`, damit die Go-Erweiterung alle Tests erkennt.

## Tests

Alle Übungen:

```bash
go test ./...
```

Nur Rekursion:

```bash
go test ./Recursion/...
```

Eine einzelne Übung:

```bash
go test ./Recursion/01_sum
```

Die Starterfunktionen enthalten absichtlich TODOs. Deshalb schlagen die Tests zunächst fehl.
