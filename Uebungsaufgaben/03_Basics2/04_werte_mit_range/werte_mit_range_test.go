package werte_mit_range

import (
    "reflect"
    "testing"
)

func TestCopyValues(t *testing.T) {
    input := []int{2, 4, 6, 8}
    want := []int{2, 4, 6, 8}

    got := CopyValues(input)

    if !reflect.DeepEqual(got, want) {
        t.Errorf("bekommen: %v; erwartet: %v", got, want)
    }
}
