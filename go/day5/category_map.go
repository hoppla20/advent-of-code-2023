package day5

import (
	"aoc/errors"
	"fmt"
	"regexp"
)

type DstCategory struct {
	Name     string
	Mappings []Mapping
}

func (c DstCategory) String() string {
	return fmt.Sprintf("DstCategory{ dst: %s }", c.Name)
}

func (c DstCategory) Convert(item int) int {
	result := item

	for _, mapping := range c.Mappings {
		n, applicable := mapping.Convert(item)
		if applicable {
			result = n
			break
		}
	}

	return result
}

func ParseCategoryMapHeading(s string) (string, string, error) {
	parseError := errors.ParseError{
		TargetType: "Category Map Heading",
		Input:      s,
	}

	r := regexp.MustCompile(`^([a-z]+)-to-([a-z]+) map:$`)
	matches := r.FindStringSubmatch(s)

	if len(matches) != 3 {
		return "", "", &parseError
	}

	return matches[1], matches[2], nil
}

func ParseCategoryMap(input []string) (string, DstCategory, error) {
	var srcCategory string
	result := DstCategory{}

	srcCategory, dstCategory, err := ParseCategoryMapHeading(input[0])
	if err != nil {
		return "", DstCategory{}, err
	}
	result.Name = dstCategory

	for _, line := range input[1:] {
		mapping, err := ParseMapping(line)
		if err != nil {
			return "", DstCategory{}, err
		}
		result.Mappings = append(result.Mappings, mapping)
	}

	return srcCategory, result, nil
}

type CategoryMap map[string]DstCategory
