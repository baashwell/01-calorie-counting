package rucksack_reorganization

import (
	"strconv"
	"testing"
)

var character_priority_test_set = []struct {
	character         rune
	expected_priority int
}{
	{'a', 1},
	{'b', 2},
	{'c', 3},
	{'d', 4},
	{'m', 13},
	{'r', 18},
	{'t', 20},
	{'v', 22},
	{'z', 26},
	{'A', 27},
	{'B', 28},
	{'C', 29},
	{'D', 30},
	{'M', 39},
	{'R', 44},
	{'T', 46},
	{'V', 48},
	{'Z', 52},
}

func TestGetPriorityValueOfCharacter(t *testing.T) {
	t.Run("Test Priority is X when character is Y", func(t *testing.T) {
		for _, test := range character_priority_test_set {
			result := GetPriorityValueOfCharacter(test.character)
			if test.expected_priority != result {
				t.Errorf("expected %q result %q", strconv.Itoa(test.expected_priority), strconv.Itoa(result))
			}
		}
	})
}

var characters_in_both_slices_test_set = []struct {
	slice_1 []rune
	slice_2 []rune
	result  []rune
}{
	{[]rune{'a'}, []rune{'a'}, []rune{'a'}},
	{[]rune{'a', 'b', 'c'}, []rune{'d', 'e', 'f'}, []rune{}},
	{[]rune{'a', 'b', 'c'}, []rune{'a', 'b', 'f'}, []rune{'a', 'b'}},
	{[]rune{'d', 'e', 'f'}, []rune{'d', 'e', 'f'}, []rune{'d', 'e', 'f'}},
}

func TestGetCharactersExistingInBothSlices(t *testing.T) {
	t.Run("Test characters are X when slices are Y and Z", func(t *testing.T) {
		for _, test := range characters_in_both_slices_test_set {
			result := GetCharactersExistingInBothSlices(test.slice_1, test.slice_2)
			for i, character := range result {
				if character != test.result[i] {
					t.Errorf("expected %q result %q", character, result[i])
				}
			}
		}
	})
}

var total_priority_of_slice_test_set = []struct {
	slice_1           []rune
	expected_priority int
}{
	{[]rune{'a'}, 1},
	{[]rune{'a', 'f', 't'}, 27},
	{[]rune{'A', 'F', 'T'}, 105},
	{[]rune{'a', 'B', 'c', 'D'}, 62},
	{[]rune{'a', 'b', 'C', 'D'}, 62},
}

func TestGetTotalPriorityOfSlice(t *testing.T) {
	t.Run("Test total priority of slice is X when slice is Y", func(t *testing.T) {
		for _, test := range total_priority_of_slice_test_set {
			result := GetTotalPriorityOfSlice(test.slice_1)
			if test.expected_priority != result {
				t.Errorf("expected %q result %q", strconv.Itoa(test.expected_priority), strconv.Itoa(result))
			}
		}
	})
}

var get_priority_of_duplicated_characters_test_set = []struct {
	rucksack          []rune
	expected_priority int
}{
	{[]rune{'a', 'a'}, 1},
	{[]rune{'a', 'b', 'c', 'd'}, 0},
	{[]rune{'a', 'b', 'c', 'a'}, 1},
	{[]rune{'F', 'G', 'F', 'G'}, 65},
}

func TestGetPriorityOfDuplicatedCharactersInRucksackCompartments(t *testing.T) {
	t.Run("Test total priority of duplicated items in rucksack compartments is X when rucksack is Y", func(t *testing.T) {
		for _, test := range get_priority_of_duplicated_characters_test_set {
			result := GetPriorityOfDuplicatedCharactersInRucksackCompartments(test.rucksack)
			if test.expected_priority != result {
				t.Errorf("expected %q result %q", strconv.Itoa(test.expected_priority), strconv.Itoa(result))
			}
		}
	})
}
