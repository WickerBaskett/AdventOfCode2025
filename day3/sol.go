// tmpl.go
// Elliott R. Lewandowski
// 2025-12-05
// Solution to Advent of Code 2025 day 3
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

	sol1, sol2 := solve(string(content))

	fmt.Println("*************************  RESULTS  *************************")
	fmt.Println("Part 1: ", sol1)
	fmt.Println("Part 2: ", sol2)

}

func solve(content string) (int, int) {

	part_1 := select_n_batteries(content, 2)
	part_2 := select_n_batteries(content, 12)

	return part_1, part_2
}

func select_n_batteries(content string, n int) int {
	sum := 0
	for _, bank := range strings.Split(content, "\n") {
		if len(bank) < 2 {
			continue
		}

		selected := make([]int, n)

		for i := 0; i < len(bank); i++ {

			// Get current digit of bank
			curr_dig, err := strconv.Atoi(string(bank[i]))
			if err != nil {
				log.Fatal(err)
			}

			// Find earliest digit we can change while still selecting n digits
			index := 0
			if len(bank)-i < n {
				index = n - (len(bank) - i)
			}

			// Find first entry in selected that we can modify and has a value lower than curr_digit
			for j := index; j < len(selected); j++ {
				if selected[j] < curr_dig {
					selected[j] = curr_dig
					// Clear digits after the most recently updated one
					for k := j + 1; k < len(selected); k++ {
						selected[k] = -1
					}
					break
				}
			}
		}

		str := ""
		for digit := range selected {
			str += strconv.Itoa(selected[digit])
		}

		val, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}

		sum += val
	}

	return sum
}
