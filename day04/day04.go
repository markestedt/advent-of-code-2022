package main

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/markestedt/advent-of-code-2022/utils"
)

type Elf struct {
	Start int
	End   int
}

func main() {
	lines := utils.GetLines("day04.txt")
	start := time.Now()

	contained := 0
	overlapping := 0

	for _, line := range lines {
		pairs := strings.Split(line, ",")

		pair1 := strings.Split(pairs[0], "-")
		pair2 := strings.Split(pairs[1], "-")

		firstElfStart, _ := strconv.Atoi(pair1[0])
		firstElfEnd, _ := strconv.Atoi(pair1[1])
		secondElfStart, _ := strconv.Atoi(pair2[0])
		secondElfEnd, _ := strconv.Atoi(pair2[1])

		firstElf := Elf{firstElfStart, firstElfEnd}
		secondElf := Elf{secondElfStart, secondElfEnd}

		if isContainedBy(firstElf, secondElf) ||
			isContainedBy(secondElf, firstElf) {
			contained++
			overlapping++
		} else if isOverlapping(firstElf, secondElf) {
			overlapping++
		}
	}

	timeElapsed := time.Since(start)
	log.Printf("Part1: %d. Part2: %d. Took: %dms", contained, overlapping, timeElapsed.Milliseconds())
}

func isOverlapping(firstElf Elf, secondElf Elf) bool {
	return firstElf.End >= secondElf.Start && firstElf.Start <= secondElf.End
}

func isContainedBy(firstElf Elf, secondElf Elf) bool {
	return firstElf.Start >= secondElf.Start && firstElf.End <= secondElf.End
}
