package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func rtoi(r rune) (int, bool) {
	switch {
	case r >= 49 && r <= 57:
		return int(r - 48), true
	default:
		return 0, false
	}
}

func main() {
	file, err := os.Open("../input.txt")
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

		for _, r := range line {
			if v, ok := rtoi(r); ok {
				first = v
				break
			}
		}

		for i := range line {
			r := rune(line[len(line)-1-i])
			if v, ok := rtoi(r); ok {
				last = v
				break
			}
		}

		sum += (first * 10) + last
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result: ", sum)
}
