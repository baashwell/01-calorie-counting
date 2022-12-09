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
	elves := 0
	elves_food_supply := make([][]string, 1)
	split_elves_food_supply := strings.Split(elves_food_supply_string, "\n")

	for _, food_item := range split_elves_food_supply {
		if food_item == "" {
			elves = elves + 1
			elves_food_supply = append(elves_food_supply, make([]string, 0))
		} else {
			elves_food_supply[elves] = append(elves_food_supply[elves], food_item)
		}
	}

	return elves_food_supply
}

func GetHighestCalorieAmount(elves_food_supply_string string) int {
	highest_amount := 0

	elves_food_supply := GetElvesFoodSupply(elves_food_supply_string)

	for _, elf := range elves_food_supply {
		elf_calories := GetElvesCalories(elf)
		if elf_calories > highest_amount {
			highest_amount = elf_calories
		}
	}

	return highest_amount
}
