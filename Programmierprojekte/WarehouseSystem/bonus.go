package main

// BONUS-AUFGABEN
//
// Diese Funktionen werden vom Hauptprogramm nicht benötigt.
// Sie sind für Teilnehmer gedacht, die früher fertig sind.

// CountLowStock zählt alle Produkte, die nachbestellt werden sollten.
//
// BONUS 1:
// Nutze IsLowStock.
func CountLowStock(warehouse Warehouse) int {
	// TODO
	return 0
}

// MostValuableProduct sucht das Produkt mit dem höchsten Gesamtwert.
//
// Gesamtwert eines Produkts:
//
//	Quantity * PriceCents
//
// BONUS 2:
// Gibt false zurück, wenn das Lager leer ist.
//
// Beispiel für mehrere Rückgabewerte:
//
//	product, ok := MostValuableProduct(warehouse)
func MostValuableProduct(warehouse Warehouse) (Product, bool) {
	// TODO
	return Product{}, false
}

// TransferStock verschiebt eine Menge von einem Produkt zu einem anderen.
//
// BONUS 3:
// Diese Funktion ist absichtlich etwas künstlich, aber eine gute Pointer-Übung.
//
// Regeln:
// - amount > 0
// - source muss genügend Bestand besitzen
// - bei Fehler darf sich KEIN Bestand ändern
func TransferStock(source *Product, target *Product, amount int) bool {
	// TODO
	return false
}
