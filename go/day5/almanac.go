package day5

import (
	"aoc/errors"
	"log/slog"
	"strconv"
	"strings"
)

func ParseSeeds(s string) ([]int, error) {
	_, numbersString, found := strings.Cut(s, ": ")
	var result []int
	parseError := errors.ParseError{
		TargetType: "Seed List",
		Input:      s,
	}

	if !found {
		return []int{}, &parseError
	}

	for _, s := range strings.Split(numbersString, " ") {
		n, err := strconv.Atoi(s)
		if err != nil {
			slog.Error(err.Error())
			return []int{}, &parseError
		}
		result = append(result, n)
	}

	return result, nil
}

func ParseAlmanac(lines []string) ([]int, CategoryMap, error) {
	categoryMap := make(CategoryMap)

	seeds, err := ParseSeeds(lines[0])
	if err != nil {
		return []int{}, CategoryMap{}, err
	}

	lines = lines[2:]

	start := 0
	for i := 0; i <= len(lines); i++ {
		if i == len(lines) || lines[i] == "" {
			srcCategory, dstCategory, err := ParseCategoryMap(lines[start:i])
			if err != nil {
				return []int{}, CategoryMap{}, err
			}

			categoryMap[srcCategory] = dstCategory

			start = i + 1
		}
	}

	return seeds, categoryMap, nil
}
