package gerade_zahlen_filtern

import (
    "reflect"
    "testing"
)

func TestFilterEven(t *testing.T) {
    tests := []struct {
        input []int
        want  []int
    }{
        {[]int{3, 8, 11, 14, 20, 25, 30}, []int{8, 14, 20, 30}},
        {[]int{1, 3, 5}, []int{}},
        {[]int{2, 4, 6}, []int{2, 4, 6}},
    }

    for _, tt := range tests {
        got := FilterEven(tt.input)
        if !reflect.DeepEqual(got, tt.want) {
            t.Errorf("FilterEven(%v) = %v; erwartet %v", tt.input, got, tt.want)
        }
    }
}
