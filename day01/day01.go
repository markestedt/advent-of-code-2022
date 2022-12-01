package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/markestedt/advent-of-code-2022/utils"
)

func main() {
	file, err := os.Open("day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	start := time.Now()

	part1, part2 := day1(file)
	timeElapsed := time.Since(start)

	log.Printf("Part1: %d. Part2: %d. Took: %dms", part1, part2, timeElapsed.Milliseconds())
}

func day1(input *os.File) (int, int) {
	elves := []int{}
	currentElf := []int{}

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		val := scanner.Text()

		if val == "" {
			// if empty line = new elf
			sum := utils.Sum(currentElf)
			elves = append(elves, sum)

			currentElf = []int{}
		} else {
			// if not empty = add to current elf
			cal, err := strconv.Atoi(val)

			if err != nil {
				log.Fatal(err)
			}
			currentElf = append(currentElf, cal)
		}
	}
	sort.Ints(elves)
	part1 := elves[len(elves)-1]
	part2 := utils.Sum(utils.Last(elves, 3))

	return part1, part2
}
