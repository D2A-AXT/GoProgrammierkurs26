package factorial

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 1}, {1, 1}, {2, 2}, {3, 6}, {5, 120}, {7, 5040},
	}

	for _, test := range tests {
		if got := Factorial(test.n); got != test.want {
			t.Errorf("Factorial(%d) = %d, erwartet %d", test.n, got, test.want)
		}
	}
}
