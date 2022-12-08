package main

import (
	"github.com/markestedt/advent-of-code-2022/utils"
	"log"
	"sort"
)

func getCol(m map[utils.Point]int, x int) []utils.Point {
	var col []utils.Point

	for p, _ := range m {
		if p.X == x {
			col = append(col, p)
		}
	}
	sort.SliceStable(col, func(i, j int) bool {
		return col[i].Y < col[j].Y
	})

	return col
}

func getRow(m map[utils.Point]int, y int) []utils.Point {
	var row []utils.Point

	for p, _ := range m {
		if p.Y == y {
			row = append(row, p)
		}
	}

	sort.SliceStable(row, func(i, j int) bool {
		return row[i].X < row[j].X
	})

	return row
}

func main() {
	lines := utils.GetLines("day08/day08test.txt")
	grid := utils.BuildGrid(lines)

	part1 := 0
	part2 := 0

	for p, _ := range grid {
		score, visible := solve(grid, p)

		if visible {
			part1++
		}

		if score > part2 {
			part2 = score
		}

	}

	log.Println(part1)
	log.Println(part2)
}

func solve(forest map[utils.Point]int, tree utils.Point) (int, bool) {
	currentTree := forest[tree]
	scoreLeft := 0
	scoreRight := 0
	scoreUp := 0
	scoreDown := 0

	visibleLeft := true
	visibleRight := true
	visibleUp := true
	visibleDown := true

	row := getRow(forest, tree.Y)
	rowBefore := row[:tree.X]
	rowAfter := row[tree.X+1:]

	col := getCol(forest, tree.X)
	colBefore := col[:tree.Y]
	colAfter := col[tree.Y+1:]

	for i := len(rowBefore) - 1; i >= 0; i-- {
		nextTree := forest[rowBefore[i]]
		scoreLeft++
		if nextTree >= currentTree {
			visibleLeft = false
			break
		}
	}
	for i := 0; i < len(rowAfter); i++ {
		nextTree := forest[rowAfter[i]]
		scoreRight++
		if nextTree >= currentTree {
			visibleRight = false
			break
		}
	}

	for i := len(colBefore) - 1; i >= 0; i-- {
		nextTree := forest[colBefore[i]]
		scoreUp++
		if nextTree >= currentTree {
			visibleUp = false
			break
		}
	}
	for i := 0; i < len(colAfter); i++ {
		nextTree := forest[colAfter[i]]
		scoreDown++
		if nextTree >= currentTree {
			visibleDown = false
			break
		}
	}
	visible := visibleLeft || visibleRight || visibleUp || visibleDown
	return scoreLeft * scoreRight * scoreUp * scoreDown, visible
}
