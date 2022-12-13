package main

import (
	_ "embed"
	"encoding/json"
	"log"
	"os"
	"sort"
	"strings"
)


func main() {
	data,_ := os.ReadFile("Day13/day13test.txt")
	input := strings.Split(string(data), "\n\n")

	var packets []any
	var sum int

	for i, pairs := range input {
		pair := strings.Split(pairs, "\n")

		var first any
		var second any

		json.Unmarshal([]byte(pair[0]), &first)
		json.Unmarshal([]byte(pair[1]), &second)

		packets = append(packets, first, second)
		if equal(first, second) <= 0 {
			sum += i + 1
		}
	}
	log.Println(sum)

	var add1 any
	var add2 any
	divider1 := "[[2]]"
	divider2 := "[[6]]"

	json.Unmarshal([]byte(divider1), &add1)
	json.Unmarshal([]byte(divider2), &add2)

	packets = append(packets, add1, add2)

	sort.Slice(packets, func(i, j int) bool {
		return equal(packets[i], packets[j]) < 0
	})

	part2 := 1
	for i, pack := range packets {
		p,_ := json.Marshal(pack)
		if string(p) == divider1 || string(p) == divider2 {
			part2 *= (i + 1)
		}
	}
	log.Println(part2)
}

func equal(first any, second any) int {
	firstNumber, firstIsNumber := first.(float64)
	secondNumber, secondIsNumber := second.(float64)
	if firstIsNumber && secondIsNumber {
		return int(firstNumber) - int(secondNumber)
	}
	var firstList []any
	var secondList []any

	switch first.(type) {
	case []any, []float64:
		firstList = first.([]any)
	case float64:
		firstList = []any{first}
	}

	switch second.(type) {
	case []any, []float64:
		secondList = second.([]any)
	case float64:
		secondList = []any{second}
	}

	for i := range firstList {
		if len(secondList) <= i {
			return 1
		}
		if r := equal(firstList[i], secondList[i]); r != 0 {
			return r
		}
	}
	if len(secondList) == len(firstList) {
		return 0
	}
	return -1
}
