package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSeeds(t *testing.T) {
	input := "seeds: 79 14 55 13"
	expectedSeeds := []int{79, 14, 55, 13}

	seeds, err := ParseSeeds(input)
	assert.Nil(t, err)
	assert.Equal(t, expectedSeeds, seeds)
}

func TestParseAlmanac(t *testing.T) {
	input := []string{
		"seeds: 79 14 55 13",
		"",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
	}

	expectedSeeds := []int{79, 14, 55, 13}
	expectedCategoryMap := CategoryMap{
		"seed": DstCategory{
			Name: "soil",
			Mappings: []Mapping{
				{50, 98, 2},
				{52, 50, 48},
			},
		},
	}

	seeds, categoryMap, err := ParseAlmanac(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedSeeds, seeds)
	assert.Equal(t, expectedCategoryMap, categoryMap)
}
