package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	WindowWidth  = 980
	WindowHeight = 660
)

type Game struct {
	Rooms  []Room
	Player Player

	SelectedItem int
	SelectedExit int

	Message string
	Won     bool
}

func NewGame() *Game {
	g := &Game{}
	g.Reset()
	return g
}

func (g *Game) Reset() {
	g.Rooms = CreateWorld()
	g.Player = Player{
		CurrentRoom: "office",
	}
	g.SelectedItem = 0
	g.SelectedExit = 0
	g.Message = "Finde einen Weg nach draußen."
	g.Won = false
}

func (g *Game) currentRoom() *Room {
	return FindRoom(g.Rooms, g.Player.CurrentRoom)
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.Reset()
		return nil
	}

	if g.Won {
		return nil
	}

	room := g.currentRoom()
	if room == nil {
		g.Message = "Fehler: aktueller Raum existiert nicht."
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyW) || inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.moveSelection(-1, room)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyS) || inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.moveSelection(1, room)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		g.takeSelectedItem(room)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.useSelectedExit(room)
	}

	return nil
}

func (g *Game) moveSelection(delta int, room *Room) {
	total := len(room.Items) + len(room.Exits)
	if total == 0 {
		return
	}

	flat := g.SelectedItem
	if g.SelectedExit >= 0 {
		flat = len(room.Items) + g.SelectedExit
	}

	flat += delta
	if flat < 0 {
		flat = total - 1
	}
	if flat >= total {
		flat = 0
	}

	if flat < len(room.Items) {
		g.SelectedItem = flat
		g.SelectedExit = -1
	} else {
		g.SelectedItem = -1
		g.SelectedExit = flat - len(room.Items)
	}
}

func (g *Game) ensureSelection(room *Room) {
	total := len(room.Items) + len(room.Exits)
	if total == 0 {
		g.SelectedItem = -1
		g.SelectedExit = -1
		return
	}

	if g.SelectedItem >= 0 && g.SelectedItem < len(room.Items) {
		return
	}
	if g.SelectedExit >= 0 && g.SelectedExit < len(room.Exits) {
		return
	}

	if len(room.Items) > 0 {
		g.SelectedItem = 0
		g.SelectedExit = -1
	} else {
		g.SelectedItem = -1
		g.SelectedExit = 0
	}
}

func (g *Game) takeSelectedItem(room *Room) {
	g.ensureSelection(room)

	if g.SelectedItem < 0 || g.SelectedItem >= len(room.Items) {
		g.Message = "Wähle zuerst einen Gegenstand."
		return
	}

	item := room.Items[g.SelectedItem]

	if TakeItem(&g.Player, room, item.ID) {
		g.Message = fmt.Sprintf("%s aufgenommen.", item.Name)
	} else {
		g.Message = "Gegenstand konnte nicht aufgenommen werden."
	}

	g.ensureSelection(room)
}

func (g *Game) useSelectedExit(room *Room) {
	g.ensureSelection(room)

	if g.SelectedExit < 0 || g.SelectedExit >= len(room.Exits) {
		g.Message = "Wähle zuerst einen Ausgang."
		return
	}

	exit := room.Exits[g.SelectedExit]

	moved, message := MovePlayer(&g.Player, g.Rooms, exit.Name)
	if message != "" {
		g.Message = message
	}

	if moved {
		g.SelectedItem = 0
		g.SelectedExit = -1

		if IsEscaped(g.Player) {
			g.Won = true
			g.Message = "Geschafft! Du bist entkommen. R: Neustart"
			return
		}

		if newRoom := g.currentRoom(); newRoom != nil {
			g.Message = fmt.Sprintf("Du betrittst: %s", newRoom.Name)
			g.ensureSelection(newRoom)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return WindowWidth, WindowHeight
}
