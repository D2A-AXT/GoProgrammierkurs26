package structpointers

import "testing"

func TestBirthday(t *testing.T) {
	s := Student{Name: "Ada", Age: 19}
	Birthday(&s)
	if s.Age != 20 {
		t.Errorf("Age = %d, want 20", s.Age)
	}
}

func TestRename(t *testing.T) {
	s := Student{Name: "Ada", Age: 19}
	Rename(&s, "Grace")
	if s.Name != "Grace" {
		t.Errorf("Name = %q, want Grace", s.Name)
	}
}

func TestOlder(t *testing.T) {
	a := Student{Name: "A", Age: 20}
	b := Student{Name: "B", Age: 25}

	if got := Older(&a, &b); got != &b {
		t.Error("Older sollte einen Pointer auf b zurückgeben")
	}

	b.Age = 20
	if got := Older(&a, &b); got != &a {
		t.Error("Bei gleichem Alter sollte Older einen Pointer auf a zurückgeben")
	}
}
