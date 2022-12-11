package rucksack_reorganization

import "strings"

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
		if DoesCharacterExistInSlice(slice_1_char, slice_2) {
			if !DoesCharacterExistInSlice(slice_1_char, duplicate_characters) {
				duplicate_characters = append(duplicate_characters, slice_1_char)
			}
		}
	}
	return duplicate_characters
}

func DoesCharacterExistInSlice(character rune, slice []rune) bool {
	for _, slice_char := range slice {
		if character == slice_char {
			return true
		}
	}

	return false
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

func GetPriorityOfMultipleRucksacks(input string) int {
	rucksacks := strings.Split(input, "\n")
	total_priority := 0

	for _, rucksack := range rucksacks {
		rucksack_as_runes := make([]rune, 0)

		for _, character := range rucksack {
			rucksack_as_runes = append(rucksack_as_runes, character)
		}

		total_priority += GetPriorityOfDuplicatedCharactersInRucksackCompartments(rucksack_as_runes)
	}

	return total_priority
}
