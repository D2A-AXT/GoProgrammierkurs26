package mergesort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{}, []int{}},
		{[]int{5}, []int{5}},
		{[]int{3, 1, 2}, []int{1, 2, 3}},
		{[]int{8, 3, 6, 2, 7, 1}, []int{1, 2, 3, 6, 7, 8}},
		{[]int{5, 2, 5, 1, 2}, []int{1, 2, 2, 5, 5}},
	}

	for _, test := range tests {
		got := MergeSort(test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("MergeSort(%v) = %v, erwartet %v", test.input, got, test.want)
		}
	}
}
