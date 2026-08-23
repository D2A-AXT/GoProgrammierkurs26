package memory

import (
	"testing"
	"unsafe"
)

func TestGetSizes(t *testing.T) {
	got := GetSizes()

	if got.BoolSize != unsafe.Sizeof(bool(false)) {
		t.Errorf("BoolSize = %d, want %d", got.BoolSize, unsafe.Sizeof(bool(false)))
	}
	if got.Int64Size != unsafe.Sizeof(int64(0)) {
		t.Errorf("Int64Size = %d, want %d", got.Int64Size, unsafe.Sizeof(int64(0)))
	}
	if got.BadLayoutSize != unsafe.Sizeof(BadLayout{}) {
		t.Errorf("BadLayoutSize = %d, want %d", got.BadLayoutSize, unsafe.Sizeof(BadLayout{}))
	}
	if got.BetterLayoutSize != unsafe.Sizeof(BetterLayout{}) {
		t.Errorf("BetterLayoutSize = %d, want %d", got.BetterLayoutSize, unsafe.Sizeof(BetterLayout{}))
	}
}

func TestMakeLabel(t *testing.T) {
	if got := MakeLabel("Ada"); got != "Student: Ada" {
		t.Errorf("MakeLabel(Ada) = %q, want %q", got, "Student: Ada")
	}
}
