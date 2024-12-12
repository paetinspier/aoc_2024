package day11

import (
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func Run() {
	f, err := os.ReadFile("/home/paetin/code/aoc_2024/day11/input.txt")
	if err != nil {
		panic(err)
	}

	strs := strings.Split(string(f), " ")
	nums := []int{}
	for _, str := range strs {
		str = strings.TrimSuffix(str, "\n")
		num, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}

		nums = append(nums, num)
	}
	//fmt.Println(nums)

	blinks := 1000
	var total big.Int

	s := time.Now()
	for _, num := range nums {
		t := total
		tc := totalChildrenAfterN(blinks, num)
		total.Add(&t, &tc)
	}
	e := time.Now()
	elapsed := e.Sub(s)
	fmt.Println(total.String(), "n=", blinks, ", t=", elapsed)
}

var cache sync.Map

func totalChildrenAfterN(n int, num int) big.Int {

	cacheKey := fmt.Sprintf("%d,%d", n, num)

	if cacheValue, ok := cache.Load(cacheKey); ok {
		return cacheValue.(big.Int)
	}

	var result big.Int

	if n == 0 {
		result.SetInt64(1)
	} else if num == 0 {
		t := totalChildrenAfterN(n-1, 1)
		result.Set(&t)
	} else {
		digitLen := getDigitLength(num)
		if digitLen%2 == 0 {
			num1, num2 := cutNumber(num)
			t := totalChildrenAfterN(n-1, num1)
			t2 := totalChildrenAfterN(n-1, num2)
			result.Add(&t, &t2)
		} else {
			t := totalChildrenAfterN(n-1, num*2024)
			result.Set(&t)
		}
	}

	cache.Store(cacheKey, result)
	return result
}

func cutNumber(num int) (int, int) {
	digits := getDigitLength(num)
	mid := int(math.Pow(10, float64(digits/2)))
	return num / mid, num % mid
}

func stitchArray(h1 []int, h2 []int, n1 int, n2 int) []int {
	newArr := []int{}

	for _, n := range h1 {
		newArr = append(newArr, n)
	}

	newArr = append(newArr, n1)
	newArr = append(newArr, n2)

	for _, n := range h2 {
		newArr = append(newArr, n)
	}

	return newArr
}

func getDigitLength(n int) int {
	if n == 0 {
		return 1
	}
	return int(math.Log10(float64(n))) + 1
}
