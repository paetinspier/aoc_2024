package day4

import (
	"fmt"
	"os"
	"strings"
)

func Run() {
	file, err := os.ReadFile("/home/paetin/code/aoc_2024/day4/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")
	lines = lines[:len(lines)-1]
	matrix := [][]string{}
	for _, line := range lines {
		matrix = append(matrix, strings.Split(line, ""))
	}

	findXmas(matrix)
}

func findXmas(matrix [][]string) {
	// initialize xmas key
	key := make(map[string]string)
	key["X"] = "M"
	key["M"] = "A"
	key["A"] = "S"

	total := 0
	// iterate over matrix looking for X. once found deploy XMAS Validator
	for y, row := range matrix {
		for x, l := range row {
			if l == "X" {
				//fmt.Println(x, y)
				total += validXMAS(matrix, x, y, key, "X", None)
			}
		}
	}

	fmt.Println("total XMAS", total)

	//for _, line := range matrix {
	//	fmt.Println(line)
	//}

	total = 0
	// iterate over matrix looking for X. once found deploy XMAS Validator
	for y, row := range matrix {
		if y == 0 || y == len(matrix)-1 {
			continue
		}
		for x, l := range row {
			if x == 0 || x == len(row)-1 {
				continue
			}
			if l == "A" {
				//fmt.Println(x, y)
				total += findX_MAS(matrix, x, y)
			}
		}
	}
	fmt.Println("total X-MAS", total)
}

func findX_MAS(matrix [][]string, x int, y int) int {
	northeast := []int{x - 1, y - 1}
	northwest := []int{x + 1, y - 1}
	southwest := []int{x + 1, y + 1}
	southeast := []int{x - 1, y + 1}

	wings := [][]int{northeast, northwest, southwest, southeast}
	mCount := 0
	sCount := 0
	for _, d := range wings {
		if d[1] >= len(matrix) || d[1] < 0 || d[0] >= len(matrix[0]) || d[0] < 0 {
			//fmt.Println("OB", x, y)
			return 0
		} else {
			l := matrix[d[1]][d[0]]
			if l != "M" && l != "S" {
				//fmt.Println("invalid letter", l, x, y)
				return 0
			}
			if l == "M" {
				mCount++
			}
			if l == "S" {
				sCount++
			}
			if sCount > 2 || mCount > 2 {
				//fmt.Println("over count", x, y)
				return 0
			}
		}
	}

	ne := matrix[northeast[1]][northeast[0]]
	nw := matrix[northwest[1]][northwest[0]]
	se := matrix[southeast[1]][southeast[0]]
	sw := matrix[southwest[1]][southwest[0]]

	if ne == sw || nw == se {
		return 0
	}

	//fmt.Println(x, y)
	return 1
}

type Direction int

const (
	None      = 0
	North     = 1
	NorthWest = 2
	West      = 3
	SouthWest = 4
	South     = 5
	SouthEast = 6
	East      = 7
	NorthEast = 8
)

func validXMAS(matrix [][]string, x int, y int, key map[string]string, prev string, d Direction) int {
	// if out of bounds return 0
	if y >= len(matrix) || y < 0 || x >= len(matrix[0]) || x < 0 {
		//fmt.Println("OB", x, y, prev)
		return 0
	}
	// if not XMAS return 0
	if matrix[y][x] != prev {
		//fmt.Println("Invalid", x, y, matrix[y][x])
		return 0
	} else if prev == "S" {
		//fmt.Println("valid", x, y, matrix[y][x])
		return 1
	}

	// continue searching in all directions
	if d == None {
		north := validXMAS(matrix, x, y-1, key, key[prev], North)
		northwest := validXMAS(matrix, x+1, y-1, key, key[prev], NorthWest)
		west := validXMAS(matrix, x+1, y, key, key[prev], West)
		southwest := validXMAS(matrix, x+1, y+1, key, key[prev], SouthWest)
		south := validXMAS(matrix, x, y+1, key, key[prev], South)
		southeast := validXMAS(matrix, x-1, y+1, key, key[prev], SouthEast)
		east := validXMAS(matrix, x-1, y, key, key[prev], East)
		northeast := validXMAS(matrix, x-1, y-1, key, key[prev], NorthEast)

		return north + northwest + west + southwest + south + southeast + east + northeast
	} else {
		switch d {
		case North:
			return validXMAS(matrix, x, y-1, key, key[prev], North)
		case NorthWest:
			return validXMAS(matrix, x+1, y-1, key, key[prev], NorthWest)
		case West:
			return validXMAS(matrix, x+1, y, key, key[prev], West)
		case SouthWest:
			return validXMAS(matrix, x+1, y+1, key, key[prev], SouthWest)
		case South:
			return validXMAS(matrix, x, y+1, key, key[prev], South)
		case SouthEast:
			return validXMAS(matrix, x-1, y+1, key, key[prev], SouthEast)
		case East:
			return validXMAS(matrix, x-1, y, key, key[prev], East)
		case NorthEast:
			return validXMAS(matrix, x-1, y-1, key, key[prev], NorthEast)
		}
	}
	fmt.Println("issue")
	return 0
}
