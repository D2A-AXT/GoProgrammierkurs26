package slicesum

import "testing"

func TestSumSlice(t *testing.T) {
	tests := []struct {
		numbers []int
		want    int
	}{
		{[]int{}, 0},
		{[]int{5}, 5},
		{[]int{1, 2, 3}, 6},
		{[]int{10, -5, 3}, 8},
		{[]int{2, 4, 6, 8}, 20},
	}

	for _, test := range tests {
		got := SumSlice(test.numbers)
		if got != test.want {
			t.Errorf("SumSlice(%v) = %d, erwartet %d", test.numbers, got, test.want)
		}
	}
}
