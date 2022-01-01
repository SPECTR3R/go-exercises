package palindromeChecker_test

import (
	"testing"

	"github.com/SPECTR3R/go-exercises/palindromeChecker"
)

func TestIsAPalindrome(t *testing.T) {
	tt := []struct {
		name  string
		input string
		want  bool
	}{
		{"no string is not a Palindrome", "", false},
		{"'mom' is a Palindrome", "mom", true},
		{"'bill' is not a Palindrome", "bill", false},
		{"'Mom' is a Palindrome", "Mom", true},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := palindromeChecker.IsAPalindrome(tc.input)
			if s != tc.want {
				t.Errorf("it should be able to tell that %v, got %t expect %t", tc.name, s, tc.want)
			}
		})
	}
}
