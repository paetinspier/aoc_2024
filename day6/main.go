package day6

import (
	"fmt"
	"os"
	"strings"
)

func Run() {
	matrix_a := getMatrixFromInput()
	sum := mapPath(matrix_a)

	fmt.Println("Positions:", sum)

	matrix_b := getMatrixFromInput()
	loops := 0
	for i, row := range matrix_b {
		for j, cell := range row {
			if cell == "." {
				matrix_c := getMatrixFromInput()
				matrix_c[i][j] = "#"
				loopFound := mapPathLoops(matrix_c)
				if loopFound {
					loops++
				}
			}
		}
	}

	fmt.Println("loops: ", loops)
}

type Direction int

const (
	Up Direction = iota + 1
	Down
	Left
	Right
)

func mapPath(matrix [][]string) int {
	positions := 0

	x, y, err := findGuard(matrix)
	if err != nil {
		panic(err)
	}

	//fmt.Println(x, y)

	matrix[x][y] = "x"
	positions++

	direction := Up

	for !isOB(matrix, x, y) {
		if matrix[x][y] == "." {
			positions++
			matrix[x][y] = "x"
		}

		// move the guard
		switch direction {
		case Up:
			if x-1 >= 0 && matrix[x-1][y] == "#" {
				direction = Right
			} else {
				x -= 1
			}
			break
		case Down:
			if x+1 < len(matrix) && matrix[x+1][y] == "#" {
				direction = Left
			} else {
				x += 1
			}
			break
		case Left:
			if y-1 >= 0 && matrix[x][y-1] == "#" {
				direction = Up
			} else {
				y -= 1
			}
			break
		case Right:
			if y+1 < len(matrix[x]) && matrix[x][y+1] == "#" {
				direction = Down
			} else {
				y += 1
			}
			break
		}
	}
	//for _, line := range matrix {
	//	fmt.Println(line)
	//}
	return positions
}

func mapPathLoops(matrix [][]string) bool {
	positions := 0

	x, y, err := findGuard(matrix)
	if err != nil {
		panic(err)
	}

	//fmt.Println(x, y)

	matrix[x][y] = "x"
	positions++

	direction := Up

	matrixSize := len(matrix) * len(matrix[0])
	loopDetector := 0
	for !isOB(matrix, x, y) {
		if matrix[x][y] == "." {
			positions++
			matrix[x][y] = "x"
		}

		loopDetector++
		if loopDetector > matrixSize {
			return true
		}
		// move the guard
		switch direction {
		case Up:
			if x-1 >= 0 && matrix[x-1][y] == "#" {
				direction = Right
			} else {
				x -= 1
			}
			break
		case Down:
			if x+1 < len(matrix) && matrix[x+1][y] == "#" {
				direction = Left
			} else {
				x += 1
			}
			break
		case Left:
			if y-1 >= 0 && matrix[x][y-1] == "#" {
				direction = Up
			} else {
				y -= 1
			}
			break
		case Right:
			if y+1 < len(matrix[x]) && matrix[x][y+1] == "#" {
				direction = Down
			} else {
				y += 1
			}
			break
		}
	}
	//for _, line := range matrix {
	//	fmt.Println(line)
	//}
	return false
}

func findGuard(matrix [][]string) (int, int, error) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == "^" {
				return i, j, nil
			}
		}
	}
	return -1, -1, fmt.Errorf("could not find guard")
}

func isOB(matrix [][]string, x int, y int) bool {
	return x < 0 || y < 0 || x >= len(matrix) || y >= len(matrix[0])
}

func getMatrixFromInput() [][]string {
	file, err := os.ReadFile("/home/paetin/code/aoc_2024/day6/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")
	matrix := [][]string{}
	for i, line := range lines {
		if i == len(lines)-1 {
			continue
		}
		matrix = append(matrix, strings.Split(line, ""))
	}

	return matrix
}
