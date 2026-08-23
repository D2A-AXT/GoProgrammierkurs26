package musterloesungen

// 04. Werte mit range durchlaufen

func CopyValues(numbers []int) []int {
    var result []int
    for _, value := range numbers {
        result = append(result, value)
    }
    return result
}
