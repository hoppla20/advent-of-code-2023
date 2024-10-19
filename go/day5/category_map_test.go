package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCategoryMapHeading(t *testing.T) {
	inputs := map[string][2]string{
		"seed-to-soil map:":       {"seed", "soil"},
		"soil-to-fertilizer map:": {"soil", "fertilizer"},
	}

	for input, expected := range inputs {
		srcCat, dstCat, err := ParseCategoryMapHeading(input)

		assert.Nil(t, err)
		assert.Equal(t, expected, [2]string{srcCat, dstCat})
	}
}

func TestParseCategoryMapHeadingInvalid(t *testing.T) {
	inputs := []string{
		"invalid",
		"soil-in-fertilizer map:",
	}

	for _, input := range inputs {
		srcCat, dstCat, err := ParseCategoryMapHeading(input)

		assert.NotNil(t, err)
		assert.Zero(t, srcCat, dstCat)
	}
}

func TestParseCategoryMap(t *testing.T) {
	input := []string{
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
	}

	expectedSrcCategory := "seed"
	expectedDstCategory := DstCategory{
		Name: "soil",
		Mappings: []Mapping{
			{50, 98, 2},
			{52, 50, 48},
		},
	}

	srcCategory, dstCategory, err := ParseCategoryMap(input)
	assert.Nil(t, err)
	assert.Equal(t, expectedSrcCategory, srcCategory)
	assert.Equal(t, expectedDstCategory.Name, dstCategory.Name)

	assert.Equal(t, len(expectedDstCategory.Mappings), len(dstCategory.Mappings))
	for i := range dstCategory.Mappings {
		assert.Equal(t, expectedDstCategory.Mappings[i], dstCategory.Mappings[i])
	}
}
