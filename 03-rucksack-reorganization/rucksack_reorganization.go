package rucksack_reorganization

func GetPriorityValueOfCharacter(character rune) int {
	differential := 38

	if character >= 97 {
		differential = 96
	}

	return int(character) - differential
}
