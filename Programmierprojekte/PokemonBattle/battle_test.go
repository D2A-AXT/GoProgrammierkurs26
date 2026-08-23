package main

import "testing"

func sampleMonster() Monster {
	return Monster{
		Name:    "Flamara",
		HP:      80,
		MaxHP:   100,
		Attack:  20,
		Defense: 8,
		Potions: 2,
		Moves: []Move{
			{Name: "Tackle", Power: 5},
			{Name: "Fire Burst", Power: 12},
			{Name: "Quick Hit", Power: 8},
		},
	}
}

func TestIsDefeated(t *testing.T) {
	if IsDefeated(Monster{HP: 10}) {
		t.Error("monster with positive HP must not be defeated")
	}

	if !IsDefeated(Monster{HP: 0}) {
		t.Error("monster with 0 HP must be defeated")
	}

	if !IsDefeated(Monster{HP: -5}) {
		t.Error("monster with negative HP must be defeated")
	}
}

func TestCalculateDamage(t *testing.T) {
	attacker := Monster{Attack: 20}
	defender := Monster{Defense: 8}
	move := Move{Power: 10}

	got := CalculateDamage(attacker, defender, move)
	want := 22

	if got != want {
		t.Fatalf("damage = %d, want %d", got, want)
	}
}

func TestCalculateDamageMinimumOne(t *testing.T) {
	attacker := Monster{Attack: 2}
	defender := Monster{Defense: 100}
	move := Move{Power: 1}

	if got := CalculateDamage(attacker, defender, move); got != 1 {
		t.Fatalf("damage = %d, want minimum damage 1", got)
	}
}

func TestPerformAttack(t *testing.T) {
	attacker := Monster{Attack: 20}
	defender := Monster{HP: 50, Defense: 5}
	move := Move{Power: 10}

	damage := PerformAttack(attacker, &defender, move)

	if damage != 25 {
		t.Fatalf("damage = %d, want 25", damage)
	}

	if defender.HP != 25 {
		t.Fatalf("defender HP = %d, want 25", defender.HP)
	}
}

func TestPerformAttackDoesNotCreateNegativeHP(t *testing.T) {
	attacker := Monster{Attack: 50}
	defender := Monster{HP: 10, Defense: 0}
	move := Move{Power: 20}

	PerformAttack(attacker, &defender, move)

	if defender.HP != 0 {
		t.Fatalf("defender HP = %d, want 0", defender.HP)
	}
}

func TestHeal(t *testing.T) {
	monster := Monster{HP: 40, MaxHP: 100}

	healed := Heal(&monster, 25)

	if healed != 25 {
		t.Fatalf("healed = %d, want 25", healed)
	}

	if monster.HP != 65 {
		t.Fatalf("HP = %d, want 65", monster.HP)
	}
}

func TestHealClampsToMaxHP(t *testing.T) {
	monster := Monster{HP: 90, MaxHP: 100}

	healed := Heal(&monster, 50)

	if healed != 10 {
		t.Fatalf("healed = %d, want 10", healed)
	}

	if monster.HP != 100 {
		t.Fatalf("HP = %d, want 100", monster.HP)
	}
}

func TestUsePotion(t *testing.T) {
	monster := Monster{HP: 40, MaxHP: 100, Potions: 2}

	healed := UsePotion(&monster, 30)

	if healed != 30 {
		t.Fatalf("healed = %d, want 30", healed)
	}

	if monster.Potions != 1 {
		t.Fatalf("potions = %d, want 1", monster.Potions)
	}
}

func TestUsePotionAtFullHPDoesNotConsumePotion(t *testing.T) {
	monster := Monster{HP: 100, MaxHP: 100, Potions: 2}

	healed := UsePotion(&monster, 30)

	if healed != 0 {
		t.Fatalf("healed = %d, want 0", healed)
	}

	if monster.Potions != 2 {
		t.Fatalf("potions = %d, want 2", monster.Potions)
	}
}

func TestUsePotionWithoutPotions(t *testing.T) {
	monster := Monster{HP: 40, MaxHP: 100, Potions: 0}

	healed := UsePotion(&monster, 30)

	if healed != 0 {
		t.Fatalf("healed = %d, want 0", healed)
	}

	if monster.HP != 40 {
		t.Fatalf("HP = %d, want unchanged 40", monster.HP)
	}
}

func TestIsValidMove(t *testing.T) {
	monster := sampleMonster()

	if !IsValidMove(monster, 0) {
		t.Error("move 0 should be valid")
	}

	if !IsValidMove(monster, 2) {
		t.Error("move 2 should be valid")
	}

	if IsValidMove(monster, -1) {
		t.Error("negative move index must be invalid")
	}

	if IsValidMove(monster, 3) {
		t.Error("index equal to len(Moves) must be invalid")
	}
}

func TestChooseEnemyActionHealsAtLowHP(t *testing.T) {
	monster := sampleMonster()
	monster.HP = 30
	monster.MaxHP = 100
	monster.Potions = 1

	if got := ChooseEnemyAction(monster); got != -1 {
		t.Fatalf("action = %d, want -1 for potion", got)
	}
}

func TestChooseEnemyActionChoosesStrongestMove(t *testing.T) {
	monster := sampleMonster()
	monster.HP = 80

	// Fire Burst at index 1 has the highest Power.
	if got := ChooseEnemyAction(monster); got != 1 {
		t.Fatalf("action = %d, want strongest move index 1", got)
	}
}

func TestChooseEnemyActionWithoutMoves(t *testing.T) {
	monster := sampleMonster()
	monster.Moves = nil

	if got := ChooseEnemyAction(monster); got != -1 {
		t.Fatalf("action = %d, want -1 when there are no moves", got)
	}
}
