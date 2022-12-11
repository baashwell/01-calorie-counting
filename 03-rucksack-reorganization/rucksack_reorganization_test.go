package rucksack_reorganization

import "testing"

var character_priority_test_set = []struct {
	character         string
	expected_priority int
}{
	{"a", 1},
}

func TestGetPriorityValueOfCharacter(t *testing.T) {
	t.Run("Test Priority is X when character is Y", func(t *testing.T) {
		for _, test := range character_priority_test_set {
			result := GetPriorityValueOfCharacter(test.character)
			if test.expected_priority != result {
				t.Errorf("expected %q result %q", test.expected_priority, result)
			}
		}
	})
}
