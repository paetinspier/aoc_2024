package day1

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Run() {
	text, err := os.ReadFile("/home/paetin/code/aoc_2024/day1/input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	input := string(text)

	lines := strings.Split(input, "\n")

	leftList := []int{}
	rightList := []int{}

	for i, line := range lines[:len(lines)-1] {
		temp := strings.Split(line, "   ")

		leftNum, err := strconv.Atoi(temp[0])
		if err != nil {
			fmt.Printf("Error at line %v left number: %v\n", i, err)
		}

		rightNum, err := strconv.Atoi(temp[1])
		if err != nil {
			fmt.Printf("Error at line %v right number: %v\n", i, err)
		}

		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}

	res := findDiff(leftList, rightList)
	fmt.Println("find diff =", res)
	res2 := findSimScore(leftList, rightList)
	fmt.Println("find sim score =", res2)
}

func findDiff(left []int, right []int) int {
	result := 0
	sort.Slice(left, func(i, j int) bool { return left[i] > left[j] })
	sort.Slice(right, func(i, j int) bool { return right[i] > right[j] })

	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]
		if diff < 0 {
			diff *= -1
		}
		result += diff
	}

	return result
}

func findSimScore(left []int, right []int) int {
	simScore := 0

	sort.Slice(left, func(i, j int) bool { return left[i] > left[j] })

	uniqueLeftNums := []int{}

	lib := make(map[int]int)

	for i := 0; i < len(right); i++ {
		leftNum := left[i]
		rightNum := right[i]

		// create set for left numbers
		if len(uniqueLeftNums) == 0 {
			uniqueLeftNums = append(uniqueLeftNums, leftNum)
		} else if uniqueLeftNums[len(uniqueLeftNums)-1] != leftNum {
			uniqueLeftNums = append(uniqueLeftNums, leftNum)
		}

		// create map of right numbers and their frequencies
		libVal, ok := lib[rightNum]
		if ok {
			lib[rightNum] = libVal + 1
		} else {
			lib[rightNum] = 1
		}
	}

	for _, num := range uniqueLeftNums {
		frequency, ok := lib[num]
		if ok {
			simScore += num * frequency
		}
	}

	return simScore
}

