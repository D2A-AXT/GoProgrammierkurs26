# Go Playground

Dieser Playground ist für den Einstieg in Go gedacht.  
Es ist nicht schlimm, wenn du noch nie programmiert hast. Die Aufgaben sind dafür da, Dinge auszuprobieren, Code zu verändern und zu beobachten, was passiert.

## 1. Ordner öffnen

Öffne den Ordner mit dem Playground in **Visual Studio Code**.

Der Ordner sollte ungefähr so aussehen:

```text
GoPlayground/
└── main.go
```

Wichtig: Öffne in VS Code den **gesamten Ordner**, nicht nur die Datei `main.go`.

## 2. Terminal öffnen

Öffne in Visual Studio Code ein Terminal:

**Terminal → New Terminal**

Unten in VS Code sollte sich jetzt ein Terminal öffnen.

Prüfe, ob du dich im richtigen Ordner befindest.

Zum Beispiel:

```text
...\GoPlayground>
```

## 3. Prüfen, ob Go installiert ist

Gib im Terminal ein:

```bash
go version
```

Wenn Go korrekt installiert ist, sollte ungefähr so etwas erscheinen:

```text
go version go1.xx.x windows/amd64
```

Die genaue Versionsnummer kann anders sein.

## 4. Go-Modul erstellen

Dieser Schritt ist nur nötig, wenn noch keine `go.mod` Datei im Ordner vorhanden ist.

Bevor du das Programm mit `go run .` starten kannst, erstellen wir einmalig eine `go.mod`-Datei.

Gib im Terminal ein:

```bash
go mod init playground
```

Danach sollte dein Ordner so aussehen:

```text
GoPlayground/
├── go.mod
└── main.go
```

Diesen Schritt musst du nur **einmal** machen.

## 5. Das Programm zum ersten Mal starten

Starte das Programm mit:

```bash
go run .
```

Wenn alles funktioniert, solltest du eine Ausgabe im Terminal sehen.

Zum Beispiel:

```text
Aufgabe 1: Sortiert...
```

Glückwunsch – du hast dein erstes Go-Programm ausgeführt.

## 6. Mit den Aufgaben beginnen

Öffne jetzt die Datei `main.go`.

Ganz oben in der Funktion `main()` findest du:

```go
exercise := 1
```

Die Zahl bestimmt, welche Aufgabe gestartet wird.

Für Aufgabe 1:

```go
exercise := 1
```

Für Aufgabe 2:

```go
exercise := 2
```

und so weiter.

Nach einer Änderung:

1. Datei speichern mit `Strg + S`
2. Im Terminal erneut ausführen:

```bash
go run .
```

## 7. Was tun, wenn ein Fehler erscheint?

Fehler gehören beim Programmieren dazu.

Wenn etwas nicht funktioniert:

1. Lies die Fehlermeldung.
2. Schau dir die angegebene Zeilennummer an.
3. Überlege, was dort falsch sein könnte.
4. Ändere den Code.
5. Speichere die Datei.
6. Starte erneut:

```bash
go run .
```

Du darfst den Code ausdrücklich kaputt machen und ausprobieren, was passiert.

## Die wichtigsten Befehle

```bash
# Prüfen, ob Go installiert ist
go version

# Ein neues Go-Modul erstellen
go mod init playground

# Das aktuelle Programm starten
go run .
```

Viel Spaß beim Ausprobieren.
