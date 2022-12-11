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
	total := 0

	for _, character := range slice {
		total += GetPriorityValueOfCharacter(character)
	}

	return total
}

func GetPriorityOfDuplicatedCharactersInRucksackCompartments(rucksack []rune) int {
	compartments := GetRucksackCompartments(rucksack)
	duplicate_items := GetCharactersExistingInBothSlices(compartments[0], compartments[1])
	return GetTotalPriorityOfSlice(duplicate_items)
}

func GetRucksackCompartments(rucksack []rune) [][]rune {
	compartments := make([][]rune, 2)
	compartment_size := len(rucksack) / 2

	for i, character := range rucksack {
		if i >= compartment_size {
			compartments[1] = append(compartments[1], character)
		} else {
			compartments[0] = append(compartments[0], character)
		}
	}

	return compartments
}
