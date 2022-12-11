package rucksack_reorganization

func GetPriorityValueOfCharacter(character rune) int {
	differential := 38

	if character >= 97 {
		differential = 96
	}

	return int(character) - differential
}

func GetCharactersExistingInBothSlices(slice_1 []rune, slice_2 []rune) []rune {
	duplicate_characters := []rune{}

	for _, slice_1_char := range slice_1 {
		for _, slice_2_char := range slice_2 {
			if slice_1_char == slice_2_char {
				duplicate_characters = append(duplicate_characters, slice_1_char)
			}
		}
	}
	return duplicate_characters
}

func GetTotalPriorityOfSlice(slice []rune) int {
	return 1
}
