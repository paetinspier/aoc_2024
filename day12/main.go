package day12

import (
	"fmt"
	"os"
	"strings"
)

func Run() {
	f, err := os.ReadFile("/home/paetin/code/aoc_2024/day12/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(f), "\n")
	lines = lines[:len(lines)-1]
	m := [][]string{}
	for _, line := range lines {
		plots := strings.Split(line, "")
		m = append(m, plots)
	}
	sum := 0
	for i, row := range m {
		for j, l := range row {
			if l >= "A" && l <= "Z" {
				area, p := getPlotAreaAndCorners(m, i, j, l)
				fmt.Printf("area %d, perimeter %d\n", area, p)
				sum += area * p
			}
		}
	}

	fmt.Println(sum)
}

func getPlotAreaAndPerimeter(m [][]string, i int, j int, t string) (int, int) {
	if i < 0 || j < 0 || i >= len(m) || j >= len(m[0]) {
		return 0, 0
	}
	area := 0
	perimeter := 0
	if m[i][j] == strings.ToLower(t) {
		return 0, 1
	} else if m[i][j] == t {
		m[i][j] = strings.ToLower(t)
		area++
		perimeter++
	} else {
		return 0, 0
	}

	ua, up := getPlotAreaAndPerimeter(m, i-1, j, t)
	da, dp := getPlotAreaAndPerimeter(m, i+1, j, t)
	la, lp := getPlotAreaAndPerimeter(m, i, j-1, t)
	ra, rp := getPlotAreaAndPerimeter(m, i, j+1, t)

	perimeter += up + dp + lp + rp

	return area + ua + da + la + ra, perimeter
}

func getPlotAreaAndCorners(m [][]string, i int, j int, t string) (int, int) {
	if i < 0 || j < 0 || i >= len(m) || j >= len(m[0]) {
		return 0, 0
	}
	area := 0
	corners := 0
	if m[i][j] == strings.ToLower(t) {
		return 0, corners
	} else if m[i][j] == t {
		m[i][j] = strings.ToLower(t)
		area++
		nc := countCorners2(m, i, j)
		corners += nc
	} else {
		return 0, corners
	}

	ua, uc := getPlotAreaAndCorners(m, i-1, j, t)
	da, dc := getPlotAreaAndCorners(m, i+1, j, t)
	la, lc := getPlotAreaAndCorners(m, i, j-1, t)
	ra, rc := getPlotAreaAndCorners(m, i, j+1, t)

	return area + ua + da + la + ra, corners + uc + dc + lc + rc
}

func countCorners2(m [][]string, i int, j int) int {
	count := 0
	v := strings.ToLower(m[i][j])

	n := false
	e := false
	s := false
	w := false
	// check to see if any of the directions have more of the plot
	if i-1 >= 0 && strings.ToLower(m[i-1][j]) == v {
		n = true
	}
	if i+1 < len(m) && strings.ToLower(m[i+1][j]) == v {
		s = true
	}
	if j-1 >= 0 && strings.ToLower(m[i][j-1]) == v {
		w = true
	}
	if j+1 < len(m[0]) && strings.ToLower(m[i][j+1]) == v {
		e = true
	}

	// check ne outer corner
	if !n && !e {
		//fmt.Println("ne outer corner")
		count += 1
	}
	// check ne inner corner
	if n && e && strings.ToLower(m[i-1][j+1]) != v {
		//fmt.Println("ne inner corner")
		count += 1
	}

	// check se outer corner
	if !s && !e {
		//fmt.Println("se outer corner")
		count += 1
	}
	// check se inner corner
	if s && e && strings.ToLower(m[i+1][j+1]) != v {
		//fmt.Println("se inner corner")
		count += 1
	}

	// check sw outer corner
	if !s && !w {
		//fmt.Println("sw outer corner")
		count += 1
	}
	// check sw inner corner
	if s && w && strings.ToLower(m[i+1][j-1]) != v {
		//fmt.Println("sw inner corner")
		count += 1
	}

	// check nw outer corner
	if !n && !w {
		//fmt.Println("nw outer corner")
		count += 1
	}
	// check nw inner corner
	if n && w && strings.ToLower(m[i-1][j-1]) != v {
		//fmt.Println("nw inner corner")
		count += 1
	}

	fmt.Println(i, j, count, n, e, s, w, m[i][j])
	return count
}
