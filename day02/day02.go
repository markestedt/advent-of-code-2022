package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

type Score struct {
	part1 int
	part2 int
}

func main() {
	file, err := os.Open("day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	start := time.Now()

	points := map[string]Score{
		"A X": {4, 3},
		"B X": {1, 1},
		"C X": {7, 2},

		"A Y": {8, 4},
		"B Y": {5, 5},
		"C Y": {2, 6},

		"A Z": {3, 8},
		"B Z": {9, 9},
		"C Z": {6, 7},
	}

	part1, part2 := day2(file, points)
	timeElapsed := time.Since(start)
	log.Printf("Part1: %d. Part2: %d. Took: %dms", part1, part2, timeElapsed.Milliseconds())
}

func day2(input *os.File, points map[string]Score) (int, int) {
	scorePart1 := 0
	scorePart2 := 0
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		val := scanner.Text()
		scorePart1 += points[val].part1
		scorePart2 += points[val].part2
	}
	return scorePart1, scorePart2
}
