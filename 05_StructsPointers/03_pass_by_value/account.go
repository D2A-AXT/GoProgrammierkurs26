package account

// Übung 3 – Structs werden kopiert
// Schwierigkeit: ★★☆☆☆
//
// Lernziele:
//   - Go arbeitet mit Pass by Value
//   - Auch ein Struct-Parameter ist zunächst eine Kopie
//   - Änderungen an der Kopie verändern den Aufrufer nicht

type Account struct {
	Owner   string
	Balance int
}

// DepositCopy erhält Account als WERT.
// Erhöhe die Balance der lokalen Kopie und gib diese Kopie zurück.
func DepositCopy(account Account, amount int) Account {
	// TODO
	return Account{}
}

// RenameCopy soll analog nur die Kopie umbenennen und zurückgeben.
func RenameCopy(account Account, newOwner string) Account {
	// TODO
	return Account{}
}

// Experiment:
//
// Ergänze testweise innerhalb einer Funktion:
//     fmt.Printf("Adresse lokal: %p\n", &account)
//
// und im Test/Playground:
//     fmt.Printf("Adresse original: %p\n", &original)
//
// Sind die Adressen gleich? Warum nicht?
//
// Dafür müsstest du vorübergehend `import "fmt"` ergänzen.
