package rock_paper_scissors

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
	return Rock
}
