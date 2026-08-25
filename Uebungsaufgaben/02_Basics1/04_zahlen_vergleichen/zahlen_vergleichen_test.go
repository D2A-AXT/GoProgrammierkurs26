package zahlen_vergleichen

import "testing"

func TestIsGreater(t *testing.T) {
    tests := []struct {
        a, b int
        want bool
    }{
        {10, 5, true},
        {5, 10, false},
        {5, 5, false},
    }

    for _, tt := range tests {
        got := IsGreater(tt.a, tt.b)
        if got != tt.want {
            t.Errorf("IsGreater(%d, %d) = %v; erwartet %v", tt.a, tt.b, got, tt.want)
        }
    }
}
