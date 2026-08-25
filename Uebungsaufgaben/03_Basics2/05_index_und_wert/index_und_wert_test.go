package index_und_wert

import (
    "reflect"
    "testing"
)

func TestLabelValues(t *testing.T) {
    input := []string{"Apfel", "Banane", "Orange"}
    want := []string{"0: Apfel", "1: Banane", "2: Orange"}

    got := LabelValues(input)

    if !reflect.DeepEqual(got, want) {
        t.Errorf("bekommen: %v; erwartet: %v", got, want)
    }
}
