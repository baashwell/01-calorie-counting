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

func GetGameResult(player_choice Choice, opponent_choice Choice) Result {
	game_results := map[Choice]map[Choice]Result{
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
	}

	return game_results[player_choice][opponent_choice]
}
