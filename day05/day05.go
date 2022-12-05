package main

import (
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/markestedt/advent-of-code-2022/utils"
)

func main() {
	lines := utils.GetLines("day05.txt")
	start := time.Now()
	timeElapsed := time.Since(start)
	log.Printf("Part1: %s. Part2: %s. Took: %dms", part1(lines, getStacks()), part2(lines, getStacks()), timeElapsed.Milliseconds())

}

func part2(lines []string, stacks map[int][]string) string {
	for _, line := range lines {
		re := regexp.MustCompile("[0-9]+")
		numbers := re.FindAllString(line, -1)

		count, _ := strconv.Atoi(numbers[0])
		from, _ := strconv.Atoi(numbers[1])
		to, _ := strconv.Atoi(numbers[2])

		origin := stacks[from]
		destination := stacks[to]

		crates := origin[0:count]

		newOrigin := origin[count:]

		newDestination := append([]string{}, crates...)
		newDestination = append(newDestination, destination...)

		stacks[from] = newOrigin
		stacks[to] = newDestination

	}

	return getAnswer(stacks)
}

func part1(lines []string, stacks map[int][]string) string {
	for _, line := range lines {
		re := regexp.MustCompile("[0-9]+")
		numbers := re.FindAllString(line, -1)

		count, _ := strconv.Atoi(numbers[0])
		from, _ := strconv.Atoi(numbers[1])
		to, _ := strconv.Atoi(numbers[2])

		origin := stacks[from]
		destination := stacks[to]

		for i := 0; i < count; i++ {
			crate := origin[0]
			origin = origin[1:]

			destination = append([]string{crate}, destination...)

			stacks[from] = origin
			stacks[to] = destination
		}
	}

	return getAnswer(stacks)
}

func getAnswer(stacks map[int][]string) string {
	var answer string
	for i := 1; i <= len(stacks); i++ {
		answer += stacks[i][0]
	}

	return answer
}

func getStacks() map[int][]string {
	one := []string{"P", "V", "Z", "W", "D", "T"}
	two := []string{"D", "J", "F", "V", "W", "S", "L"}
	three := []string{"H", "B", "T", "V", "S", "L", "M", "Z"}
	four := []string{"J", "S", "R"}
	five := []string{"W", "L", "M", "F", "G", "B", "Z", "C"}
	six := []string{"B", "G", "R", "Z", "H", "V", "W", "Q"}
	seven := []string{"N", "D", "B", "C", "P", "J", "V"}
	eight := []string{"Q", "B", "T", "P"}
	nine := []string{"C", "R", "Z", "G", "H"}

	stacks := map[int][]string{
		1: one,
		2: two,
		3: three,
		4: four,
		5: five,
		6: six,
		7: seven,
		8: eight,
		9: nine,
	}

	return stacks
}
