package main

import (
	"aoc/day5"
	"aoc/internal/utils"
	"fmt"
	"log"
)

func srcToDstCategory(srcCategory string, input []int, categoryMap day5.CategoryMap) (string, []int) {
	output := input

	dstCategory := categoryMap[srcCategory]
	for i, item := range input {
		output[i] = dstCategory.Convert(item)
	}

	return dstCategory.Name, output
}

func main() {
	utils.SetLogLevel()

	lines := utils.ReadFile("../inputs/input.txt")

	currentCategory := "seed"
	currentItems, categoryMap, err := day5.ParseAlmanac(lines)
	if err != nil {
		log.Fatal(err)
	}

	for currentCategory != "location" {
		currentCategory, currentItems = srcToDstCategory(currentCategory, currentItems, categoryMap)
	}

	result, _ := utils.Lowest(currentItems)
	fmt.Println("Result:", result)
}
