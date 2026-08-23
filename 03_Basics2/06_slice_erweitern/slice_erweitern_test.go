package slice_erweitern

import (
    "reflect"
    "testing"
)

func TestExtendNames(t *testing.T) {
    input := []string{"Alice", "Bob"}
    first := "Charlie"
    second := "Dora"
    want := []string{"Alice", "Bob", "Charlie", "Dora"}

    got := ExtendNames(input, first, second)

    if !reflect.DeepEqual(got, want) {
        t.Errorf("bekommen: %v; erwartet: %v", got, want)
    }
}
