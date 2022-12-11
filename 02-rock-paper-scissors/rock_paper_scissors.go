package rock_paper_scissors

import (
	"strings"
)

type Choice int64

const (
	Rock Choice = iota
	Paper
	Scissors
)

type Result int64

const (
	Win Result = iota
	Draw
	Lose
)

var game_results = map[Choice]map[Choice]Result{
	Rock: {
		Rock:     Draw,
		Paper:    Lose,
		Scissors: Win,
	},
	Paper: {
		Rock:     Win,
		Paper:    Draw,
		Scissors: Lose,
	},
	Scissors: {
		Rock:     Lose,
		Paper:    Win,
		Scissors: Draw,
	},
}

func GetGameResult(player_choice Choice, opponent_choice Choice) Result {
	return game_results[player_choice][opponent_choice]
}

func GetGamePoints(player_choice Choice, game_result Result) int {
	return GetResultPoints(game_result) + GetChoicePoints(player_choice)
}

func GetResultPoints(game_result Result) int {
	switch game_result {
	case Win:
		return 6
	case Lose:
		return 0
	default:
		return 3
	}
}

func GetChoicePoints(player_choice Choice) int {
	switch player_choice {
	case Rock:
		return 1
	case Paper:
		return 2
	default:
		return 3
	}
}

func GetChoiceFromString(input string) Choice {
	switch input {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	default:
		return Scissors
	}
}

func GetResultFromString(input string) Result {
	switch input {
	case "X":
		return Lose
	case "Y":
		return Draw
	default:
		return Win
	}
}

func GetTotalPointsWhenWeGivenChoices(input string) int {
	total_points := 0
	rounds := strings.Split(input, "\n")

	for _, round := range rounds {
		choices := strings.Split(round, " ")
		opponent_choice := GetChoiceFromString(choices[0])
		player_choice := GetChoiceFromString(choices[1])

		result := GetGameResult(player_choice, opponent_choice)
		total_points += GetGamePoints(player_choice, result)
	}

	return total_points
}

func GetTotalPointsWhenGivenOpponentChoiceAndResult(input string) int {
	total_points := 0
	rounds := strings.Split(input, "\n")

	for _, round := range rounds {
		choices := strings.Split(round, " ")
		opponent_choice := GetChoiceFromString(choices[0])
		result := GetResultFromString(choices[1])
		player_choice := GetChoiceToMake(opponent_choice, result)

		total_points += GetGamePoints(player_choice, result)
	}

	return total_points
}

func GetChoiceToMake(opponent_choice Choice, result Result) Choice {
	choices := game_results[opponent_choice]
	opposite_result := GetOppositeOfResult(result)

	for key, element := range choices {
		if element == opposite_result {
			return key
		}
	}

	return opponent_choice
}

func GetOppositeOfResult(result Result) Result {
	switch result {
	case Win:
		return Lose
	case Lose:
		return Win
	default:
		return Draw
	}
}
