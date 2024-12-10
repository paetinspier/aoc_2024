package day9

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	f, err := os.ReadFile("/home/paetin/code/aoc_2024/day9/input.txt")
	if err != nil {
		panic(err)
	}

	values := strings.Split(string(f), "")
	values = values[:len(values)-1]

	//fmt.Println(values)
	fileCount := 0
	arr := []string{}
	isFile := true
	index := 0
	for _, value := range values {
		v, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		if isFile {
			fileCount++
			for j := 0; j < v; j++ {
				arr = append(arr, fmt.Sprint(index))
			}
			index++
			isFile = false
		} else {
			for j := 0; j < v; j++ {
				arr = append(arr, ".")
			}

			isFile = true
		}
	}

	//fmt.Println(arr)

	// part 2
	i := len(arr) - 1
	for i >= 0 {
		if arr[i] == "." {
			i--
		} else {
			v := arr[i]
			end := i

			for i-1 >= 0 && arr[i-1] == v {
				i--
			}
			start := i
			searchForSpace(arr, start, end)

			//fmt.Println(arr)
			i--
		}
	}

	// part 1
	//l := 0
	//r := len(arr) - 1

	//for l < r {
	//	if arr[l] == "." && arr[r] != "." {
	//		arr[l], arr[r] = arr[r], arr[l]
	//		l++
	//		r--
	//	} else {
	//		if arr[l] != "." {
	//			l++
	//		}
	//		if arr[r] == "." {
	//			r--
	//		}
	//	}
	//}

	fmt.Println(arr)

	sum := 0

	for i, value := range arr {
		num, err := strconv.Atoi(value)
		if err != nil {
			continue
		}

		sum += num * i
	}

	fmt.Println("sum:", sum)
}

func searchForSpace(arr []string, start int, end int) {
	fileLength := end - start + 1
	//fmt.Println(arr[start:end+1], fileLength)
	fileCode := arr[start]
	i := 0

	for i < start {
		if arr[i] != "." && arr[i] != fileCode {
			i++
		} else {
			startSpace := i

			for i+1 < end && arr[i+1] == "." {
				i++
			}
			endSpace := i

			spaceLength := endSpace - startSpace + 1

			if spaceLength >= fileLength {
				for j := startSpace; j < startSpace+fileLength; j++ {
					arr[j] = fileCode
				}

				for j := start; j <= end; j++ {
					arr[j] = "."
				}

				return
			}
			i++
		}
	}
}
