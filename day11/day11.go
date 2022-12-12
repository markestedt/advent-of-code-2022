package main

import (
	"github.com/markestedt/advent-of-code-2022/utils"
	"log"
	"math"
	"sort"
)

type monkey struct {
	Operation    func(int64) int64
	Test         int64
	ThrowIfTrue  int
	ThrowIfFalse int
	Items        []int64
	Inspected    int
}

func main() {
	monkeys := getMonkeysTest()
	part1 := false
	var rounds int

	var superMod int64 = 1
	for i := 0; i < len(monkeys); i++ {
		superMod *= monkeys[i].Test
	}

	if part1 {
		rounds = 20
	} else {
		rounds = 10000
	}

	for i := 0; i < rounds; i++ {
		for j := 0; j < len(monkeys); j++ {
			m := monkeys[j]

			for len(m.Items) > 0 {
				item := m.Items[0]
				m.Items = m.Items[1:]

				newItem := m.Operation(item)
				if part1 {
					newItem = int64(math.Floor(float64(newItem / 3)))
				} else {
					newItem %= superMod
				}

				m.Inspected++

				var throwTo int
				if newItem%m.Test == 0 {
					throwTo = m.ThrowIfTrue
				} else {
					throwTo = m.ThrowIfFalse
				}

				monkeys[throwTo].Items = append(monkeys[throwTo].Items, newItem)
			}
		}
	}

	var inspected []int
	for j := 0; j < len(monkeys); j++ {
		log.Println(monkeys[j].Inspected)
		inspected = append(inspected, monkeys[j].Inspected)
	}

	sort.Ints(inspected)
	topTwo := utils.Last(inspected, 2)
	log.Println(topTwo)
	log.Println(topTwo[0] * topTwo[1])
}

func getMonkeysTest() map[int]*monkey {
	monkey0 := monkey{
		Items: []int64{79, 98},
		Operation: func(old int64) int64 {
			return old * 19
		},
		Test:         23,
		ThrowIfTrue:  2,
		ThrowIfFalse: 3,
		Inspected:    0,
	}

	monkey1 := monkey{
		Items: []int64{54, 65, 75, 74},
		Operation: func(old int64) int64 {
			return old + 6
		},
		Test:         19,
		ThrowIfTrue:  2,
		ThrowIfFalse: 0,
		Inspected:    0,
	}

	monkey2 := monkey{
		Items: []int64{79, 60, 97},
		Operation: func(old int64) int64 {
			return old * old
		},
		Test:         13,
		ThrowIfTrue:  1,
		ThrowIfFalse: 3,
		Inspected:    0,
	}

	monkey3 := monkey{
		Items: []int64{74},
		Operation: func(old int64) int64 {
			return old + 3
		},
		Test:         17,
		ThrowIfTrue:  0,
		ThrowIfFalse: 1,
		Inspected:    0,
	}

	monkeys := map[int]*monkey{
		0: &monkey0,
		1: &monkey1,
		2: &monkey2,
		3: &monkey3,
	}
	return monkeys
}
