package day7

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	equations := parseInput()

	sum := 0
	//fmt.Println(equations)
	for _, e := range equations {
		sum += solveEquation(e)
	}

	fmt.Println("sum:", sum)
	sum = 0
	for _, e := range equations {
		sum += solveEquation2(e)
	}

	fmt.Println("sum:", sum)
}

func solveEquation(e []int) int {
	target := e[0]
	permutations := [][]string{}
	getPermutations(len(e)-2, &permutations, []string{})

	for _, p := range permutations {
		j := 1
		total := e[j]
		j++
		for i := 0; i < len(p); i++ {
			if p[i] == "+" {
				total += e[j]
			} else {
				total *= e[j]
			}
			j++
		}
		if total == target {
			return target
		}
	}
	return 0
}

func solveEquation2(e []int) int {
	target := e[0]
	permutations := [][]string{}
	getPermutations2(len(e)-2, &permutations, []string{})

	for _, p := range permutations {
		j := 1
		total := e[j]
		j++
		for i := 0; i < len(p); i++ {
			if p[i] == "+" {
				total += e[j]
			} else if p[i] == "*" {
				total *= e[j]
			} else {
				total = combineNums(total, e[j])
			}
			j++
		}
		if total == target {
			return target
		}
	}
	return 0
}

func combineNums(x, y int) int {
	str1 := strconv.Itoa(x)
	str2 := strconv.Itoa(y)

	totalstr := fmt.Sprintf("%s%s", str1, str2)

	total, err := strconv.Atoi(totalstr)
	if err != nil {
		panic(err)
	}
	return total
}

func getPermutations(r int, permutations *[][]string, current []string) {
	if r == 0 {
		// Append a copy of the current slice to the permutations slice
		*permutations = append(*permutations, append([]string{}, current...))
		return
	}

	current = append(current, "+")
	getPermutations(r-1, permutations, current)

	// Backtrack and change the last element to "*"
	current[len(current)-1] = "*"
	getPermutations(r-1, permutations, current)
}

func getPermutations2(r int, permutations *[][]string, current []string) {
	if r == 0 {
		// Append a copy of the current slice to the permutations slice
		*permutations = append(*permutations, append([]string{}, current...))
		return
	}

	current = append(current, "+")
	getPermutations2(r-1, permutations, current)

	// Backtrack and change the last element to "*"
	current[len(current)-1] = "*"
	getPermutations2(r-1, permutations, current)

	// Backtrack and change the last element to "||"
	current[len(current)-1] = "||"
	getPermutations2(r-1, permutations, current)
}

func parseInput() [][]int {
	file, err := os.ReadFile("/home/paetin/code/aoc_2024/day7/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")
	equations := [][]int{}
	for i, line := range lines {
		if i == len(lines)-1 {
			continue
		}

		characters := strings.Split(line, "")

		equationValues := parseCharacters(characters)
		equations = append(equations, equationValues)
	}
	return equations
}

func parseCharacters(characters []string) []int {
	nums := []int{}

	num := 0
	for i, c := range characters {
		n, err := strconv.Atoi(c)
		if err != nil {
			if num != 0 {
				nums = append(nums, num)
				num = 0
			}
			continue
		} else {
			num *= 10
			num += n
			if i == len(characters)-1 {
				nums = append(nums, num)
			}
		}
	}

	return nums
}
