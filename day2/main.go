package day2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	file, err := os.ReadFile("/home/paetin/code/aoc_2024/day2/input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		return
	}

	levels := [][]string{}
	lines := strings.Split(string(file), "\n")
	for i := 0; i < len(lines)-1; i++ {
		line := lines[i]
		values := strings.Split(line, " ")
		levels = append(levels, values)
	}

	safeLevels := checkLevels(levels)
	fmt.Printf("Safe levels: %v\n", safeLevels)
}

func checkLevels(levels [][]string) int {
	safeLevels := 0
	for i := 0; i < len(levels); i++ {
		curr := levels[i]
		if isSafeP2(curr) {
			safeLevels++
		}
	}
	return safeLevels
}

func isSafeP2(level []string) bool {
	if isSafe(level){
		return true
	}

	for i := 0; i < len(level); i++ {
		curr := removeNumAt(level, i)
		if isSafe(curr) {
			return true
		}
	}

	return false
}

func isSafe(level []string) bool {
	fmt.Println(level)
	isInc := true
	for i := 1; i < len(level); i++ {
		prev, err := strconv.Atoi(level[i-1])
		if err != nil {
			panic("failed ATOI")
		}
		curr, err := strconv.Atoi(level[i])
		if err != nil {
			panic("failed ATOI")
		}
		diff := prev - curr
		if diff == 0 {
			return false
		}
		if i == 1 && diff < 0 {
			isInc = true
		} else if i == 1 && diff > 0 {
			isInc = false
		} else if diff < 0 && isInc == false {
			return false
		} else if diff > 0 && isInc == true {
			return false
		}

		if diff < 0 {
			diff *= -1
		}

		if diff > 3 {
			return false
		}
	}

	fmt.Println("safe", level)
	return true
}

func removeNumAt(level []string, index int) []string {
	slice := []string{}

	for i := 0; i < len(level); i++ {
		if i != index {
			slice = append(slice, level[i])
		}
	}

	return slice
}
