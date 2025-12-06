// tmpl.go
// Elliott R. Lewandowski
// 2025-12-05
// Template for solving Advent of Code challenges
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

	content, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	sol := solve(string(content))

	fmt.Println("*************************  RESULTS  *************************")
	fmt.Println("Part 2: ", sol)
}

func solve(content string) int {
	part_1 := 0
	part_2 := 0
	for _, line := range strings.Split(content, ",") {

		id_range := strings.Split(line, "-")
		if string(id_range[1][len(id_range[1])-1]) == "\n" {
			id_range[1] = id_range[1][:len(id_range[1])-1]
		}

		beg, err1 := strconv.Atoi(id_range[0])
		if err1 != nil {
			log.Fatal(err1)
		}
		end, err2 := strconv.Atoi(id_range[1])
		if err2 != nil {
			log.Fatal(err2)
		}

		for val := beg; val <= end; val++ {
			id := strconv.Itoa(val)

			//////////////
			//  Part 1  //
			//////////////

			mid := (len(id) / 2)
			valid := false

			for i := 0; i < mid; i++ {
				if len(id)%2 != 0 {
					valid = true
					continue
				}
				if id[i] != id[mid+i] {
					valid = true
					break
				}
			}
			if !valid {
				num, err := strconv.Atoi(id)

				if err != nil {
					log.Fatal(err)
				}

				part_1 += num
			}

			//////////////
			//  Part 2  //
			//////////////

			//fmt.Println("******  ", id, "  ******")

			valid2 := true
			pattern := ""
			checked_digits := ""
			pat_index := 0
			for i := 0; i < len(id); i++ {

				//fmt.Println("Pattern: ", pattern, "\tChecked Digit: ", checked_digits)

				if len(pattern) == 0 {
					pattern += string(id[i])
					continue
				}

				if len(pattern) == pat_index {
					pat_index = 0
				}

				if pattern[pat_index] == id[i] {
					pat_index += 1
					checked_digits += string(id[i])
					if i == len(id)-1 && pat_index == len(pattern) {
						valid2 = false
						break
					}
				} else {
					if len(checked_digits) > 0 {
						i -= len(checked_digits)
						pattern += string(checked_digits[0])
						checked_digits = ""
					} else {
						pattern += string(id[i])
					}
					pat_index = 0
				}

			}

			if !valid2 {

				//fmt.Println("Found ", id)

				num, err := strconv.Atoi(id)

				if err != nil {
					log.Fatal(err)
				}

				part_2 += num
			}
		}
	}
	return part_2
}

// Highest result: 30940224581
