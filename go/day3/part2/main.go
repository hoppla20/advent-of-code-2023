package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

/*
* Helper functions
 */

func rtoi(r rune) (int, bool) {
	switch {
	case r >= 48 && r <= 57:
		return int(r - 48), true
	default:
		return 0, false
	}
}

/*
* Position type
 */

type position struct {
	x int
	y int
}

func (p *position) left() position {
	return position{
		x: p.x - 1,
		y: p.y,
	}
}

func (p *position) right() position {
	return position{
		x: p.x + 1,
		y: p.y,
	}
}

func (p *position) up() position {
	return position{
		x: p.x,
		y: p.y - 1,
	}
}

func (p *position) down() position {
	return position{
		x: p.x,
		y: p.y + 1,
	}
}

/*
* EngineSchematic type
 */

type engineSchematic [][]rune

func makeEngineSchematic(input *[]byte) (result engineSchematic) {
	lines := strings.Split(string(*input), "\n")

	result = make([][]rune, len(lines))
	for i, line := range lines {
		result[i] = []rune(line)
	}

	return
}

func (s *engineSchematic) runeAtPos(pos position) (rune, bool) {
	if pos.x < 0 || pos.y < 0 {
		return 0, false
	}
	if pos.y >= len(*s) || pos.x >= len((*s)[pos.y]) {
		return 0, false
	}
	return (*s)[pos.y][pos.x], true
}

func (s *engineSchematic) digitAtPos(pos position) (int, bool) {
	r, ok := s.runeAtPos(pos)
	if !ok {
		return 0, false
	}

	result, ok := rtoi(r)
	if !ok {
		return 0, false
	}

	return result, true
}

func (s *engineSchematic) numberAtPos(pos position) (int, int, bool) {
	rune, ok := s.runeAtPos(pos)
	if !ok {
		return 0, 0, false
	}

	result, isDigit := rtoi(rune)
	if !isDigit {
		return 0, 0, false
	}

	length := 1
	curPos := pos.right()
	curDigit, ok := s.digitAtPos(curPos)
	for ok {
		result = result*10 + curDigit
		length++
		curPos = curPos.right()
		curDigit, ok = s.digitAtPos(curPos)
	}

	return result, length, true
}

func (s *engineSchematic) isSymbol(pos position) bool {
	rune, ok := s.runeAtPos(pos)
	if !ok {
		return false
	}

	if _, isDigit := s.digitAtPos(pos); isDigit {
		return false
	}

	return rune != '.'
}

func (s *engineSchematic) isStar(pos position) bool {
	rune, ok := s.runeAtPos(pos)
	if !ok {
		return false
	}

	return rune == '*'
}

func (s *engineSchematic) isEnginePartNumber(pos position, length int) bool {
	if s.isSymbol(pos.left()) || s.isSymbol(position{pos.x + length, pos.y}) {
		return true
	}

	for x := pos.x - 1; x < pos.x+1+length; x++ {
		if s.isSymbol(position{x, pos.y - 1}) || s.isSymbol(position{x, pos.y + 1}) {
			return true
		}
	}

	return false
}

func (s *engineSchematic) findNumberStartPosition(pos position) (position, bool) {
	_, isDigit := s.digitAtPos(pos)

	for isDigit {
		prevPos := pos.left()
		_, isDigit = s.digitAtPos(prevPos)
		if !isDigit {
			return pos, true
		}
		pos = prevPos
	}

	return position{}, false
}

/*
* Main
 */

func main() {
	fileContent, err := os.ReadFile("../inputs/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	schema := makeEngineSchematic(&fileContent)

	result := 0
	for y := range schema {
		for x := range schema[y] {
			pos := position{x, y}
			log.Println("Looking at Position:", pos)
			if schema.isStar(pos) {
				foundNumbers := 0
				gearRatio := 1

				if numPos, isNumber := schema.findNumberStartPosition(pos.left()); isNumber {
					foundNumbers++
					num, _, _ := schema.numberAtPos(numPos)
					gearRatio *= num
				}

				if num, _, isNumber := schema.numberAtPos(pos.right()); isNumber {
					foundNumbers++
					gearRatio *= num
				}

				for _, curY := range []int{y - 1, y + 1} {
					for curX := x - 1; curX <= x+1; {
						numPos, isNumber := schema.findNumberStartPosition(position{curX, curY})
						if isNumber {
							foundNumbers++
							num, length, _ := schema.numberAtPos(numPos)
							gearRatio *= num
							curX = numPos.x + length
						} else {
							curX++
						}
					}
				}

				if foundNumbers == 2 {
					result += gearRatio
				}
			}
		}
	}

	fmt.Println("Result:", result)
}
