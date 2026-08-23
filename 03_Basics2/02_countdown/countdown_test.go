package countdown

import (
    "reflect"
    "testing"
)

func TestCountdown(t *testing.T) {
    input := 5
    want := []int{5, 4, 3, 2, 1}

    got := Countdown(input)

    if !reflect.DeepEqual(got, want) {
        t.Errorf("Countdown(%d) = %v; erwartet %v", input, got, want)
    }
}
