package main

// Product beschreibt einen Artikel im Lager.
//
// PriceCents wird absichtlich als int gespeichert.
// Geldbeträge sollten nicht unnötig als float64 gespeichert werden,
// da Gleitkommazahlen Rundungsfehler besitzen können.
type Product struct {
	ID         string
	Name       string
	Quantity   int
	Minimum    int
	PriceCents int
}

// Warehouse enthält alle Produkte.
type Warehouse struct {
	Products []Product
}

// FindProductIndex sucht nach einer Produkt-ID.
//
// Rückgabe:
//
//	0..n  Index des Produkts
//	-1    Produkt wurde nicht gefunden
//
// AUFGABE 1:
// Durchlaufe warehouse.Products mit einer Schleife.
func FindProductIndex(warehouse Warehouse, id string) int {
	// TODO
	return -1
}

// FindProduct sucht ein Produkt und gibt einen Pointer darauf zurück.
//
// Wenn das Produkt nicht existiert, soll nil zurückgegeben werden.
//
// WICHTIG:
// Der Pointer soll auf das Produkt IM Slice zeigen.
// So kann der Aufrufer das echte Produkt im Lager verändern.
//
// AUFGABE 2:
// Nutze FindProductIndex.
func FindProduct(warehouse *Warehouse, id string) *Product {
	// TODO
	return nil
}

// AddStock lagert eine bestimmte Menge ein.
//
// Regeln:
// - amount muss größer als 0 sein.
// - Bei Erfolg wird product.Quantity erhöht.
// - Rückgabe true bei Erfolg, sonst false.
//
// AUFGABE 3:
// Hier wird der Pointer benutzt, um das Originalprodukt zu verändern.
func AddStock(product *Product, amount int) bool {
	// TODO
	return false
}

// RemoveStock lagert eine bestimmte Menge aus.
//
// Regeln:
// - amount muss größer als 0 sein.
// - Es darf nie mehr ausgelagert werden als vorhanden ist.
// - Quantity darf niemals negativ werden.
// - Rückgabe true bei Erfolg, sonst false.
//
// AUFGABE 4:
func RemoveStock(product *Product, amount int) bool {
	// TODO
	return false
}

// IsLowStock prüft, ob ein Produkt nachbestellt werden sollte.
//
// Ein Produkt gilt als knapp, wenn:
//
//	Quantity <= Minimum
//
// AUFGABE 5:
func IsLowStock(product Product) bool {
	// TODO
	return false
}

// TotalQuantity berechnet die Gesamtanzahl aller Artikel im Lager.
//
// Beispiel:
// Produkt A: 10 Stück
// Produkt B: 20 Stück
// Ergebnis: 30
//
// AUFGABE 6:
func TotalQuantity(warehouse Warehouse) int {
	// TODO
	return 0
}

// InventoryValueCents berechnet den gesamten Lagerwert in Cent.
//
// Für jedes Produkt gilt:
//
//	Quantity * PriceCents
//
// Danach werden alle Werte addiert.
//
// AUFGABE 7:
func InventoryValueCents(warehouse Warehouse) int {
	// TODO
	return 0
}

// AddProduct fügt ein neues Produkt zum Lager hinzu.
//
// Regeln:
// - ID darf nicht leer sein.
// - Name darf nicht leer sein.
// - Quantity, Minimum und PriceCents dürfen nicht negativ sein.
// - Eine ID darf nur einmal vorkommen.
// - Bei Erfolg wird das Produkt an warehouse.Products angehängt.
//
// AUFGABE 8:
// Verwende append und möglichst FindProductIndex.
func AddProduct(warehouse *Warehouse, product Product) bool {
	// TODO
	return false
}
