package persoenliche_begruessung

import "testing"

func TestGreeting(t *testing.T) {
    input := "Alice"
    want := "Hallo, Alice!"

    got := Greeting(input)

    if got != want {
        t.Errorf("Greeting(%q) = %q; erwartet %q", input, got, want)
    }
}
