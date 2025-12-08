// tmpl.go
// Elliott R. Lewandowski
// 2025-12-05
// Solution to Advent of Code 2025 day 5
package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
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

	lines := strings.Split(content, "\n")

	// This overcounts, I should fix that
	id_ranges := make([]int, 0)
	seen := make(map[int]bool)
	ids := make([]int, 0)

	// Parse Input into two slices
	first_half := true
	for i := 0; i < len(lines); i++ {
		if first_half {
			if lines[i] == "" {
				first_half = false
				continue
			}

			vals := strings.Split(lines[i], "-")
			start, err := strconv.Atoi(vals[0])
			if err != nil {
				log.Fatal(err)
			}

			end, err := strconv.Atoi(vals[1])
			if err != nil {
				log.Fatal(err)
			}
			for j := start; j <= end; j++ {
				if !seen[j] {
					seen[j] = true
					id_ranges = append(id_ranges, j)
				}
			}
		} else {
			if lines[i] == "" {
				continue
			}
			val, err := strconv.Atoi(lines[i])
			if err != nil {
				log.Fatal(err)
			}
			ids = append(ids, val)
		}
	}

	for k := range ids {
		if seen[ids[k]] == true {
			part_1 += 1
		}
	}

	sort.Ints(id_ranges)

	fmt.Println("ID RANGES: ", id_ranges)

	return part_1, part_2
}

func merge_ranges(ranges [][]int) {

}
