package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type gameSet struct {
	red   int
	green int
	blue  int
}

var maxColorNum = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func parseGameString(gameString string) []gameSet {
	var parsedGame []gameSet

	gameElements := strings.FieldsFunc(gameString, func(r rune) bool {
		if r == ':' || r == ';' {
			return true
		}
		return false
	})

	log.Println(gameElements[0])
	gameNum, err := strconv.Atoi(strings.Fields(gameElements[0])[1])
	if err != nil {
		log.Fatal("Failed parse game number ", gameNum)
	}
	log.Println("Parsed Game Number:", gameNum)

	log.Println("Elements:", gameElements[1:])
	for i, set := range gameElements[1:] {
		colorStrings := strings.Split(set, ", ")
		log.Println("Game", gameNum, "Set", i)
		parsedSet := gameSet{}
		for _, color := range colorStrings {
			colorParts := strings.Split(strings.Trim(color, " "), " ")
			num, err := strconv.Atoi(colorParts[0])
			if err != nil {
				log.Println("Failed to parse color number:")
				log.Fatal(err)
			}
			c := colorParts[1]
			log.Println("Color", c, "Num", num)
			switch c {
			case "red":
				parsedSet.red = num
			case "green":
				parsedSet.green = num
			case "blue":
				parsedSet.blue = num
			default:
				log.Fatal("Could not parse color:", c)
			}
		}
		log.Println("Parsed Set:", parsedSet)
		parsedGame = append(parsedGame, parsedSet)
	}
	log.Println("Parsed Game:", parsedGame)

	return parsedGame
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var parsedGames [][]gameSet
	for scanner.Scan() {
		gameString := scanner.Text()
		parsedGames = append(parsedGames, parseGameString(gameString))
	}
	log.Println("Parsed Games:", parsedGames)

	result := 0
	for i, game := range parsedGames {
		fail := false
		for _, set := range game {
			if set.red > maxColorNum["red"] || set.green > maxColorNum["green"] || set.blue > maxColorNum["blue"] {
				fail = true
				break
			}
		}

		log.Println("Game", i, !fail)

		if !fail {
			result += i + 1
		}
	}

	fmt.Println("Result:", result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
