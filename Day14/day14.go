package main

import (
	"fmt"
	"github.com/markestedt/advent-of-code-2022/utils"
	"log"
	"strings"
)

const rock = '#'
const sand = 'o'

var down = utils.Point{X: 0, Y: 1}
var downLeft = utils.Point{X: -1, Y: 1}
var downRight = utils.Point{X: 1, Y: 1}

func main() {
	lines := utils.GetLines("Day14/day14.txt")
	cave1, floor := buildCave(lines)

	//Since map is reference type, I can't use the same 'cave' for both parts.
	cave2 := utils.Copy(cave1)

	part1(cave1, floor)
	part2(cave2, floor+2)
}

func part2(cave map[utils.Point]rune, floor int) {
	startingPoint := utils.Point{X: 500, Y: 0}
	done := false
	unitsOfSand := 0

	for done == false {
		sandPoint := startingPoint

		rest, restPoint := fallPart2(cave, sandPoint, floor)

		if rest {
			unitsOfSand++
		}
		done = restPoint == startingPoint
	}
	log.Println(unitsOfSand)
}

func part1(cave map[utils.Point]rune, floor int) {
	startingPoint := utils.Point{X: 500, Y: 0}
	done := false
	unitsOfSand := 0

	for done == false {
		sandPoint := startingPoint

		rest := fallPart1(cave, sandPoint, floor)
		if rest {
			unitsOfSand++
		}
		done = !rest
	}
	log.Println(unitsOfSand)
}

func fallPart2(cave map[utils.Point]rune, point utils.Point, floor int) (bool, utils.Point) {
	pointDown := point.Add(down)
	pointDownLeft := point.Add(downLeft)
	pointDownRight := point.Add(downRight)

	blockedDown := part2Blocked(cave, pointDown, floor)
	blockedDownLeft := part2Blocked(cave, pointDownLeft, floor)
	blockedDownRight := part2Blocked(cave, pointDownRight, floor)

	if blockedDown && blockedDownLeft && blockedDownRight {
		cave[point] = sand
		return true, point
	} else if !blockedDown {
		return fallPart2(cave, pointDown, floor)
	} else if !blockedDownLeft {
		return fallPart2(cave, pointDownLeft, floor)
	} else {
		return fallPart2(cave, pointDownRight, floor)
	}
}

func fallPart1(cave map[utils.Point]rune, point utils.Point, floor int) bool {
	pointDown := point.Add(down)
	pointDownLeft := point.Add(downLeft)
	pointDownRight := point.Add(downRight)

	blockedDown := part1Blocked(cave, pointDown)
	blockedDownLeft := part1Blocked(cave, pointDownLeft)
	blockedDownRight := part1Blocked(cave, pointDownRight)

	if blockedDown && blockedDownLeft && blockedDownRight {
		cave[point] = sand
		return true
	} else if !blockedDown {
		if pointDown.Y >= floor {
			return false
		}
		return fallPart1(cave, pointDown, floor)
	} else if !blockedDownLeft {
		if pointDownLeft.Y >= floor {
			return false
		}
		return fallPart1(cave, pointDownLeft, floor)
	} else {
		if pointDownRight.Y >= floor {
			return false
		}
		return fallPart1(cave, pointDownRight, floor)
	}
}

func part2Blocked(cave map[utils.Point]rune, point utils.Point, floor int) bool {
	if point.Y >= floor {
		return true
	}
	return part1Blocked(cave, point)
}

func part1Blocked(cave map[utils.Point]rune, point utils.Point) bool {
	if val, ok := cave[point]; ok {
		return val == rock || val == sand
	}
	return false
}

func buildCave(lines []string) (map[utils.Point]rune, int) {
	cave := make(map[utils.Point]rune)
	floor := 0

	for _, path := range lines {
		points := strings.Split(path, " -> ")

		for i, p := range points {
			var x, y int
			fmt.Sscanf(p, "%d,%d", &x, &y)

			currentPoint := utils.Point{X: x, Y: y}
			cave[currentPoint] = rock
			if y > floor {
				floor = y
			}

			if i < len(points)-1 {
				var nextX, nextY int
				fmt.Sscanf(points[i+1], "%d,%d", &nextX, &nextY)

				if nextX != x {
					if x > nextX {
						for j := x - 1; j > nextX; j-- {
							point := utils.Point{X: j, Y: y}
							cave[point] = rock
						}
					} else {
						for j := x + 1; j < nextX; j++ {
							point := utils.Point{X: j, Y: y}
							cave[point] = rock
						}
					}
				} else if nextY != y {
					if y > nextY {
						for j := y - 1; j > nextY; j-- {
							point := utils.Point{X: x, Y: j}
							cave[point] = rock
						}
					} else {
						for j := y + 1; j < nextY; j++ {
							point := utils.Point{X: x, Y: j}
							cave[point] = rock
						}
					}
				}
			}
		}
	}
	return cave, floor
}
