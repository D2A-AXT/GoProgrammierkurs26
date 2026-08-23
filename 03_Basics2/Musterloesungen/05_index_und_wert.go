package musterloesungen

// 05. Index und Wert verwenden

import "fmt"

func LabelValues(values []string) []string {
    var result []string
    for index, value := range values {
        result = append(result, fmt.Sprintf("%d: %s", index, value))
    }
    return result
}
