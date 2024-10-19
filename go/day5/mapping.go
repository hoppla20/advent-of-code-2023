package day5

import (
	"aoc/errors"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
)

type Mapping struct {
	DstStart int
	SrcStart int
	Length   int
}

func (m Mapping) String() string {
	return fmt.Sprintf("Mapping{ %d, %d, %d }", m.DstStart, m.SrcStart, m.Length)
}

func (m Mapping) Convert(item int) (int, bool) {
	diff := item - m.SrcStart
	if diff < 0 || diff >= m.Length {
		return 0, false
	}
	return m.DstStart + diff, true
}

func ParseMapping(s string) (Mapping, error) {
	parseError := &errors.ParseError{
		TargetType: "Mapping",
		Input:      s,
	}

	var result Mapping
	var intermediate [3]int

	parts := strings.Split(s, " ")

	if len(parts) != 3 {
		slog.Error("A mapping has to have exactly three parts", "got", parts)
		return result, parseError
	}

	for i := 0; i < 3; i++ {
		n, err := strconv.Atoi(parts[i])
		if err != nil {
			slog.Error(err.Error())
			return result, parseError
		}
		intermediate[i] = n
	}

	result.DstStart = intermediate[0]
	result.SrcStart = intermediate[1]
	result.Length = intermediate[2]

	return result, nil
}
