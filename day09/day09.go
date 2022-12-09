package main

import (
	"github.com/markestedt/advent-of-code-2022/utils"
	"log"
	"strconv"
	"strings"
)

type knot struct {
	Position utils.Point
	Visited  []utils.Point
}

func main() {
	instructions := utils.GetLines("day09/day09.txt")
	part1(instructions)
	part2(instructions)
}

func move(head *knot, tail *knot) {
	diffX := head.Position.X - tail.Position.X
	diffY := head.Position.Y - tail.Position.Y

	if utils.Abs(diffX) > 1 || utils.Abs(diffY) > 1 {
		tail.Position.X = tail.Position.X + utils.SignInt(diffX)
		tail.Position.Y = tail.Position.Y + utils.SignInt(diffY)

		if !utils.Contains(tail.Visited, tail.Position) {
			tail.Visited = append(tail.Visited, tail.Position)
		}
	}
}

func part1(instructions []string) {
	startingPoint := utils.Point{X: 0, Y: 0}

	head := knot{Position: startingPoint, Visited: []utils.Point{startingPoint}}
	tail := knot{Position: startingPoint, Visited: []utils.Point{startingPoint}}

	for _, inst := range instructions {
		x := strings.Split(inst, " ")
		direction := x[0]
		distance, _ := strconv.Atoi(x[1])

		for i := 0; i < distance; i++ {
			switch direction {
			case "R":
				head.Position.X++
			case "L":
				head.Position.X--
			case "U":
				head.Position.Y++
			case "D":
				head.Position.Y--
			}

			move(&head, &tail)
		}
	}

	log.Println(len(tail.Visited))
}

func part2(instructions []string) {
	startingPoint := utils.Point{X: 0, Y: 0}

	head := knot{Position: startingPoint, Visited: []utils.Point{startingPoint}}
	one := knot{Position: startingPoint, Visited: []utils.Point{startingPoint}}
	two := knot{Position: startingPoint, Visited: []utils.Point{startingPoint}}
	three := knot{Position: startingPoint, Visited: []utils.Point{startingPoint}}
	four := knot{Position: startingPoint, Visited: []utils.Point{startingPoint}}
	five := knot{Position: startingPoint, Visited: []utils.Point{startingPoint}}
	six := knot{Position: startingPoint, Visited: []utils.Point{startingPoint}}
	seven := knot{Position: startingPoint, Visited: []utils.Point{startingPoint}}
	eight := knot{Position: startingPoint, Visited: []utils.Point{startingPoint}}
	nine := knot{Position: startingPoint, Visited: []utils.Point{startingPoint}}

	knots := []*knot{&head, &one, &two, &three, &four, &five, &six, &seven, &eight, &nine}

	for _, inst := range instructions {
		x := strings.Split(inst, " ")
		direction := x[0]
		distance, _ := strconv.Atoi(x[1])

		for i := 0; i < distance; i++ {
			switch direction {
			case "R":
				head.Position.X++
			case "L":
				head.Position.X--
			case "U":
				head.Position.Y++
			case "D":
				head.Position.Y--
			}

			for i := 1; i < len(knots); i++ {
				move(knots[i-1], knots[i])
			}
		}
	}
	log.Println(len(nine.Visited))
}
