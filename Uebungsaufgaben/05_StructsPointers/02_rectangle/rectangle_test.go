package rectangle

import "testing"

func TestArea(t *testing.T) {
	r := Rectangle{Width: 4, Height: 3}
	if got := r.Area(); got != 12 {
		t.Errorf("Area() = %v, want 12", got)
	}
}

func TestPerimeter(t *testing.T) {
	r := Rectangle{Width: 4, Height: 3}
	if got := r.Perimeter(); got != 14 {
		t.Errorf("Perimeter() = %v, want 14", got)
	}
}

func TestScaledReturnsCopy(t *testing.T) {
	original := Rectangle{Width: 4, Height: 3}
	scaled := original.Scaled(2)

	if scaled.Width != 8 || scaled.Height != 6 {
		t.Errorf("Scaled(2) = %+v, want Width=8 Height=6", scaled)
	}
	if original.Width != 4 || original.Height != 3 {
		t.Errorf("Original wurde verändert: %+v", original)
	}
}
