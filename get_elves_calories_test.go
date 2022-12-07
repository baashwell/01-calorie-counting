package calorie_counting

import (
	"strconv"
	"testing"
)

func TestGetElvesCalories(t *testing.T) {
	t.Run("Test Get Elves Calories for elf with one food item returns total calories for elf", func(t *testing.T) {
		expected := 1000
		result := GetElvesCalories([]string{"1000"})

		if expected != result {
			t.Errorf("expected %q result %q", strconv.Itoa(expected), strconv.Itoa(result))
		}
	})
}
