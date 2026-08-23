package main

import "testing"

func TestToggleCell(t *testing.T) {
	world := NewWorld(3, 3)

	ToggleCell(&world, 1, 1)
	if !world.Cells[1][1] {
		t.Fatal("ToggleCell sollte eine tote Zelle lebendig machen")
	}

	ToggleCell(&world, 1, 1)
	if world.Cells[1][1] {
		t.Fatal("ToggleCell sollte eine lebendige Zelle wieder tot machen")
	}
}

func TestToggleCellOutsideWorldDoesNothing(t *testing.T) {
	world := NewWorld(2, 2)

	// Darf nicht abstürzen.
	ToggleCell(&world, -1, 0)
	ToggleCell(&world, 0, 99)
}

func TestCountNeighboursMiddle(t *testing.T) {
	world := NewWorld(3, 3)
	world.Cells[0][0] = true
	world.Cells[0][1] = true
	world.Cells[1][0] = true

	got := CountNeighbours(world, 1, 1)
	want := 3

	if got != want {
		t.Fatalf("CountNeighbours(..., 1, 1) = %d, erwartet %d", got, want)
	}
}

func TestCountNeighboursDoesNotCountCellItself(t *testing.T) {
	world := NewWorld(3, 3)
	world.Cells[1][1] = true

	got := CountNeighbours(world, 1, 1)
	if got != 0 {
		t.Fatalf("Die Zelle selbst darf nicht mitgezählt werden: erhalten %d", got)
	}
}

func TestCountNeighboursAtBorder(t *testing.T) {
	world := NewWorld(3, 3)
	world.Cells[0][1] = true
	world.Cells[1][0] = true
	world.Cells[1][1] = true

	got := CountNeighbours(world, 0, 0)
	want := 3

	if got != want {
		t.Fatalf("CountNeighbours am Rand = %d, erwartet %d", got, want)
	}
}

func TestWillBeAlive(t *testing.T) {
	tests := []struct {
		name           string
		currentlyAlive bool
		neighbours     int
		want           bool
	}{
		{"lebendig mit 0 Nachbarn stirbt", true, 0, false},
		{"lebendig mit 1 Nachbar stirbt", true, 1, false},
		{"lebendig mit 2 Nachbarn überlebt", true, 2, true},
		{"lebendig mit 3 Nachbarn überlebt", true, 3, true},
		{"lebendig mit 4 Nachbarn stirbt", true, 4, false},
		{"tot mit 2 Nachbarn bleibt tot", false, 2, false},
		{"tot mit 3 Nachbarn wird lebendig", false, 3, true},
		{"tot mit 4 Nachbarn bleibt tot", false, 4, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WillBeAlive(tt.currentlyAlive, tt.neighbours)
			if got != tt.want {
				t.Fatalf("WillBeAlive(%v, %d) = %v, erwartet %v",
					tt.currentlyAlive, tt.neighbours, got, tt.want)
			}
		})
	}
}

func TestNextGenerationBlinker(t *testing.T) {
	world := NewWorld(5, 5)
	world.Cells[2][1] = true
	world.Cells[2][2] = true
	world.Cells[2][3] = true

	next := NextGeneration(world)

	wantAlive := map[[2]int]bool{
		{1, 2}: true,
		{2, 2}: true,
		{3, 2}: true,
	}

	for row := 0; row < next.Rows(); row++ {
		for col := 0; col < next.Cols(); col++ {
			want := wantAlive[[2]int{row, col}]
			if next.Cells[row][col] != want {
				t.Fatalf("Zelle (%d,%d) = %v, erwartet %v", row, col, next.Cells[row][col], want)
			}
		}
	}
}

func TestNextGenerationDoesNotModifyInput(t *testing.T) {
	world := NewWorld(3, 3)
	world.Cells[1][0] = true
	world.Cells[1][1] = true
	world.Cells[1][2] = true

	_ = NextGeneration(world)

	if !world.Cells[1][0] || !world.Cells[1][1] || !world.Cells[1][2] {
		t.Fatal("NextGeneration darf die übergebene Welt nicht verändern")
	}
}
