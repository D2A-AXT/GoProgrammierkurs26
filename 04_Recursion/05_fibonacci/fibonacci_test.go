package fibonacci

import "testing"

func TestFib(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0}, {1, 1}, {2, 1}, {5, 5}, {10, 55}, {20, 6765},
	}

	for _, test := range tests {
		if got := Fib(test.n); got != test.want {
			t.Errorf("Fib(%d) = %d, erwartet %d", test.n, got, test.want)
		}
	}
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(30)
	}
}
