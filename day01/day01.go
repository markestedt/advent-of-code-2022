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
	start := time.Now()

	part1 := day1part1()
	part1Time := time.Since(start)

	part2 := day1part2()
	part2Time := time.Since(start)

	log.Printf("Part1. Answer: %d Took: %dms", part1, part1Time.Milliseconds())
	log.Printf("Part2. Answer: %d Took: %dms", part2, part2Time.Milliseconds())
}

func day1part1() int {
	f, err := os.Open("day1.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	currentElf := []int{}

	scanner := bufio.NewScanner(f)
	var maxVal = 0

	for scanner.Scan() {
		val := scanner.Text()

		if val == "" {
			// if empty line = new elf
			sum := utils.Sum(currentElf)

			if sum > maxVal {
				maxVal = sum
			}

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
	return maxVal
}

func day1part2() int {
	f, err := os.Open("day1.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	elves := []int{}
	currentElf := []int{}

	scanner := bufio.NewScanner(f)

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

	return elves[len(elves)-1] + elves[len(elves)-2] + elves[len(elves)-3]
}
