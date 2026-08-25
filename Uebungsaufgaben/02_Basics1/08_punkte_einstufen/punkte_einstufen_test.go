package punkte_einstufen

import "testing"

func TestClassifyPoints(t *testing.T) {
    tests := []struct {
        points int
        want   string
    }{
        {95, "sehr gut"},
        {90, "sehr gut"},
        {76, "gut"},
        {75, "gut"},
        {50, "bestanden"},
        {49, "nicht bestanden"},
    }

    for _, tt := range tests {
        got := ClassifyPoints(tt.points)
        if got != tt.want {
            t.Errorf("ClassifyPoints(%d) = %q; erwartet %q", tt.points, got, tt.want)
        }
    }
}
