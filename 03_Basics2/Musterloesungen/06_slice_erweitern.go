package musterloesungen

// 06. Slice erweitern

func ExtendNames(names []string, first string, second string) []string {
    return append(names, first, second)
}
