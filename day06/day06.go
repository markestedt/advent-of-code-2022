package main

import (
	"log"
	"time"

	"github.com/markestedt/advent-of-code-2022/utils"
)

func main() {
	input := utils.GetLines("day06/day06.txt")[0]
	start := time.Now()

	part1 := solve(input, 4)
	part2 := solve(input, 14)

	timeElapsed := time.Since(start)
	log.Printf("Part1: %d. Part2: %d. Took: %dms", part1, part2, timeElapsed.Milliseconds())
}

func solve(input string, markerSize int) int {
	var answer int

	for i := markerSize; i < len(input); i++ {
		marker := input[i-markerSize : i]

		if !hasDuplicate(marker) {
			answer = i
			break
		}
	}
	return answer
}

func hasDuplicate(input string) bool {
	var visited []rune

	for _, s := range input {
		if utils.Contains(visited, s) {
			return true
		}
		visited = append(visited, s)
	}
	return false
}
