package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var wordMapping = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func digitAt(line *string, pos int) (int, bool) {
	r := rune((*line)[pos])

	if r >= 49 && r <= 57 {
		log.Println("Found digit:", int(r-48))
		return int(r - 48), true
	}

	for word, digit := range wordMapping {
		if pos+len(word) < len(*line)+1 && (*line)[pos:pos+len(word)] == word {
			log.Println("Found word:", word)
			return digit, true
		}
	}

	return 0, false
}

func main() {
	file, err := os.Open("../inputs/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		var first int
		var last int

		for i := range line {
			if digit, ok := digitAt(&line, i); ok {
				first = digit
				break
			}
		}

		for i := range line {
			if digit, ok := digitAt(&line, len(line)-i-1); ok {
				last = digit
				break
			}
		}

		value := (first * 10) + last
		log.Println("Value:", value)

		sum += value
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result:", sum)
}
