package main

import "github.com/markestedt/advent-of-code-2022/utils"

type knot struct {
	Position utils.Point
	Visited  map[utils.Point]int
}

func (k knot) getAdjacent() []utils.Point {
	//-1-1, 0-1, 1-1
	//-10, 10
	//-11, 01, 11
	return []utils.Point{
		{X: k.Position.X - 1, Y: k.Position.Y - 1},
		{X: k.Position.X, Y: k.Position.Y - 1},
		{X: k.Position.X + 1, Y: k.Position.Y - 1},

		{X: k.Position.X + 1, Y: k.Position.Y},
		{X: k.Position.X - 1, Y: k.Position.Y},

		{X: k.Position.X - 1, Y: k.Position.Y + 1},
		{X: k.Position.X, Y: k.Position.Y + 1},
		{X: k.Position.X + 1, Y: k.Position.Y + 1},
	}
}

func main() {
	instructions := utils.GetLines("day09.txt")

	startingPoint := utils.Point{X: 0, Y: 0}

	head := knot{Position: startingPoint, Visited: map[utils.Point]int{startingPoint: 1}}
	tail := knot{Position: startingPoint, Visited: map[utils.Point]int{startingPoint: 1}}

}
