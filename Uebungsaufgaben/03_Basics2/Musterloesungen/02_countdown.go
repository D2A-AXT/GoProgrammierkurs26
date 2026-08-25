package musterloesungen

// 02. Countdown

func Countdown(start int) []int {
    var result []int
    for i := start; i >= 1; i-- {
        result = append(result, i)
    }
    return result
}
