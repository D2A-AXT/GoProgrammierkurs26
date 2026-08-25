package tree

import "testing"

func TestSumTree(t *testing.T) {
	tree := &Node{
		Value: 10,
		Children: []*Node{
			{Value: 5, Children: []*Node{{Value: 2}, {Value: 3}}},
			{Value: 8},
		},
	}

	if got := SumTree(tree); got != 28 {
		t.Errorf("SumTree(tree) = %d, erwartet 28", got)
	}
}

func TestSumTreeNil(t *testing.T) {
	if got := SumTree(nil); got != 0 {
		t.Errorf("SumTree(nil) = %d, erwartet 0", got)
	}
}
