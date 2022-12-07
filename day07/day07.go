package main

import (
	"github.com/markestedt/advent-of-code-2022/utils"
	"log"
	"strconv"
	"strings"
)

type directory struct {
	Name        string
	Directories []*directory
	Files       []file
	Parent      *directory
	Size        int64
}

type file struct {
	Name string
	Size int64
}

const totalSpace = 70000000
const spaceNeeded = 30000000
const smallDirectoryLimit = 100000

func main() {
	instructions := utils.GetLines("day07/day07.txt")
	root := buildTree(instructions)

	unusedSpace := totalSpace - root.Size
	spaceToFreeUp := spaceNeeded - unusedSpace

	part1, part2 := solve(&root, spaceToFreeUp, root.Size)

	log.Println(part1)
	log.Println(part2)
}

func solve(dir *directory, spaceNeeded int64, spaceToDelete int64) (int64, int64) {
	var part1 int64
	var part2 = spaceToDelete

	if dir.Size <= smallDirectoryLimit {
		part1 += dir.Size
	}

	if dir.Size >= spaceNeeded && dir.Size < part2 {
		part2 = dir.Size
	}

	for _, d := range dir.Directories {
		temp1, temp2 := solve(d, spaceNeeded, part2)

		part1 += temp1
		part2 = temp2
	}

	return part1, part2
}

func buildTree(instructions []string) directory {
	var root directory
	var currentDir = &root

	for _, instruction := range instructions {
		instr := strings.Split(instruction, " ")
		if strings.HasPrefix(instruction, "$") {
			// command
			switch instr[1] {
			case "cd":
				if instr[2] == ".." {
					currentDir = currentDir.Parent
				} else {
					for _, dir := range currentDir.Directories {
						if dir.Name == instr[2] {
							currentDir = dir
						}
					}
				}
			case "ls":
			}
		} else {
			// listing dirs/files
			switch instr[0] {
			case "dir":
				// create dir
				dir := directory{Name: instr[1], Parent: currentDir}
				currentDir.Directories = append(currentDir.Directories, &dir)
			default:
				// create file
				size, _ := strconv.ParseInt(instr[0], 10, 64)
				newFile := file{Name: instr[1], Size: size}

				currentDir.Files = append(currentDir.Files, newFile)
				increaseSize(size, currentDir)
			}
		}
	}
	return root
}

func increaseSize(size int64, dir *directory) {
	dir.Size += size
	if dir.Parent != nil {
		increaseSize(size, dir.Parent)
	}
}
