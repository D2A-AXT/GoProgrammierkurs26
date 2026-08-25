package typableitung

import "testing"

func TestCourseData(t *testing.T) {
    wantCourse := "Go Grundlagen"
    wantParticipants := 12
    wantActive := true

    gotCourse, gotParticipants, gotActive := CourseData()

    if gotCourse != wantCourse || gotParticipants != wantParticipants || gotActive != wantActive {
        t.Errorf("bekommen: (%q, %d, %v); erwartet: (%q, %d, %v)",
            gotCourse, gotParticipants, gotActive,
            wantCourse, wantParticipants, wantActive)
    }
}
