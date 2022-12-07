package calorie_counting

import (
	"strconv"
	"strings"
)

func GetElvesCalories(food_items []string) int {
	total_calories := 0

	for _, food_item := range food_items {
		calories, _ := strconv.Atoi(food_item)
		total_calories += calories
	}

	return total_calories
}

func GetElvesFoodSupply(elves_food_supply_string string) [][]string {
	elves_food_supply := make([][]string, 1)
	split_elves_food_supply := strings.Split(elves_food_supply_string, "\n")

	for _, food_item := range split_elves_food_supply {
		elves_food_supply[0] = append(elves_food_supply[0], food_item)
	}

	return elves_food_supply
}
