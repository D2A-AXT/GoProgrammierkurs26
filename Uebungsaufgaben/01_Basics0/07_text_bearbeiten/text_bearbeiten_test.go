package text_bearbeiten

import "testing"

func TestChangeCase(t *testing.T) {
    input := "Go bei ABB"
    wantUpper := "GO BEI ABB"
    wantLower := "go bei abb"

    gotUpper, gotLower := ChangeCase(input)

    if gotUpper != wantUpper || gotLower != wantLower {
        t.Errorf("bekommen: (%q, %q); erwartet: (%q, %q)",
            gotUpper, gotLower, wantUpper, wantLower)
    }
}
