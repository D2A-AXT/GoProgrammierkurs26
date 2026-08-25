package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		text string
		want bool
	}{
		{"", true}, {"a", true}, {"anna", true}, {"otto", true},
		{"racecar", true}, {"hello", false}, {"recursive", false},
	}

	for _, test := range tests {
		got := IsPalindrome(test.text)
		if got != test.want {
			t.Errorf("IsPalindrome(%q) = %v, erwartet %v", test.text, got, test.want)
		}
	}
}
