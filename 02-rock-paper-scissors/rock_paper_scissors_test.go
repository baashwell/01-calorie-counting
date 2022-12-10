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
	t.Run("Test Does Player Lose when player has Rock and opponent has Paper", func(t *testing.T) {
		expected := Lose
		result := GetGameResult(Rock, Paper)

		if expected != result {
			t.Errorf("expected %q result %q", expected, result)
		}
	})
	t.Run("Test Does Player Draw when player has Rock and opponent has Rock", func(t *testing.T) {
		expected := Draw
		result := GetGameResult(Rock, Rock)

		if expected != result {
			t.Errorf("expected %q result %q", expected, result)
		}
	})
}
