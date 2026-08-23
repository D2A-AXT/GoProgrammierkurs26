package addresses

import "testing"

func TestPointerTo(t *testing.T) {
	p := PointerTo(42)
	if p == nil {
		t.Fatal("PointerTo(42) returned nil")
	}
	if *p != 42 {
		t.Errorf("*p = %d, want 42", *p)
	}
}

func TestSameAddress(t *testing.T) {
	x := 10
	a := &x
	b := &x
	y := 10
	c := &y

	if !SameAddress(a, b) {
		t.Error("Pointer auf dieselbe Variable sollten dieselbe Adresse haben")
	}
	if SameAddress(a, c) {
		t.Error("Pointer auf verschiedene Variablen sollten verschiedene Adressen haben")
	}
}
