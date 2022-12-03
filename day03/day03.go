package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"

	"github.com/markestedt/advent-of-code-2022/utils"
)

func main() {
	file, err := os.Open("day03.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	start := time.Now()

	alphabet := map[rune]int{
		'a': 1,
		'b': 2,
		'c': 3,
		'd': 4,
		'e': 5,
		'f': 6,
		'g': 7,
		'h': 8,
		'i': 9,
		'j': 10,
		'k': 11,
		'l': 12,
		'm': 13,
		'n': 14,
		'o': 15,
		'p': 16,
		'q': 17,
		'r': 18,
		's': 19,
		't': 20,
		'u': 21,
		'v': 22,
		'w': 23,
		'x': 24,
		'y': 25,
		'z': 26,
		'A': 27,
		'B': 28,
		'C': 29,
		'D': 30,
		'E': 31,
		'F': 32,
		'G': 33,
		'H': 34,
		'I': 35,
		'J': 36,
		'K': 37,
		'L': 38,
		'M': 39,
		'N': 40,
		'O': 41,
		'P': 42,
		'Q': 43,
		'R': 44,
		'S': 45,
		'T': 46,
		'U': 47,
		'V': 48,
		'W': 49,
		'X': 50,
		'Y': 51,
		'Z': 52,
	}

	part1, part2 := day3(file, alphabet)

	timeElapsed := time.Since(start)
	log.Printf("Part1: %d. Part2: %d. Took: %dms", part1, part2, timeElapsed.Milliseconds())
}

func day3(input *os.File, alphabet map[rune]int) (int, int) {
	scanner := bufio.NewScanner(input)

	var lines []string
	var prios1 []int
	for scanner.Scan() {
		val := scanner.Text()
		lines = append(lines, val)

		first, last := utils.Split(val)

		for _, r := range first {
			if strings.ContainsRune(last, r) {
				prios1 = append(prios1, alphabet[r])
				break
			}
		}
	}

	chunks := utils.Chunks(lines, 3)
	var prios2 []int

	for _, c := range chunks {
		for _, r := range c[0] {
			if strings.ContainsRune(c[1], r) && strings.ContainsRune(c[2], r) {
				prios2 = append(prios2, alphabet[r])
				break
			}
		}
	}

	return utils.Sum(prios1), utils.Sum(prios2)
}
