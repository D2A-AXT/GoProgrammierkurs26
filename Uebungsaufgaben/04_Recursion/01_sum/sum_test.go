package sum

import "testing"

func TestSumTo(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0}, {1, 1}, {2, 3}, {5, 15}, {10, 55},
	}

	for _, test := range tests {
		got := SumTo(test.n)
		if got != test.want {
			t.Errorf("SumTo(%d) = %d, erwartet %d", test.n, got, test.want)
		}
	}
}
