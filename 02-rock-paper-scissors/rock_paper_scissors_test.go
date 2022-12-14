package rock_paper_scissors

import (
	"strconv"
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
	{Rock, Draw, 4},
	{Paper, Draw, 5},
	{Scissors, Draw, 6},
	{Rock, Lose, 1},
	{Paper, Lose, 2},
	{Scissors, Lose, 3},
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

var get_choice_from_string_test_set = []struct {
	string_input    string
	expected_choice Choice
}{
	{"A", Rock},
	{"B", Paper},
	{"C", Scissors},
	{"X", Rock},
	{"Y", Paper},
	{"Z", Scissors},
}

func TestGetChoiceFromString(t *testing.T) {
	t.Run("Test Get Choice is X when string is Y", func(t *testing.T) {
		for _, test := range get_choice_from_string_test_set {
			result := GetChoiceFromString(test.string_input)
			if test.expected_choice != result {
				t.Errorf("expected %q result %q", test.expected_choice, result)
			}
		}
	})
}

var get_total_points_for_player_test_set = []struct {
	string_input    string
	expected_points int
}{
	{"A Y", 8},
	{
		`A Y
B X
C Z`,
		15,
	},
}

func TestGetTotalPointsWhenWeGivenChoices(t *testing.T) {
	t.Run("Test Get total points is X when string is Y", func(t *testing.T) {
		for _, test := range get_total_points_for_player_test_set {
			result := GetTotalPointsWhenWeGivenChoices(test.string_input)
			if test.expected_points != result {
				t.Errorf("expected %q result %q", strconv.Itoa(test.expected_points), strconv.Itoa(result))
			}
		}
	})
}

func TestPuzzleInput(t *testing.T) {
	t.Run("Test Get total points for puzzle input", func(t *testing.T) {
		expected := 12794
		result := GetTotalPointsWhenWeGivenChoices(puzzle_input)
		if expected != result {
			t.Errorf("expected %q result %q", strconv.Itoa(expected), strconv.Itoa(result))
		}

	})
}

var get_choice_to_make_test_set = []struct {
	opponents_choice Choice
	result           Result
	expected_choice  Choice
}{
	{Rock, Win, Paper},
	{Rock, Lose, Scissors},
	{Rock, Draw, Rock},
	{Paper, Win, Scissors},
	{Paper, Lose, Rock},
	{Paper, Draw, Paper},
	{Scissors, Win, Rock},
	{Scissors, Lose, Paper},
	{Scissors, Draw, Scissors},
}

func TestGetChoiceToMakeResult(t *testing.T) {
	t.Run("Test Get Choice x When opponents choice is Y and result is Z", func(t *testing.T) {
		for _, test := range get_choice_to_make_test_set {
			result := GetChoiceToMake(test.opponents_choice, test.result)
			if test.expected_choice != result {
				t.Errorf("expected %q result %q", test.expected_choice, result)
			}
		}
	})
}

var get_total_points_for_player_with_result_test_set = []struct {
	string_input    string
	expected_points int
}{
	{"A Y", 4},
	{
		`A Y
B X
C Z`,
		12,
	},
}

func TestGetTotalPointsWhenGivenOpponentChoiceAndResult(t *testing.T) {
	t.Run("Test Get total points is X when string is Y", func(t *testing.T) {
		for _, test := range get_total_points_for_player_with_result_test_set {
			result := GetTotalPointsWhenGivenOpponentChoiceAndResult(test.string_input)
			if test.expected_points != result {
				t.Errorf("expected %q result %q", strconv.Itoa(test.expected_points), strconv.Itoa(result))
			}
		}
	})
}

func TestPuzzleInput2(t *testing.T) {
	t.Run("Test Get total points for puzzle input 2", func(t *testing.T) {
		expected := 14979
		result := GetTotalPointsWhenGivenOpponentChoiceAndResult(puzzle_input)
		if expected != result {
			t.Errorf("expected %q result %q", strconv.Itoa(expected), strconv.Itoa(result))
		}

	})
}
