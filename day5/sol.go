// tmpl.go
// Elliott R. Lewandowski
// 2025-12-05
// Solution to Advent of Code 2025 day 5
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	///////////////////////
	//  Read Input File  //
	///////////////////////

	content, err := os.ReadFile("demoinput.txt")

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

	id_ranges := make([][]int, 0)
	ids := make([]int, 0)

	///////////////////
	//  Parse Input  //
	///////////////////

	first_half := true
	for i := 0; i < len(lines); i++ {
		if first_half {
			if lines[i] == "" {
				first_half = false
				continue
			}

			curr_range := make([]int, 2)

			vals := strings.Split(lines[i], "-")
			start, err := strconv.Atoi(vals[0])
			if err != nil {
				log.Fatal(err)
			}

			end, err := strconv.Atoi(vals[1])
			if err != nil {
				log.Fatal(err)
			}

			curr_range[0] = start
			curr_range[1] = end

			id_ranges = append(id_ranges, curr_range)
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

	//////////////
	//  Part 1  //
	//////////////

	// Loop over ids
	for k := range ids {
		// Loop over ranges per id
		for l := range len(id_ranges) {
			if ids[k] >= id_ranges[l][0] && ids[k] <= id_ranges[l][1] {
				part_1 += 1
				break
			}
		}
	}

	return part_1, part_2
}
