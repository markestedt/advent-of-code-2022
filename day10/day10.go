package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/markestedt/advent-of-code-2022/utils"
)

func draw(cycle int, spritePositions []int, rowNumber int) string {
	for _, position := range spritePositions {
		if (cycle-1)-(rowNumber*40) == position {
			return "#"
		}
	}
	return " "
}

func main() {
	lines := utils.GetLines("day10.txt")

	cycles := make(map[int]int)
	x := 1
	cycle := 0

	var image [6][40]string
	var row [40]string
	var crtPosition int
	spritePositions := []int{0, 1, 2}
	rowNumber := 0

	var val int

	for _, line := range lines {

		commands := strings.Split(line, " ")
		command := commands[0]

		if len(commands) > 1 {
			val, _ = strconv.Atoi(commands[1])
		} else {
			val = 0
		}

		cycle++
		cycles[cycle] = x

		crtPosition = (cycle - 1) - (rowNumber * 40)
		row[crtPosition] = draw(cycle, spritePositions, rowNumber)

		if cycle%40 == 0 {
			image[rowNumber] = row
			rowNumber++
		}

		if command == "addx" {
			cycle++
			cycles[cycle] = x

			crtPosition = (cycle - 1) - (rowNumber * 40)
			row[crtPosition] = draw(cycle, spritePositions, rowNumber)

			if cycle%40 == 0 {
				image[rowNumber] = row
				rowNumber++
			}

			x += val

			spritePositions[0] = x - 1
			spritePositions[1] = x
			spritePositions[2] = x + 1
		}
	}

	answer := (cycles[20] * 20) + (cycles[60] * 60) + (cycles[100] * 100) + (cycles[140] * 140) + (cycles[180] * 180) + (cycles[220] * 220)
	log.Println(answer)
	for _, row := range image {
		log.Println(row)
	}
}
