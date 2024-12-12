package day10

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	f, err := os.ReadFile("/home/paetin/code/aoc_2024/day10/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(f), "\n")
	lines = lines[:len(lines)-1]

	m := [][]int{}

	for _, s := range lines {
		row := []int{}
		for _, rn := range s {
			num, err := strconv.Atoi(string(rn))
			if err != nil {
				panic(err)
			}
			row = append(row, num)

		}
		m = append(m, row)
	}
	sum := 0
	for i, row := range m {
		for j, num := range row {
			if num == 0 {
				rate := getTrailheadScore2(m, j, i, -1)
				fmt.Println(i, j, rate)
				sum += rate
				//reset(m)
			}
		}
	}

	fmt.Println("sum", sum)
}

func getTrailheadScore(m [][]int, x int, y int, prev int) int {
	if x < 0 || y < 0 || x >= len(m[0]) || y >= len(m) {
		return 0
	}
	if m[y][x] == prev+1 || prev == -2 {
		if prev == 8 {
			m[y][x] = -1
			return 1
		}

		return getTrailheadScore(m, x-1, y, m[y][x]) + getTrailheadScore(m, x+1, y, m[y][x]) + getTrailheadScore(m, x, y-1, m[y][x]) + getTrailheadScore(m, x, y+1, m[y][x])
	}

	return 0
}

func getTrailheadScore2(m [][]int, x int, y int, prev int) int {
	if x < 0 || y < 0 || x >= len(m[0]) || y >= len(m) {
		return 0
	}
	if m[y][x] == prev+1 || prev == -1 {
		if prev == 8 {
			return 1
		}

		return getTrailheadScore2(m, x-1, y, m[y][x]) + getTrailheadScore2(m, x+1, y, m[y][x]) + getTrailheadScore2(m, x, y-1, m[y][x]) + getTrailheadScore2(m, x, y+1, m[y][x])
	}

	return 0
}

func reset(m [][]int) {
	for i, row := range m {
		for j := range row {
			if m[i][j] == -1 {
				m[i][j] = 9
			}
		}
	}
}
