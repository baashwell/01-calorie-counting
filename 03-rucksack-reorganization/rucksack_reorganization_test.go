package rucksack_reorganization

import "testing"

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
