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
	lines := utils.GetLines("day08/day08.txt")
	log.Println(part1(lines))

	grid := utils.BuildGrid(lines)
	answerPart2 := 0

	for p, _ := range grid {
		score := part2(grid, p)

		if score > answerPart2 {
			answerPart2 = score
		}
	}
	log.Println(answerPart2)
}

func part1(lines []string) int {
	var visibleTrees []utils.Point

	//left to right
	for i, line := range lines {
		highestInLine := -1
		for j, t := range line {
			val := int(t - '0')
			if val > highestInLine {
				highestInLine = val
				point := utils.Point{Y: i, X: j}
				if !utils.Contains(visibleTrees, point) {
					visibleTrees = append(visibleTrees, point)
				}
			}
		}
	}

	//right to left
	for i, line := range lines {
		highestInLine := -1
		for j := len(line) - 1; j >= 0; j-- {
			val := int(line[j] - '0')
			if val > highestInLine {
				highestInLine = val
				point := utils.Point{Y: i, X: j}
				if !utils.Contains(visibleTrees, point) {
					visibleTrees = append(visibleTrees, point)
				}
			}
		}
	}

	//top to bottom
	for i := 0; i < len(lines); i++ {
		highestInCol := -1
		for j := 0; j < len(lines); j++ {
			val := int(lines[j][i] - '0')
			if val > highestInCol {
				highestInCol = val
				point := utils.Point{Y: j, X: i}
				if !utils.Contains(visibleTrees, point) {
					visibleTrees = append(visibleTrees, point)
				}
			}
		}
	}

	//bottom to top
	for i := 0; i < len(lines); i++ {
		highestInCol := -1
		for j := len(lines) - 1; j >= 0; j-- {
			val := int(lines[j][i] - '0')
			if val > highestInCol {
				highestInCol = val
				point := utils.Point{Y: j, X: i}
				if !utils.Contains(visibleTrees, point) {
					visibleTrees = append(visibleTrees, point)
				}
			}
		}
	}

	return len(visibleTrees)
}

func part2(forest map[utils.Point]int, tree utils.Point) int {
	currentTree := forest[tree]
	scoreLeft := 0
	scoreRight := 0
	scoreUp := 0
	scoreDown := 0

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
			break
		}
	}
	for i := 0; i < len(rowAfter); i++ {
		nextTree := forest[rowAfter[i]]
		scoreRight++
		if nextTree >= currentTree {
			break
		}
	}

	for i := len(colBefore) - 1; i >= 0; i-- {
		nextTree := forest[colBefore[i]]
		scoreUp++
		if nextTree >= currentTree {
			break
		}
	}
	for i := 0; i < len(colAfter); i++ {
		nextTree := forest[colAfter[i]]
		scoreDown++
		if nextTree >= currentTree {
			break
		}
	}

	return scoreLeft * scoreRight * scoreUp * scoreDown
}
