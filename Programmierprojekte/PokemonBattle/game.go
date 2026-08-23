package main

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	WindowWidth  = 900
	WindowHeight = 620

	PotionHeal = 35
)

type BattleState int

const (
	PlayerTurn BattleState = iota
	EnemyTurn
	BattleWon
	BattleLost
)

type Game struct {
	Player Monster
	Enemy  Monster

	State BattleState

	Message     string
	SecondLine  string
	turnReadyAt time.Time
}

func NewGame() *Game {
	g := &Game{}
	g.Reset()
	return g
}

func newPlayerMonster() Monster {
	return Monster{
		Name:    "Voltling",
		HP:      110,
		MaxHP:   110,
		Attack:  18,
		Defense: 7,
		Potions: 2,
		Moves: []Move{
			{Name: "Funkenstoß", Power: 7},
			{Name: "Kraftschlag", Power: 11},
			{Name: "Blitzhieb", Power: 15},
		},
	}
}

func newEnemyMonster() Monster {
	return Monster{
		Name:    "Pyronox",
		HP:      120,
		MaxHP:   120,
		Attack:  17,
		Defense: 9,
		Potions: 1,
		Moves: []Move{
			{Name: "Glut", Power: 6},
			{Name: "Feuerklaue", Power: 13},
			{Name: "Flammenstoß", Power: 16},
		},
	}
}

func (g *Game) Reset() {
	g.Player = newPlayerMonster()
	g.Enemy = newEnemyMonster()
	g.State = PlayerTurn
	g.Message = "Wähle eine Aktion."
	g.SecondLine = "1/2/3: Attacke   H: Heiltrank"
	g.turnReadyAt = time.Now()
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.Reset()
		return nil
	}

	switch g.State {
	case PlayerTurn:
		g.updatePlayerTurn()
	case EnemyTurn:
		if time.Now().After(g.turnReadyAt) {
			g.performEnemyTurn()
		}
	}

	return nil
}

func (g *Game) updatePlayerTurn() {
	switch {
	case inpututil.IsKeyJustPressed(ebiten.Key1):
		g.playerAttack(0)
	case inpututil.IsKeyJustPressed(ebiten.Key2):
		g.playerAttack(1)
	case inpututil.IsKeyJustPressed(ebiten.Key3):
		g.playerAttack(2)
	case inpututil.IsKeyJustPressed(ebiten.KeyH):
		g.playerPotion()
	}
}

func (g *Game) playerAttack(moveIndex int) {
	if !IsValidMove(g.Player, moveIndex) {
		g.Message = "Diese Attacke existiert nicht."
		return
	}

	move := g.Player.Moves[moveIndex]
	damage := PerformAttack(g.Player, &g.Enemy, move)

	g.Message = fmt.Sprintf("%s benutzt %s!", g.Player.Name, move.Name)
	g.SecondLine = fmt.Sprintf("%s verliert %d HP.", g.Enemy.Name, damage)

	if IsDefeated(g.Enemy) {
		g.State = BattleWon
		g.Message = fmt.Sprintf("%s wurde besiegt!", g.Enemy.Name)
		g.SecondLine = "Du hast gewonnen!   R: Neustart"
		return
	}

	g.startEnemyTurn()
}

func (g *Game) playerPotion() {
	healed := UsePotion(&g.Player, PotionHeal)

	if healed == 0 {
		if g.Player.Potions <= 0 {
			g.Message = "Keine Heiltränke mehr vorhanden."
		} else {
			g.Message = "Dein Monster hat bereits volle HP."
		}
		g.SecondLine = "Wähle eine andere Aktion."
		return
	}

	g.Message = fmt.Sprintf("%s benutzt einen Heiltrank.", g.Player.Name)
	g.SecondLine = fmt.Sprintf("%d HP geheilt. Noch %d Tränke.", healed, g.Player.Potions)

	g.startEnemyTurn()
}

func (g *Game) startEnemyTurn() {
	g.State = EnemyTurn
	g.turnReadyAt = time.Now().Add(850 * time.Millisecond)
}

func (g *Game) performEnemyTurn() {
	action := ChooseEnemyAction(g.Enemy)

	if action == -1 {
		healed := UsePotion(&g.Enemy, PotionHeal)

		if healed > 0 {
			g.Message = fmt.Sprintf("%s benutzt einen Heiltrank.", g.Enemy.Name)
			g.SecondLine = fmt.Sprintf("%s heilt %d HP.", g.Enemy.Name, healed)
		} else {
			g.Message = fmt.Sprintf("%s zögert...", g.Enemy.Name)
			g.SecondLine = "Keine Aktion möglich."
		}
	} else if IsValidMove(g.Enemy, action) {
		move := g.Enemy.Moves[action]
		damage := PerformAttack(g.Enemy, &g.Player, move)

		g.Message = fmt.Sprintf("%s benutzt %s!", g.Enemy.Name, move.Name)
		g.SecondLine = fmt.Sprintf("%s verliert %d HP.", g.Player.Name, damage)
	}

	if IsDefeated(g.Player) {
		g.State = BattleLost
		g.Message = fmt.Sprintf("%s wurde besiegt!", g.Player.Name)
		g.SecondLine = "Du hast verloren.   R: Neustart"
		return
	}

	g.State = PlayerTurn
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return WindowWidth, WindowHeight
}
