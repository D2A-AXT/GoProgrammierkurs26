package summe_bilden

import "testing"

func TestSumTo(t *testing.T) {
    tests := []struct {
        input int
        want  int
    }{
        {5, 15},
        {10, 55},
        {1, 1},
    }

    for _, tt := range tests {
        got := SumTo(tt.input)
        if got != tt.want {
            t.Errorf("SumTo(%d) = %d; erwartet %d", tt.input, got, tt.want)
        }
    }
}
