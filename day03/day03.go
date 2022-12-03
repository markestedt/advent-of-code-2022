package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"

	"github.com/markestedt/advent-of-code-2022/utils"
)

const Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	file, err := os.Open("day03.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	start := time.Now()

	part1, part2 := day3(file)

	timeElapsed := time.Since(start)
	log.Printf("Part1: %d. Part2: %d. Took: %dms", part1, part2, timeElapsed.Milliseconds())
}

func day3(input *os.File) (int, int) {
	scanner := bufio.NewScanner(input)

	var lines []string
	var prios1 []int
	for scanner.Scan() {
		val := scanner.Text()
		lines = append(lines, val)

		first, last := utils.Split(val)

		for _, r := range first {
			if strings.ContainsRune(last, r) {
				prios1 = append(prios1, getPriority(r))
				break
			}
		}
	}

	chunks := utils.Chunks(lines, 3)
	var prios2 []int

	for _, c := range chunks {
		for _, r := range c[0] {
			if strings.ContainsRune(c[1], r) && strings.ContainsRune(c[2], r) {
				prios2 = append(prios2, getPriority(r))
				break
			}
		}
	}

	return utils.Sum(prios1), utils.Sum(prios2)
}

func getPriority(r rune) int {
	return strings.IndexRune(Alphabet, r) + 1
}
