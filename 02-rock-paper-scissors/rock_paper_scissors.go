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
	return 7
}
