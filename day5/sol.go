// tmpl.go
// Elliott R. Lewandowski
// 2025-12-08
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

	////////////////////////////////
	//  Merge Overlapping Ranges  //
	////////////////////////////////

	// Pre sort so we only have to look at the next element when trying to merge ranges
	sort.Slice(id_ranges, func(i int, j int) bool {
		return id_ranges[i][0] < id_ranges[j][0]
	})

	for j := 0; j < len(id_ranges)-1; j++ {
		if id_ranges[j][1]+1 >= id_ranges[j+1][0] {
			fmt.Println("WE SHOULD MERGE: ", id_ranges[j], " : ", id_ranges[j+1])
			if id_ranges[j][1] < id_ranges[j+1][1] {
				id_ranges[j][1] = id_ranges[j+1][1]
			}
			id_ranges = append(id_ranges[:j+1], id_ranges[j+2:]...)
			j--
		}
	}

	//////////////
	//  Part 1  //
	//////////////

	for k := range ids {
		for l := range len(id_ranges) {
			if ids[k] >= id_ranges[l][0] && ids[k] <= id_ranges[l][1] {
				part_1 += 1
				break
			}
		}
	}

	//////////////
	//  Part 2  //
	//////////////

	for m := range len(id_ranges) {
		part_2 += id_ranges[m][1] - id_ranges[m][0] + 1
	}

	return part_1, part_2
}
