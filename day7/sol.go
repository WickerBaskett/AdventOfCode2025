// sol.go
// Elliott R. Lewandowski
// 2025-12-05
// Solution to Advent of Code 2025 day 7
package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {

	///////////////////////
	//  Read Input File  //
	///////////////////////

	content, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	sol1, sol2 := solve(string(content))

	fmt.Println("*************************  RESULTS  *************************")
	fmt.Println("Part 1: ", sol1)
	fmt.Println("Part 2: ", sol2)

}

func solve(content string) (int, int) {

	part_1 := 0
	part_2 := 0

	manifold := make([][]string, 0)
	start := [2]int{0, 0}

	for i, line := range strings.Split(content, "\n") {

		if len(line) <= 0 {
			continue
		}

		if len(manifold) <= i {
			manifold = append(manifold, make([]string, 0))
		}

		for j, char := range line {
			if string(char) == "S" {
				start[0] = i
				start[1] = j
			}
			manifold[i] = append(manifold[i], string(char))
		}
	}

	manifold[start[0]][start[1]] = "|"

	manifold_clone := make([][]string, 0)
	for i := range manifold {
		manifold_clone = append(manifold_clone, slices.Clone(manifold[i]))
	}

	part_1 = sol_part_1(manifold_clone)
	part_2 = sol_part_2(manifold, start)

	return part_1, part_2
}

func sol_part_1(manifold [][]string) int {
	part_1 := 0
	for x := range len(manifold) - 1 {
		for y := range manifold[x] {
			if manifold[x][y] == "|" {
				if manifold[x+1][y] != "^" {
					manifold[x+1][y] = "|"
					continue
				}

				part_1 += 1

				if manifold[x+1][y+1] == "." && manifold[x][y+1] != "|" {
					manifold[x+1][y+1] = "|"
				}

				if manifold[x+1][y-1] == "." {
					manifold[x+1][y-1] = "|"
				}

			}
		}
	}
	return part_1
}

func sol_part_2(manifold [][]string, pos [2]int) int {
	part_2 := 0

	for x := pos[0]; x < len(manifold)-1; x++ {

		y := pos[1]

		// BASE CASE
		if x == len(manifold)-2 {
			part_2 += 1
		}

		// RECURSIVE CASE

		if manifold[x+1][y] != "^" {
			manifold[x+1][y] = "|"
			continue
		}

		// Look one tile right
		if manifold[x+1][y+1] == "." {
			manifold_clone := make([][]string, 0)
			for i := range manifold {
				manifold_clone = append(manifold_clone, slices.Clone(manifold[i]))
			}

			part_2 += sol_part_2(manifold_clone, [2]int{x + 1, y + 1})
		}

		// Look on tile left
		if manifold[x+1][y-1] == "." {
			manifold_clone := make([][]string, 0)
			for i := range manifold {
				manifold_clone = append(manifold_clone, slices.Clone(manifold[i]))
			}

			part_2 += sol_part_2(manifold_clone, [2]int{x + 1, y - 1})
		}

		break
	}

	return part_2
}
