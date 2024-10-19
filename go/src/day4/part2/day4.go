package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func ReadFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

type ScratchCard struct {
	wNumbers []int
	yNumbers []int
	num      int
}

func parseCardsString(s *string) []int {
	var result []int

	for i := 0; i < len(*s); i = i + 3 {
		n, e := strconv.Atoi(strings.Trim((*s)[i:i+2], " "))
		if e != nil {
			log.Fatal(e)
		}
		result = append(result, n)
	}

	return result
}

func parseLine(line *string) ScratchCard {
	var card ScratchCard

	parts := strings.Split(*line, ": ")

	cParts := strings.Split(parts[0], " ")
	n, e := strconv.Atoi(cParts[len(cParts)-1])
	if e != nil {
		log.Fatal(e)
	}
	card.num = n

	numParts := strings.Split(parts[1], " | ")

	card.wNumbers = parseCardsString(&numParts[0])
	card.yNumbers = parseCardsString(&numParts[1])

	return card
}

type gameElement struct {
	numberOfCards   int
	numberOfWinning int
}

func main() {
	input := ReadFile("../input.txt")

	var game []gameElement

	// parse cards
	for _, line := range input {
		card := parseLine(&line)
		numW := 0

		for _, n := range card.yNumbers {
			if slices.Contains(card.wNumbers, n) {
				numW++
			}
		}

		game = append(game, gameElement{1, numW})
	}

	// duplicate cards
	for cardNum, g := range game {
		for i := 1; i <= g.numberOfWinning; i++ {
			game[cardNum+i].numberOfCards += g.numberOfCards
		}
	}

	// calculate points
	numCards := 0
	for _, g := range game {
		numCards += g.numberOfCards
		fmt.Println(g)
	}

	fmt.Println(numCards)
}
