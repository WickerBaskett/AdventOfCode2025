// sol.go
// Elliott R. Lewandowski
// 2025-12-05
// Solution to Advent of Code 2025 day 4
package main

import (
	"fmt"
	"log"
	"os"
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

	// Assumes number of columns will be equal for all rows
	rows := strings.Count(content, "\n")
	cols := len(strings.Split(content, "\n")[0])
	matrix := make([][]int, rows)

	// Generate matrix corresponding to input
	for row := range rows {
		matrix[row] = make([]int, cols)
		for col := range cols {
			if string(strings.Split(content, "\n")[row][col]) == "@" {
				matrix[row][col] = 1
			} else {
				matrix[row][col] = 0
			}
		}
	}

	fmt.Println("Matrix:")
	for row := range matrix {
		fmt.Println("\t", matrix[row])
	}

	part_1 += remove_rolls_1(matrix)

	prev_val := -1
	for part_2 != prev_val {

		prev_val = part_2
		part_2 += remove_rolls_2(matrix)
	}

	return part_1, part_2
}

func remove_rolls_1(matrix [][]int) int {
	sum := 0
	rows := len(matrix)
	cols := len(matrix[0])
	for row := range rows {
		for col := range cols {
			if matrix[row][col] != 1 {
				continue
			}

			adj_sum := 0
			if row > 0 {
				if col > 0 {
					adj_sum += matrix[row-1][col-1]
				}

				adj_sum += matrix[row-1][col]

				if col < cols-1 {
					adj_sum += matrix[row-1][col+1]
				}
			}

			if col > 0 {
				adj_sum += matrix[row][col-1]
			}
			if col < cols-1 {
				adj_sum += matrix[row][col+1]
			}

			if row < rows-1 {
				if col > 0 {
					adj_sum += matrix[row+1][col-1]
				}

				adj_sum += matrix[row+1][col]

				if col < cols-1 {
					adj_sum += matrix[row+1][col+1]
				}
			}

			if adj_sum < 4 {
				sum += 1
			}
		}
	}

	return sum
}

func remove_rolls_2(matrix [][]int) int {
	sum := 0
	rows := len(matrix)
	cols := len(matrix[0])
	for row := range rows {
		for col := range cols {
			if matrix[row][col] != 1 {
				continue
			}

			adj_sum := 0
			if row > 0 {
				if col > 0 {
					adj_sum += matrix[row-1][col-1]
				}

				adj_sum += matrix[row-1][col]

				if col < cols-1 {
					adj_sum += matrix[row-1][col+1]
				}
			}

			if col > 0 {
				adj_sum += matrix[row][col-1]
			}
			if col < cols-1 {
				adj_sum += matrix[row][col+1]
			}

			if row < rows-1 {
				if col > 0 {
					adj_sum += matrix[row+1][col-1]
				}

				adj_sum += matrix[row+1][col]

				if col < cols-1 {
					adj_sum += matrix[row+1][col+1]
				}
			}

			if adj_sum < 4 {
				matrix[row][col] = 0
				sum += 1
			}
		}
	}

	return sum
}
