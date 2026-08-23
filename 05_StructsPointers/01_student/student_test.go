package student

import "testing"

func TestNewStudent(t *testing.T) {
	s := NewStudent("Anna", 19, 1)

	if s.Name != "Anna" {
		t.Errorf("Name = %q, want %q", s.Name, "Anna")
	}
	if s.Age != 19 {
		t.Errorf("Age = %d, want 19", s.Age)
	}
	if s.Semester != 1 {
		t.Errorf("Semester = %d, want 1", s.Semester)
	}
}

func TestIsFirstSemester(t *testing.T) {
	if !IsFirstSemester(Student{Name: "Anna", Semester: 1}) {
		t.Error("Student im 1. Semester sollte true liefern")
	}
	if IsFirstSemester(Student{Name: "Ben", Semester: 3}) {
		t.Error("Student im 3. Semester sollte false liefern")
	}
}
