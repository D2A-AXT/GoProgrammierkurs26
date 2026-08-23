package musterloesungen

// 01. Bis fünf zählen

func CountToFive() []int {
    var result []int
    for i := 1; i <= 5; i++ {
        result = append(result, i)
    }
    return result
}
