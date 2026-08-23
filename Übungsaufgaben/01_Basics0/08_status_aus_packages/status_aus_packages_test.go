package status_aus_packages

import "testing"

func TestBuildStatus(t *testing.T) {
    name := "go kurs"
    number := 8
    want := "GO KURS - Aufgabe 8"

    got := BuildStatus(name, number)

    if got != want {
        t.Errorf("bekommen: %q; erwartet: %q", got, want)
    }
}
