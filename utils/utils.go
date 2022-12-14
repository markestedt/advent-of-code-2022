package utils

import (
	"bufio"
	"log"
	"os"
)

type Number interface {
	int | int32 | int64 | float32 | float64
}

func Sum[T Number](array []T) T {
	var result T
	for _, v := range array {
		result += v
	}
	return result
}

func Copy[T1 comparable, T2 any](source map[T1]T2) map[T1]T2{
	target := make(map[T1]T2, len(source))
	for key, value := range source {
		target[key] = value
	}
	return target
}

func Last[T any](array []T, count int) []T {
	result := []T{}

	if count < 1 {
		return result
	}
	return array[len(array)-count:]
}

func Contains[T comparable](s []T, e T) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Split(s string) (string, string) {
	middle := len(s) / 2

	return s[:middle], s[middle:]
}

func Chunks[T any](arr []T, size int) [][]T {
	var chunks [][]T
	for i := 0; i < len(arr); i += size {
		end := i + size

		if end > len(arr) {
			end = len(arr)
		}

		chunks = append(chunks, arr[i:end])
	}

	return chunks
}

func SignInt(x int) int {
	if x < 0 {
		return -1
	}
	if x == 0 {
		return 0
	}
	return 1
}

func Abs(x int) int {
	if x < 0 {
		return 0 - x
	}
	return x - 0
}

type Point struct {
	X int
	Y int
}

func (p Point) Add (x Point) Point {
	return Point{X: p.X + x.X, Y: p.Y + x.Y}
}

func BuildGrid(lines []string) map[Point]int {
	var grid = make(map[Point]int)

	for i, line := range lines {
		for j, val := range line {

			newPoint := Point{X: j, Y: i}
			height := int(val - '0')

			grid[newPoint] = height
		}
	}
	return grid
}

func GetLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}

	return lines
}
