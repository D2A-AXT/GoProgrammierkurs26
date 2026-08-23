package dereference

import "testing"

func TestRead(t *testing.T) {
	x := 7
	if got := Read(&x); got != 7 {
		t.Errorf("Read(&x) = %d, want 7", got)
	}
}

func TestWrite(t *testing.T) {
	x := 7
	Write(&x, 99)
	if x != 99 {
		t.Errorf("x = %d after Write, want 99", x)
	}
}

func TestAdd(t *testing.T) {
	x := 10
	Add(&x, 5)
	if x != 15 {
		t.Errorf("x = %d after Add, want 15", x)
	}
}
