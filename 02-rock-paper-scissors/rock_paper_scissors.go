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
	return Win
}
