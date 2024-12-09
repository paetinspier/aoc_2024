package day8

import "os"


func Run() {

}

func parseInput() {
	file, err := os.ReadFile("/home/paetin/code/aoc_2024/day8/input_test.txt")
	if err != nil {
		panic(err)
	}
}
