package rock_paper_scissors

import (
	"testing"
)

var get_game_result_test_set = []struct {
	player_choice   Choice
	opponent_choice Choice
	expected_result Result
}{
	{Rock, Scissors, Win},
	{Rock, Paper, Lose},
	{Rock, Rock, Draw},
	{Paper, Scissors, Lose},
	{Paper, Paper, Draw},
	{Paper, Rock, Win},
}

func TestGetGameResult(t *testing.T) {
	t.Run("Test Game Result is X when player has Y and opponent has Z", func(t *testing.T) {
		for _, test := range get_game_result_test_set {
			result := GetGameResult(test.player_choice, test.opponent_choice)
			if test.expected_result != result {
				t.Errorf("expected %q result %q", test.expected_result, result)
			}
		}
	})
}
