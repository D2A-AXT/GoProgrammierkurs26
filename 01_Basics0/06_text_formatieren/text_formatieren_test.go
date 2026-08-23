package text_formatieren

import "testing"

func TestFormatCourse(t *testing.T) {
    course := "Go Grundlagen"
    participants := 12
    want := "Go Grundlagen hat 12 Teilnehmer."

    got := FormatCourse(course, participants)

    if got != want {
        t.Errorf("bekommen: %q; erwartet: %q", got, want)
    }
}
