package rucksack_reorganization

func GetPriorityValueOfCharacter(character rune) int {
	differential := 38

	if character >= 97 {
		differential = 96
	}

	return int(character) - differential
}

func GetCharactersExistingInBothSlices(slice_1 []rune, slice_2 []rune) []rune {
	return []rune{'a'}
}
