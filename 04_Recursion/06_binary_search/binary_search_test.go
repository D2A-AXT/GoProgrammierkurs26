package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	numbers := []int{2, 4, 7, 10, 15, 20, 31}
	tests := []struct {
		target int
		want   int
	}{
		{2, 0}, {7, 2}, {10, 3}, {20, 5}, {31, 6}, {5, -1}, {100, -1},
	}

	for _, test := range tests {
		got := BinarySearch(numbers, test.target)
		if got != test.want {
			t.Errorf("BinarySearch(%v, %d) = %d, erwartet %d", numbers, test.target, got, test.want)
		}
	}
}
