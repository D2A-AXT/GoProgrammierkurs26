package engine

// CellAt rechnet eine Mausposition in eine Rasterposition um.
// Die Funktion ist unabhängig von Game of Life und kann später z. B.
// auch für Labyrinth, Snake oder Verkehrssimulation verwendet werden.
func CellAt(
	mouseX, mouseY int,
	gridX, gridY int,
	cellSize int,
	rows, cols int,
) (row, col int, ok bool) {
	if mouseX < gridX || mouseY < gridY {
		return 0, 0, false
	}

	col = (mouseX - gridX) / cellSize
	row = (mouseY - gridY) / cellSize

	if row < 0 || row >= rows || col < 0 || col >= cols {
		return 0, 0, false
	}

	return row, col, true
}
