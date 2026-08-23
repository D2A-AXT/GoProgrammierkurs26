package main

// InitialWorld erzeugt ein kleines Startmuster, damit direkt nach
// dem Start etwas im Fenster zu sehen ist.
func InitialWorld(rows, cols int) World {
	world := NewWorld(rows, cols)

	// Glider
	setAlive(&world, 4, 5)
	setAlive(&world, 5, 6)
	setAlive(&world, 6, 4)
	setAlive(&world, 6, 5)
	setAlive(&world, 6, 6)

	// Blinker
	setAlive(&world, 12, 16)
	setAlive(&world, 12, 17)
	setAlive(&world, 12, 18)

	// Kleines statisches Quadrat (Still Life)
	setAlive(&world, 20, 28)
	setAlive(&world, 20, 29)
	setAlive(&world, 21, 28)
	setAlive(&world, 21, 29)

	return world
}

// setAlive ist nur eine interne Hilfsfunktion für die vorgegebenen
// Startmuster. Sie ist absichtlich unabhängig von ToggleCell.
func setAlive(world *World, row, col int) {
	if world.InBounds(row, col) {
		world.Cells[row][col] = true
	}
}
