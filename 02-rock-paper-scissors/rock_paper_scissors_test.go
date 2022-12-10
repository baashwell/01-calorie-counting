package rock_paper_scissors

import (
	"testing"
)

func TestGetGameResult(t *testing.T) {
	t.Run("Test Does Player Win when player has Rock and opponent has Scissors", func(t *testing.T) {
		expected := Win
		result := GetGameResult(Rock, Scissors)

		if expected != result {
			t.Errorf("expected %q result %q", expected, result)
		}
	})
}
