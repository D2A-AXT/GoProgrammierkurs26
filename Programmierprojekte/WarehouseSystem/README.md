# Projekt: Lagerverwaltung

## Ziel

In diesem Projekt implementierst du das Backend einer kleinen Lagerverwaltung.

Anders als bei Snake, Pong oder dem Roboter-Labyrinth ist dies kein Spiel.
Die Anwendung besitzt trotzdem ein eigenes Fenster und visualisiert den Zustand
des Lagers direkt.

Die komplette Benutzeroberfläche ist bereits fertig.

Du implementierst die **Lagerlogik**.

---

# Das übst du

Besonders wichtig sind:

- Structs
- Slices
- Pointer
- Funktionen
- Schleifen
- Bedingungen
- `append`
- Suchen in einem Slice
- Daten verändern, ohne unnötige Kopien zu erzeugen

Außerdem siehst du ein kleines Beispiel dafür, wie man Fachlogik von einer
Benutzeroberfläche trennt.

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
↑ / ↓        Produkt auswählen
W / S        Produkt auswählen

E            10 Stück einlagern
A            1 Stück auslagern

N            einen Demo-Artikel neu anlegen
R            Lager zurücksetzen
```

Das Fenster zeigt zusätzlich:

- aktuellen Bestand
- Mindestbestand
- Einzelpreis
- Nachbestellstatus
- Gesamtanzahl aller Artikel
- gesamten Lagerwert

---

# Wo arbeite ich?

Alle Pflichtaufgaben befinden sich in:

```text
warehouse.go
```

Suche nach:

```go
// TODO
```

Die Dateien

```text
game.go
ui.go
main.go
```

enthalten hauptsächlich das fertige Framework.

Für die Pflichtaufgaben musst du sie nicht verändern.

---

# Datenmodell

Ein Produkt wird durch ein Struct dargestellt:

```go
type Product struct {
    ID         string
    Name       string
    Quantity   int
    Minimum    int
    PriceCents int
}
```

Das Lager besitzt einen Slice aus Produkten:

```go
type Warehouse struct {
    Products []Product
}
```

---

# Warum PriceCents?

Der Preis wird zum Beispiel nicht so gespeichert:

```go
Price float64
```

sondern:

```go
PriceCents int
```

Ein Preis von 12,90 EUR ist also:

```go
1290
```

Das verhindert typische Rundungsprobleme von Gleitkommazahlen bei Geldbeträgen.

---

# Aufgabe 1 – Produktindex suchen

Implementiere:

```go
func FindProductIndex(
    warehouse Warehouse,
    id string,
) int
```

Durchlaufe den Slice:

```go
warehouse.Products
```

und suche nach der passenden ID.

Wenn du das Produkt nicht findest:

```go
return -1
```

---

# Aufgabe 2 – Produkt finden

Implementiere:

```go
func FindProduct(
    warehouse *Warehouse,
    id string,
) *Product
```

Nutze möglichst deine Funktion:

```go
FindProductIndex(...)
```

Wenn das Produkt existiert, soll ein Pointer auf das Produkt im Slice
zurückgegeben werden.

Das ist wichtig.

Nicht eine Kopie:

```go
product := warehouse.Products[index]
return &product
```

sondern ein Pointer auf das echte Slice-Element:

```go
return &warehouse.Products[index]
```

Warum?

Dann kann der Aufrufer später beispielsweise:

```go
product.Quantity++
```

ausführen und damit tatsächlich den Lagerbestand verändern.

---

# Aufgabe 3 – Einlagern

Implementiere:

```go
func AddStock(
    product *Product,
    amount int,
) bool
```

Regeln:

- `amount > 0`
- bei Erfolg wird `Quantity` erhöht
- Rückgabe `true` bei Erfolg
- sonst `false`

Hier kommt ein Pointer sinnvoll zum Einsatz.

---

# Aufgabe 4 – Auslagern

Implementiere:

```go
func RemoveStock(
    product *Product,
    amount int,
) bool
```

Regeln:

```text
amount muss > 0 sein
```

und:

```text
amount darf nicht größer als der aktuelle Bestand sein
```

Der Bestand darf niemals negativ werden.

---

# Aufgabe 5 – Mindestbestand

Implementiere:

```go
func IsLowStock(product Product) bool
```

Ein Produkt soll nachbestellt werden, wenn:

```text
Quantity <= Minimum
```

Wenn die Funktion funktioniert, erscheint in der GUI automatisch:

```text
NACHBESTELLEN
```

---

# Aufgabe 6 – Gesamtbestand

Implementiere:

```go
func TotalQuantity(warehouse Warehouse) int
```

Addiere die Bestände aller Produkte.

Beispiel:

```text
Kabelbinder    50
Relais          7
Sensor          3
-----------------
Gesamt         60
```

---

# Aufgabe 7 – Lagerwert

Implementiere:

```go
func InventoryValueCents(
    warehouse Warehouse,
) int
```

Der Wert eines einzelnen Produkts ist:

```text
Quantity * PriceCents
```

Addiere anschließend die Werte aller Produkte.

Die GUI wandelt Cent später automatisch in Euro um.

---

# Aufgabe 8 – Neues Produkt anlegen

Implementiere:

```go
func AddProduct(
    warehouse *Warehouse,
    product Product,
) bool
```

Prüfe zuerst:

- ID nicht leer
- Name nicht leer
- Quantity nicht negativ
- Minimum nicht negativ
- PriceCents nicht negativ
- ID noch nicht vorhanden

Anschließend:

```go
warehouse.Products = append(
    warehouse.Products,
    product,
)
```

Wenn diese Aufgabe funktioniert, kannst du im Programm `N` drücken.

Dann wird testweise ein neuer Artikel angelegt:

```text
P-500
Industrie-Switch
```

---

# Empfohlene Reihenfolge

```text
FindProductIndex
       ↓
FindProduct
       ↓
AddStock
       ↓
RemoveStock
       ↓
IsLowStock
       ↓
TotalQuantity
       ↓
InventoryValueCents
       ↓
AddProduct
```

Nach jeder Aufgabe:

```bash
go test ./...
```

---

# Einzelne Tests ausführen

Zum Beispiel:

```bash
go test -run TestFindProduct
```

oder:

```bash
go test -run TestRemoveStock
```

---

# Bonusaufgaben

In:

```text
bonus.go
```

findest du weitere Aufgaben.

Unter anderem:

- Anzahl aller Produkte mit zu niedrigem Bestand
- wertvollstes Produkt finden
- Bestand zwischen zwei Produkten verschieben

Die Bonusaufgaben werden für das fertige Hauptprogramm nicht benötigt.

---

# Architektur

Das Projekt ist bewusst getrennt:

```text
Tastatureingabe
      ↓
   game.go
      ↓
warehouse.go
      ↓
Warehouse / Product
      ↓
    ui.go
```

`warehouse.go` kennt Ebitengine überhaupt nicht.

Dadurch lässt sich die komplette Fachlogik unabhängig von der GUI testen.

Das ist ein Prinzip, das auch bei größeren realen Programmen sehr nützlich ist.
