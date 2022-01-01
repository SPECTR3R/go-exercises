package mascot_test

import (
	"testing"

	"github.com/SPECTR3R/go-exercises/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Rolo" {
		t.Error("Expected Rolo, got ", mascot.BestMascot())
	}
}
