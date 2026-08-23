package student

// Übung 1 – Mein erstes Struct
// Schwierigkeit: ★☆☆☆☆
//
// Lernziele:
//   - Structs lesen und erzeugen
//   - Felder über den Punktoperator verwenden
//   - Einen Struct-Wert aus einer Funktion zurückgeben

// Student beschreibt einen Studenten.
// Die Felder sind für den Einstieg bereits vorgegeben.
type Student struct {
	Name     string
	Age      int
	Semester int
}

// NewStudent soll einen Student-Wert mit den übergebenen Werten erzeugen.
//
// Tipp:
//
//	return Student{
//	    Name: ...,
//	    Age: ...,
//	    Semester: ...,
//	}
func NewStudent(name string, age int, semester int) Student {
	// TODO: Erzeuge und returne den Student.
	return Student{}
}

// IsFirstSemester soll true liefern, wenn der Student im 1. Semester ist.
func IsFirstSemester(s Student) bool {
	// TODO: Greife mit s.Semester auf das Feld zu.
	return false
}
