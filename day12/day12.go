package main

import (
	"fmt"
	"github.com/markestedt/advent-of-code-2022/utils"
)

func main() {
	lines := utils.GetLines("day12/day12.txt")
	var start,end utils.Point
	height := map[utils.Point]rune{}

	for x, s := range lines {
		for y, r := range s {
			height[utils.Point{X: x, Y: y}] = r
			if r == 'S' {
				start = utils.Point{X: x, Y: y}
			} else if r == 'E' {
				end = utils.Point{X: x, Y: y}
			}
		}
	}
	height[start], height[end] = 'a', 'z'

	dist := map[utils.Point]int{end: 0}
	queue := []utils.Point{end}
	var shortest *utils.Point

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if height[cur] == 'a' && shortest == nil {
			shortest = &cur
		}

		for _, d := range []utils.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
			next := cur.Add(d)
			_, seen := dist[next]
			_, valid := height[next]
			if seen || !valid || height[next] < height[cur]-1 {
				continue
			}

			dist[next] = dist[cur] + 1
			queue = append(queue, next)
		}
	}

	fmt.Println(dist[start])
	fmt.Println(dist[*shortest])
}
