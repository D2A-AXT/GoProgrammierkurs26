package slice_ausschneiden

import (
    "reflect"
    "testing"
)

func TestMiddleSection(t *testing.T) {
    input := []int{10, 20, 30, 40, 50}
    want := []int{20, 30, 40}

    got := MiddleSection(input)

    if !reflect.DeepEqual(got, want) {
        t.Errorf("bekommen: %v; erwartet: %v", got, want)
    }
}
