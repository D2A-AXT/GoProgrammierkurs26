package zutritt_pruefen

import "testing"

func TestCanEnter(t *testing.T) {
    tests := []struct {
        age   int
        hasID bool
        want  bool
    }{
        {20, true, true},
        {20, false, false},
        {17, true, false},
    }

    for _, tt := range tests {
        got := CanEnter(tt.age, tt.hasID)
        if got != tt.want {
            t.Errorf("CanEnter(%d, %v) = %v; erwartet %v",
                tt.age, tt.hasID, got, tt.want)
        }
    }
}
