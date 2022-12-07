package calorie_counting

import "strconv"

func GetElvesCalories(food_items []string) int {
	total_calories := 0

	for _, food_item := range food_items {
		calories, _ := strconv.Atoi(food_item)
		total_calories += calories
	}

	return total_calories
}
