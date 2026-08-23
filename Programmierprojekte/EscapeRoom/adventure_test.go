package main

import "testing"

func testRooms() []Room {
	return []Room{
		{
			ID:   "start",
			Name: "Start",
			Items: []Item{
				{ID: "key", Name: "Schlüssel"},
				{ID: "coin", Name: "Münze"},
			},
			Exits: []Exit{
				{Name: "Tür", TargetRoom: "hall"},
			},
		},
		{
			ID:   "hall",
			Name: "Flur",
			Exits: []Exit{
				{
					Name:         "Ausgang",
					TargetRoom:   "outside",
					RequiredItem: "key",
					LockedText:   "Die Tür ist abgeschlossen.",
				},
			},
		},
		{
			ID:   "outside",
			Name: "Draußen",
		},
	}
}

func TestFindRoom(t *testing.T) {
	rooms := testRooms()

	room := FindRoom(rooms, "hall")
	if room == nil {
		t.Fatal("expected room")
	}
	if room.Name != "Flur" {
		t.Fatalf("name = %q, want Flur", room.Name)
	}

	if FindRoom(rooms, "missing") != nil {
		t.Fatal("missing room must return nil")
	}
}

func TestFindItemIndex(t *testing.T) {
	items := []Item{
		{ID: "a"},
		{ID: "b"},
		{ID: "c"},
	}

	if got := FindItemIndex(items, "b"); got != 1 {
		t.Fatalf("index = %d, want 1", got)
	}

	if got := FindItemIndex(items, "x"); got != -1 {
		t.Fatalf("index = %d, want -1", got)
	}
}

func TestHasItem(t *testing.T) {
	player := Player{
		Inventory: []Item{
			{ID: "key"},
		},
	}

	if !HasItem(player, "key") {
		t.Error("player should have key")
	}

	if HasItem(player, "coin") {
		t.Error("player should not have coin")
	}
}

func TestTakeItem(t *testing.T) {
	room := Room{
		Items: []Item{
			{ID: "key", Name: "Schlüssel"},
			{ID: "coin", Name: "Münze"},
		},
	}
	player := Player{}

	if !TakeItem(&player, &room, "key") {
		t.Fatal("expected TakeItem to succeed")
	}

	if len(player.Inventory) != 1 || player.Inventory[0].ID != "key" {
		t.Fatalf("inventory = %+v", player.Inventory)
	}

	if len(room.Items) != 1 || room.Items[0].ID != "coin" {
		t.Fatalf("room items = %+v", room.Items)
	}
}

func TestTakeItemMissing(t *testing.T) {
	room := Room{
		Items: []Item{{ID: "coin"}},
	}
	player := Player{}

	if TakeItem(&player, &room, "key") {
		t.Fatal("missing item must fail")
	}

	if len(player.Inventory) != 0 {
		t.Fatal("inventory must remain unchanged")
	}

	if len(room.Items) != 1 {
		t.Fatal("room items must remain unchanged")
	}
}

func TestCanUseExit(t *testing.T) {
	player := Player{
		Inventory: []Item{{ID: "key"}},
	}

	if !CanUseExit(player, Exit{}) {
		t.Error("exit without required item should be usable")
	}

	if !CanUseExit(player, Exit{RequiredItem: "key"}) {
		t.Error("player has required item")
	}

	if CanUseExit(player, Exit{RequiredItem: "card"}) {
		t.Error("player does not have required item")
	}
}

func TestFindExit(t *testing.T) {
	room := Room{
		Exits: []Exit{
			{Name: "Nord"},
			{Name: "Ost"},
		},
	}

	exit := FindExit(room, "Ost")
	if exit == nil {
		t.Fatal("expected exit")
	}

	if exit.Name != "Ost" {
		t.Fatalf("exit name = %q", exit.Name)
	}

	if FindExit(room, "West") != nil {
		t.Fatal("missing exit must return nil")
	}
}

func TestMovePlayer(t *testing.T) {
	rooms := testRooms()
	player := Player{CurrentRoom: "start"}

	moved, _ := MovePlayer(&player, rooms, "Tür")

	if !moved {
		t.Fatal("expected player to move")
	}

	if player.CurrentRoom != "hall" {
		t.Fatalf("room = %q, want hall", player.CurrentRoom)
	}
}

func TestMovePlayerLocked(t *testing.T) {
	rooms := testRooms()
	player := Player{CurrentRoom: "hall"}

	moved, message := MovePlayer(&player, rooms, "Ausgang")

	if moved {
		t.Fatal("locked exit must not move player")
	}

	if player.CurrentRoom != "hall" {
		t.Fatalf("room changed to %q", player.CurrentRoom)
	}

	if message != "Die Tür ist abgeschlossen." {
		t.Fatalf("message = %q", message)
	}
}

func TestMovePlayerLockedWithKey(t *testing.T) {
	rooms := testRooms()
	player := Player{
		CurrentRoom: "hall",
		Inventory: []Item{
			{ID: "key"},
		},
	}

	moved, _ := MovePlayer(&player, rooms, "Ausgang")

	if !moved {
		t.Fatal("player with key should move")
	}

	if player.CurrentRoom != "outside" {
		t.Fatalf("room = %q, want outside", player.CurrentRoom)
	}
}

func TestMovePlayerUnknownExit(t *testing.T) {
	rooms := testRooms()
	player := Player{CurrentRoom: "start"}

	moved, message := MovePlayer(&player, rooms, "Fenster")

	if moved {
		t.Fatal("unknown exit must not move")
	}

	if message == "" {
		t.Fatal("expected an explanatory message")
	}
}

func TestIsEscaped(t *testing.T) {
	if !IsEscaped(Player{CurrentRoom: "outside"}) {
		t.Error("outside must count as escaped")
	}

	if IsEscaped(Player{CurrentRoom: "hall"}) {
		t.Error("hall must not count as escaped")
	}
}
