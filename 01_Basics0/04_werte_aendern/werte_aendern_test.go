package werte_aendern

import "testing"

func TestUpdateScore(t *testing.T) {
    wantStatus := "fertig"
    wantPoints := 50

    gotStatus, gotPoints := UpdateScore()

    if gotStatus != wantStatus || gotPoints != wantPoints {
        t.Errorf("bekommen: (%q, %d); erwartet: (%q, %d)",
            gotStatus, gotPoints, wantStatus, wantPoints)
    }
}
