package linkedlist

import "testing"

func buildList() *Node {
	return &Node{
		Value: 3,
		Next: &Node{
			Value: 8,
			Next:  &Node{Value: 5},
		},
	}
}

func TestLength(t *testing.T) {
	if got := Length(buildList()); got != 3 {
		t.Errorf("Length = %d, want 3", got)
	}
	if got := Length(nil); got != 0 {
		t.Errorf("Length(nil) = %d, want 0", got)
	}
}

func TestSum(t *testing.T) {
	if got := Sum(buildList()); got != 16 {
		t.Errorf("Sum = %d, want 16", got)
	}
	if got := Sum(nil); got != 0 {
		t.Errorf("Sum(nil) = %d, want 0", got)
	}
}

func TestFind(t *testing.T) {
	head := buildList()
	got := Find(head, 8)
	if got == nil || got.Value != 8 {
		t.Fatalf("Find(8) = %v, want Node with Value 8", got)
	}
	if got != head.Next {
		t.Error("Find soll einen Pointer auf den vorhandenen Node zurückgeben")
	}
	if Find(head, 999) != nil {
		t.Error("Find für unbekannten Wert sollte nil liefern")
	}
}

func TestAppend(t *testing.T) {
	head := buildList()
	returned := Append(head, 13)

	if returned != head {
		t.Error("Append auf nicht-leerer Liste sollte denselben head zurückgeben")
	}
	if Length(head) != 4 {
		t.Fatalf("Length after Append = %d, want 4", Length(head))
	}
	last := Find(head, 13)
	if last == nil || last.Value != 13 || last.Next != nil {
		t.Errorf("Neuer letzter Node ist falsch: %+v", last)
	}
}

func TestAppendToEmptyList(t *testing.T) {
	head := Append(nil, 42)
	if head == nil {
		t.Fatal("Append(nil, 42) returned nil")
	}
	if head.Value != 42 || head.Next != nil {
		t.Errorf("head = %+v, want single Node with Value 42", head)
	}
}
