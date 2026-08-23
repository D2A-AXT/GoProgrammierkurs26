package variablen_kennenlernen

import "testing"

func TestStudentProfile(t *testing.T) {
    wantName := "Alice"
    wantAge := 20
    wantHeight := 1.72
    wantStudent := true

    gotName, gotAge, gotHeight, gotStudent := StudentProfile()

    if gotName != wantName || gotAge != wantAge || gotHeight != wantHeight || gotStudent != wantStudent {
        t.Errorf(
            "bekommen: (%q, %d, %.2f, %v); erwartet: (%q, %d, %.2f, %v)",
            gotName, gotAge, gotHeight, gotStudent,
            wantName, wantAge, wantHeight, wantStudent,
        )
    }
}
