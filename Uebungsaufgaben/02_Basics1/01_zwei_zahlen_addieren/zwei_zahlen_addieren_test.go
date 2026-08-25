package zwei_zahlen_addieren

import "testing"

func TestAdd(t *testing.T) {
    a := 5
    b := 3
    want := 8

    got := Add(a, b)

    if got != want {
        t.Errorf("Add(%d, %d) = %d; erwartet %d", a, b, got, want)
    }
}
