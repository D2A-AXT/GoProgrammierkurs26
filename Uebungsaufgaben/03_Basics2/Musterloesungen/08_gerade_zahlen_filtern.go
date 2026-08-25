package musterloesungen

// 08. Gerade Zahlen filtern

func FilterEven(numbers []int) []int {
    var evenNumbers []int
    for _, value := range numbers {
        if value%2 == 0 {
            evenNumbers = append(evenNumbers, value)
        }
    }
    return evenNumbers
}
