package musterloesungen

// 03. Summe bilden

func SumTo(n int) int {
    sum := 0
    for i := 1; i <= n; i++ {
        sum = sum + i
    }
    return sum
}
