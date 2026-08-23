package engine

import "testing"

func TestCellAt(t *testing.T) {
	row, col, ok := CellAt(25, 35, 10, 20, 10, 5, 5)
	if !ok || row != 1 || col != 1 {
		t.Fatalf("CellAt = (%d,%d,%v), erwartet (1,1,true)", row, col, ok)
	}
}
