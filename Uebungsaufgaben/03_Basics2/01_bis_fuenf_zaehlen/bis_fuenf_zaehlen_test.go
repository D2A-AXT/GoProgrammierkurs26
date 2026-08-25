package bis_fuenf_zaehlen

import (
    "reflect"
    "testing"
)

func TestCountToFive(t *testing.T) {
    want := []int{1, 2, 3, 4, 5}

    got := CountToFive()

    if !reflect.DeepEqual(got, want) {
        t.Errorf("bekommen: %v; erwartet: %v", got, want)
    }
}
