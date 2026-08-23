package musterloesungen

// 06. Zutritt prüfen

func CanEnter(age int, hasID bool) bool {
    return age >= 18 && hasID
}
