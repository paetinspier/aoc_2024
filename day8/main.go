package day8

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Antenna struct {
	X         int
	Y         int
	Frequency string
}

type AntinodeCoordinates struct {
	X         int
	Y         int
	Frequency string
}

func Run() {
	grid := parseInput()

	antennas := findAntennas(grid)

	sort.Slice(antennas, func(i, j int) bool { return antennas[i].Frequency < antennas[j].Frequency })

	currentFrequency := antennas[0].Frequency
	likeAntennas := []Antenna{}
	antinodeCoordinates := []AntinodeCoordinates{}

	for i, antenna := range antennas {
		if i+1 >= len(antennas) {
			likeAntennas = append(likeAntennas, antenna)
			antinodeCoordinates = append(antinodeCoordinates, getAntinodeCoordinates2(grid, likeAntennas)...)
		} else if antennas[i+1].Frequency != currentFrequency {
			likeAntennas = append(likeAntennas, antenna)
			antinodeCoordinates = append(antinodeCoordinates, getAntinodeCoordinates2(grid, likeAntennas)...)
			currentFrequency = antennas[i+1].Frequency
			likeAntennas = []Antenna{}
		} else {
			likeAntennas = append(likeAntennas, antenna)
		}
	}
	//fmt.Println(len(antinodeCoordinates))
	count := 0

	for _, c := range antinodeCoordinates {
		if c.Y < 0 || c.X < 0 || c.Y >= len(grid) || c.X >= len(grid[0]) {
			continue
		}
		if grid[c.Y][c.X] != "#" {
			//fmt.Println(c.Y, c.X)
			grid[c.Y][c.X] = "#"
			count++
		}
	}

	for _, row := range grid {
		fmt.Println(row)
	}

	fmt.Println("antinodes placed", count)
}

func getAntinodeCoordinates(g [][]string, antennas []Antenna) []AntinodeCoordinates {
	//fmt.Println(antennas)
	f := antennas[0].Frequency
	coordinates := []AntinodeCoordinates{}
	for i, a := range antennas {
		for j := i + 1; j < len(antennas); j++ {
			rise := a.Y - antennas[j].Y
			run := a.X - antennas[j].X

			coordinates = append(coordinates, AntinodeCoordinates{X: a.X + run, Y: a.Y + rise, Frequency: f})
			run *= -1
			rise *= -1
			coordinates = append(coordinates, AntinodeCoordinates{X: antennas[j].X + run, Y: antennas[j].Y + rise, Frequency: f})
		}
	}

	return coordinates
}

func getAntinodeCoordinates2(g [][]string, antennas []Antenna) []AntinodeCoordinates {
	f := antennas[0].Frequency
	coordinates := []AntinodeCoordinates{}
	for i, a := range antennas {
		for j := i + 1; j < len(antennas); j++ {
			rise := a.Y - antennas[j].Y
			rise_inverse := a.Y - antennas[j].Y
			rise_inverse *= -1
			run := a.X - antennas[j].X
			run_inverse := a.X - antennas[j].X
			run_inverse *= -1

			coordinates = append(coordinates, AntinodeCoordinates{X: a.X, Y: a.Y, Frequency: f})
			coordinates = append(coordinates, AntinodeCoordinates{X: antennas[j].X, Y: antennas[j].Y, Frequency: f})

			coordinates = append(coordinates, AntinodeCoordinates{X: a.X + run, Y: a.Y + rise, Frequency: f})
			n := 2
			for a.X+run*n >= 0 && a.X+run*n < len(g[0]) && a.Y+rise*n >= 0 && a.Y+rise*n < len(g) {
				coordinates = append(coordinates, AntinodeCoordinates{X: a.X + run*n, Y: a.Y + rise*n, Frequency: f})
				n++
			}
			coordinates = append(coordinates, AntinodeCoordinates{X: antennas[j].X + run_inverse, Y: antennas[j].Y + rise_inverse, Frequency: f})
			n = 2
			for antennas[j].X+run_inverse*n >= 0 && antennas[j].X+run_inverse*n < len(g[0]) && antennas[j].Y+rise_inverse*n >= 0 && antennas[j].Y+rise_inverse*n < len(g) {
				coordinates = append(coordinates, AntinodeCoordinates{X: antennas[j].X + run_inverse*n, Y: antennas[j].Y + rise_inverse*n, Frequency: f})
				n++
			}
		}
	}

	return coordinates
}

func findAntennas(g [][]string) []Antenna {
	antennas := []Antenna{}
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			if g[i][j] != "." {
				antennas = append(antennas, Antenna{Y: i, X: j, Frequency: g[i][j]})
			}
		}
	}

	return antennas
}

func parseInput() [][]string {
	file, err := os.ReadFile("/home/paetin/code/aoc_2024/day8/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")

	grid := [][]string{}
	for i, line := range lines {
		if i == len(lines)-1 {
			continue
		}

		grid = append(grid, strings.Split(line, ""))
	}

	return grid
}
