package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMapping(t *testing.T) {
	inputs := map[string]Mapping{
		"1 2 3":    {1, 2, 3},
		"12 23 34": {12, 23, 34},
	}

	for input, expected := range inputs {
		mapping, error := ParseMapping(input)

		assert.Nil(t, error)
		assert.Equal(t, mapping, expected)
	}
}

func TestParseMappingInvalid(t *testing.T) {
	inputs := []string{
		"1 1 2 2",
		"a b c",
	}

	for _, input := range inputs {
		mapping, err := ParseMapping(input)

		assert.NotNil(t, err)
		assert.Zero(t, Mapping{}, mapping)
	}
}
