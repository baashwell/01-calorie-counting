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
	{Scissors, Scissors, Draw},
	{Scissors, Paper, Win},
	{Scissors, Rock, Lose},
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

var get_game_points_test_set = []struct {
	player_choice   Choice
	game_result     Result
	expected_points int
}{
	{Rock, Win, 7},
	{Paper, Win, 8},
	{Scissors, Win, 9},
}

func TestGetGamePoints(t *testing.T) {
	t.Run("Test Game Points is X when player choose Y and Result was Z", func(t *testing.T) {
		for _, test := range get_game_points_test_set {
			result := GetGamePoints(test.player_choice, test.game_result)
			if test.expected_points != result {
				t.Errorf("expected %q result %q", test.expected_points, result)
			}
		}
	})
}
