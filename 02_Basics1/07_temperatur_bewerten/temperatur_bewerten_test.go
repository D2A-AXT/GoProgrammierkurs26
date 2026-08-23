package temperatur_bewerten

import "testing"

func TestTemperatureLabel(t *testing.T) {
    tests := []struct {
        temp int
        want string
    }{
        {30, "warm"},
        {25, "warm"},
        {24, "kühl"},
    }

    for _, tt := range tests {
        got := TemperatureLabel(tt.temp)
        if got != tt.want {
            t.Errorf("TemperatureLabel(%d) = %q; erwartet %q", tt.temp, got, tt.want)
        }
    }
}
