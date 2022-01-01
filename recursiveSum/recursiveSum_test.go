package recursiveSum_test

import (
	"testing"

	"github.com/SPECTR3R/go-exercises/recursiveSum"
)

// TestSum can run especific test by running: go test -v -run <regex>
func TestInts(t *testing.T) {
	tt := []struct {
		name  string
		input []int
		want  int
	}{
		{"no nums", nil, 0},
		{"one and minus one", []int{1, -1}, 0},
		{"one to five", []int{1, 2, 3, 4, 5}, 15},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := recursiveSum.Ints(tc.input...)
			if s != tc.want {
				t.Errorf("got %d, want %d", s, tc.want)
			}
		})
	}
}

// This is a way of hardcofing test
// func TestInts(t *testing.T) {
// 	s := Ints(1, 2, 3, 4, 5)
// 	if s != 15 {
// 		t.Errorf("Sum of 1 to 5 should be 15, got %d", s)
// 	}
// 	s = Ints(1, -1)
// 	if s != 0 {
// 		t.Errorf("Sum of 1 and -1 0, got %d", s)
// 	}
// 	s = Ints()
// 	if s != 0 {
// 		t.Errorf("Sum of no numbers 0, got %d", s)
// 	}
// }
