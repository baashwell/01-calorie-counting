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
	t.Run("Test Get Elves Calories for elf with two food items returns total calories for elf", func(t *testing.T) {
		expected := 6000
		result := GetElvesCalories([]string{"1000", "5000"})

		if expected != result {
			t.Errorf("expected %q result %q", strconv.Itoa(expected), strconv.Itoa(result))
		}
	})
}

func TestGetElvesFoodSupply(t *testing.T) {
	t.Run("Test Get Elves Food Supply for one elf with one food item returns list of the elfs food supply", func(t *testing.T) {
		expected := [][]string{{"1000"}}
		result := GetElvesFoodSupply("1000")

		for i, elf := range expected {
			for j, food_item := range elf {
				if food_item != result[i][j] {
					t.Errorf("expected %q result %q", food_item, result[i][j])
				}
			}
		}
	})
	t.Run("Test Get Elves Food Supply for one elf with multiple food item returns list of the elfs food supply", func(t *testing.T) {
		expected := [][]string{{"1000", "2000", "3000"}}
		result := GetElvesFoodSupply("1000\n2000\n3000")

		for i, elf := range expected {
			for j, food_item := range elf {
				if food_item != result[i][j] {
					t.Errorf("expected %q result %q", food_item, result[i][j])
				}
			}
		}
	})
}
