package day3

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	file, err := os.ReadFile("/home/paetin/code/aoc_2024/day3/input.txt")
	if err != nil {
		panic(err)
	}

	tokens := strings.Split(string(file), "")
	parseTokens(tokens)
}

func parseTokens(tokens []string) {
	do := true
	slow := 0
	fast := 3
	prefix := "mul("
	prefixDont := "don't()"
	prefixDo := "do()"
	sum := 0
	for fast < len(tokens) {
		if slow+7 < len(tokens) && strings.Join(tokens[slow:slow+7], "") == prefixDont {
			do = false
		}
		if slow+4 < len(tokens) && strings.Join(tokens[slow:slow+4], "") == prefixDo {
			do = true
		}
		curr := strings.Join(tokens[slow:fast+1], "")
		if curr == prefix && do {
			nums := getNums(tokens, fast+1)
			//fmt.Println(curr, nums)
			sum += nums[0] * nums[1]
		}
		fast++
		slow++
	}
	fmt.Println("sum: ", sum)
}

func getNums(tokens []string, start int) []int {
	num1 := 0
	num2 := 0
	foundComma := false
	for i := start; i < len(tokens) && i-start < 9; i++ {
		c := tokens[i]
		if c == "," && foundComma == false {
			foundComma = true
		} else if c == "," && foundComma == true {
			return []int{0, 0}
		} else {
			n, err := strconv.Atoi(c)
			if err != nil {
				if c == ")" && foundComma {
					return []int{num1, num2}
				} else {
					return []int{0, 0}
				}
			}
			if foundComma {
				// found number for num2
				num2 *= 10
				num2 += n
			}
			if !foundComma {
				// found number for num1
				num1 *= 10
				num1 += n
			}
		}
	}

	return []int{num1, num2}
}
