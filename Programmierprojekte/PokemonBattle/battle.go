package main

// Move beschreibt eine Attacke.
type Move struct {
	Name  string
	Power int
}

// Monster beschreibt einen Kämpfer.
//
// HP      = aktuelle Lebenspunkte
// MaxHP   = maximale Lebenspunkte
// Attack  = Angriffswert
// Defense = Verteidigungswert
// Potions = verbleibende Heiltränke
type Monster struct {
	Name    string
	HP      int
	MaxHP   int
	Attack  int
	Defense int
	Potions int
	Moves   []Move
}

// IsDefeated prüft, ob ein Monster besiegt wurde.
//
// AUFGABE 1:
// Ein Monster ist besiegt, wenn seine HP kleiner oder gleich 0 sind.
func IsDefeated(monster Monster) bool {
	// TODO
	return false
}

// CalculateDamage berechnet den Schaden einer Attacke.
//
// Verwende diese vereinfachte Formel:
//
//	damage = attacker.Attack + move.Power - defender.Defense
//
// Der Schaden soll aber niemals kleiner als 1 sein.
//
// AUFGABE 2:
func CalculateDamage(attacker Monster, defender Monster, move Move) int {
	// TODO
	return 0
}

// PerformAttack führt eine Attacke aus und verändert die HP des Verteidigers.
//
// Regeln:
// - Berechne zuerst den Schaden mit CalculateDamage.
// - Ziehe den Schaden von defender.HP ab.
// - HP dürfen niemals kleiner als 0 werden.
// - Gib den tatsächlich verursachten Schaden zurück.
//
// AUFGABE 3:
// Warum ist defender hier ein Pointer?
func PerformAttack(attacker Monster, defender *Monster, move Move) int {
	// TODO
	return 0
}

// Heal heilt ein Monster.
//
// Regeln:
// - amount wird zu den HP addiert.
// - HP dürfen MaxHP nicht überschreiten.
// - Gib zurück, wie viele HP tatsächlich geheilt wurden.
//
// Beispiel:
// 70 / 100 HP + Heilung 50
// -> neue HP = 100
// -> Rückgabewert = 30
//
// AUFGABE 4:
func Heal(monster *Monster, amount int) int {
	// TODO
	return 0
}

// UsePotion verwendet einen Heiltrank.
//
// Regeln:
// - Wenn keine Tränke mehr vorhanden sind, passiert nichts.
// - Wenn das Monster bereits volle HP hat, soll KEIN Trank verbraucht werden.
// - Ein Trank heilt um healAmount.
// - Bei erfolgreicher Benutzung wird Potions um 1 reduziert.
// - Gib die tatsächlich geheilten HP zurück.
//
// AUFGABE 5:
func UsePotion(monster *Monster, healAmount int) int {
	// TODO
	return 0
}

// IsValidMove prüft, ob ein Index auf eine vorhandene Attacke zeigt.
//
// AUFGABE 6:
// Denke an negative Indizes und len(monster.Moves).
func IsValidMove(monster Monster, moveIndex int) bool {
	// TODO
	return false
}

// ChooseEnemyAction entscheidet, was der Gegner tun soll.
//
// Rückgabewerte:
//
//	 -1 = Heiltrank verwenden
//	0..n = Index einer Attacke
//
// Regeln:
//  1. Wenn das Monster höchstens ein Drittel seiner MaxHP besitzt
//     UND noch mindestens einen Trank besitzt, soll es heilen.
//  2. Sonst soll es die stärkste Attacke auswählen.
//  3. Falls das Monster keine Attacken besitzt, gib -1 zurück.
//
// AUFGABE 7:
// Hier werden Schleife, Bedingungen und Slices kombiniert.
func ChooseEnemyAction(monster Monster) int {
	// TODO
	return 0
}
