package day5

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {

	file_a, err := os.ReadFile("/home/paetin/code/aoc_2024/day5/input_a.txt")
	if err != nil {
		panic(err)
	}
	file_b, err := os.ReadFile("/home/paetin/code/aoc_2024/day5/input_b.txt")
	if err != nil {
		panic(err)
	}

	lines_a := strings.Split(string(file_a), "\n")
	lines_b := strings.Split(string(file_b), "\n")

	rules := make(map[int][]int)

	for i, line := range lines_a {
		if i == len(lines_a)-1 {
			continue
		}

		rule := strings.Split(line, "|")
		key, err := strconv.Atoi(rule[0])
		if err != nil {
			panic(err)
		}
		value, err := strconv.Atoi(rule[1])
		if err != nil {
			panic(err)
		}

		rules[key] = append(rules[key], value)
	}

	//fmt.Println(rules)

	pages := [][]int{}

	for i, line := range lines_b {
		if i == len(lines_b)-1 {
			continue
		}
		list := strings.Split(line, ",")
		page := []int{}
		for _, l := range list {
			num, err := strconv.Atoi(l)
			if err != nil {
				panic(err)
			}
			page = append(page, num)
		}

		pages = append(pages, page)

	}

	//fmt.Println(pages)

	validPages, invalidPages := validatePages(pages, rules)

	fmt.Println("valid pages middle sum: ", sumMiddleValues(validPages))

	fixedPages := updateInvalidPages(invalidPages, rules)

	fmt.Println("fix invalid pages sum: ", sumMiddleValues(fixedPages))
}

func updateInvalidPages(pages [][]int, rules map[int][]int) [][]int {
	validPages := [][]int{}

	for _, page := range pages {
		if p, isFixed := fixPage(page, rules); isFixed {
			validPages = append(validPages, p)
		}
	}

	return validPages
}

func fixPage(page []int, rules map[int][]int) ([]int, bool) {
	count := 1

	for count > 0 {
		count = 0

		for i := 1; i < len(page); i++ {
			num := page[i]
			prev := page[i-1]
			if checkCircleDependency(num, prev, rules)	{
				return []int{}, false
			}
			if hasValue(rules[prev], num) {
				page[i], page[i-1] = page[i-1], page[i]
				count++
			}
		}
	}

	return page, true
}

func hasValue(a []int, t int) bool {
	for _, n := range a {
		if n == t {
			return true
		}
	}
	return false
}

func checkCircleDependency(a int, b int, rules map[int][]int) bool {
	return hasValue(rules[a], b) && hasValue(rules[b], a)
}

func validatePages(pages [][]int, rules map[int][]int) ([][]int, [][]int) {
	validPages := [][]int{}
	invalidPages := [][]int{}

	// for each row (page) of numbers we must check to see if any of the numbers break the rules
	for _, page := range pages {
		// innocent until proven guilty
		valid := true
		// for each number in the row check the numbers before it to see if there are any rule breaking numbers
		for i, num := range page {
			// potential rule breakers we are searching for
			ruleBreakers := rules[num]

			for _, ruleBreaker := range ruleBreakers {
				if has, _ := hasTarget(page, ruleBreaker, i+1); has {
					valid = false
					break
				}
			}
		}
		if valid {
			validPages = append(validPages, page)
		} else {
			invalidPages = append(invalidPages, page)
		}
	}
	return validPages, invalidPages
}

func hasTarget(nums []int, target int, start int) (bool, int) {
	for i := 0; i < start; i++ {
		if nums[i] == target {
			return true, i
		}
	}

	return false, -1
}

func sumMiddleValues(pages [][]int) int {
	sum := 0
	for _, page := range pages {
		m := (len(page) - 1) / 2
		sum += page[m]
	}

	return sum
}
