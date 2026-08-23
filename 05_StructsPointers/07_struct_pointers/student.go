package structpointers

// Übung 7 – Pointer auf Structs
// Schwierigkeit: ★★★☆☆
//
// Lernziele:
//   - Structs über Pointer verändern
//   - Go vereinfacht den Feldzugriff automatisch
//   - student.Age funktioniert auch bei *Student

type Student struct {
	Name string
	Age  int
}

// Birthday soll das Alter des ORIGINALEN Studenten um 1 erhöhen.
func Birthday(student *Student) {
	// TODO
	// Du könntest schreiben: (*student).Age++
	// Idiomatischer ist in Go aber einfach: student.Age++
}

// Rename soll den Namen des ORIGINALEN Studenten ändern.
func Rename(student *Student, newName string) {
	// TODO
}

// Older soll einen Pointer auf den älteren der beiden Studenten zurückgeben.
// Bei gleichem Alter soll a zurückgegeben werden.
func Older(a *Student, b *Student) *Student {
	// TODO
	return nil
}
