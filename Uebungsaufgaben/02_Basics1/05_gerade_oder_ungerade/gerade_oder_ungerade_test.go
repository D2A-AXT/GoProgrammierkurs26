package gerade_oder_ungerade

import "testing"

func TestIsEven(t *testing.T) {
    tests := []struct {
        input int
        want  bool
    }{
        {8, true},
        {7, false},
        {0, true},
    }

    for _, tt := range tests {
        got := IsEven(tt.input)
        if got != tt.want {
            t.Errorf("IsEven(%d) = %v; erwartet %v", tt.input, got, tt.want)
        }
    }
}
