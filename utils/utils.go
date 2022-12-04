package utils

import (
	"bufio"
	"log"
	"os"
)

func Sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func Last[T any](array []T, count int) []T {
	result := []T{}

	if count < 1 {
		return result
	}
	return array[len(array)-count:]
}

func Split(s string) (string, string) {
	middle := len(s) / 2

	return s[:middle], s[middle:]
}

func Chunks(arr []string, size int) [][]string {
	var chunks [][]string
	for i := 0; i < len(arr); i += size {
		end := i + size

		if end > len(arr) {
			end = len(arr)
		}

		chunks = append(chunks, arr[i:end])
	}

	return chunks
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
