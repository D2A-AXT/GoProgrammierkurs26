package pointerreceivers

import "testing"

func TestAddedDoesNotChangeOriginal(t *testing.T) {
	original := Counter{Value: 10}
	copy := original.Added(5)

	if copy.Value != 15 {
		t.Errorf("copy.Value = %d, want 15", copy.Value)
	}
	if original.Value != 10 {
		t.Errorf("original.Value = %d, want unverändert 10", original.Value)
	}
}

func TestAddChangesOriginal(t *testing.T) {
	c := Counter{Value: 10}
	c.Add(5) // Go nimmt automatisch &c, weil c adressierbar ist.
	if c.Value != 15 {
		t.Errorf("c.Value = %d, want 15", c.Value)
	}
}

func TestReset(t *testing.T) {
	c := Counter{Value: 123}
	c.Reset()
	if c.Value != 0 {
		t.Errorf("c.Value = %d, want 0", c.Value)
	}
}
