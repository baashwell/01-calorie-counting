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
	t.Run("Test Get Elves Food Supply for multiple elves with one food item each returns list of the elves food supply", func(t *testing.T) {
		expected := [][]string{{"1000"}, {"2000"}, {"3000"}}
		result := GetElvesFoodSupply("1000\n\n2000\n\n3000")

		for i, elf := range expected {
			for j, food_item := range elf {
				if food_item != result[i][j] {
					t.Errorf("expected %q result %q", food_item, result[i][j])
				}
			}
		}
	})
	t.Run("Test Get Elves Food Supply for multiple elves with multiple food items each returns list of the elves food supply", func(t *testing.T) {
		expected := [][]string{{"1000", "2000"}, {"2000"}, {"3000", "1000", "7000"}}
		result := GetElvesFoodSupply("1000\n2000\n\n2000\n\n3000\n1000\n7000")

		for i, elf := range expected {
			for j, food_item := range elf {
				if food_item != result[i][j] {
					t.Errorf("expected %q result %q", food_item, result[i][j])
				}
			}
		}
	})
}

func TestGetElvesWithHighestCalories(t *testing.T) {
	t.Run("Test Get Highest Calorie Amount for one elf with one food item returns calories for the one food item", func(t *testing.T) {
		expected := 1000
		result := GetElvesWithHighestCalories("1000", 1)[0]

		if expected != result {
			t.Errorf("expected %q result %q", strconv.Itoa(expected), strconv.Itoa(result))
		}
	})
	t.Run("Test Get Highest Calorie Amount for one elf with multiple food items returns calories for the one food item", func(t *testing.T) {
		expected := 3000
		result := GetElvesWithHighestCalories("1000\n2000", 1)[0]

		if expected != result {
			t.Errorf("expected %q result %q", strconv.Itoa(expected), strconv.Itoa(result))
		}
	})
	t.Run("Test Get Highest Calorie Amount for multiple elves with multiple food items returns calories for the one food item", func(t *testing.T) {
		expected := 11000
		result := GetElvesWithHighestCalories("1000\n2000\n\n2000\n\n3000\n1000\n7000", 1)[0]

		if expected != result {
			t.Errorf("expected %q result %q", strconv.Itoa(expected), strconv.Itoa(result))
		}
	})
	t.Run("Test Get top 2 Highest Calorie Amount for multiple elves with multiple food items returns calories for the one food item", func(t *testing.T) {
		expected := []int{11000, 3000}
		result := GetElvesWithHighestCalories("1000\n2000\n\n2000\n\n3000\n1000\n7000", 2)

		for i, calorie_amount := range expected {
			if calorie_amount != result[i] {
				t.Errorf("expected %q result %q", strconv.Itoa(calorie_amount), strconv.Itoa(result[i]))
			}
		}
	})
}

func TestGetPuzzleAnswer(t *testing.T) {
	t.Run("Run", func(t *testing.T) {
		expected := 70613
		result := GetElvesWithHighestCalories(INPUT, 1)

		if expected != result[0] {
			t.Errorf("expected %q result %q", strconv.Itoa(expected), strconv.Itoa(result[0]))
		}
	})
}
