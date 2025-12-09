// tmpl.go
// Elliott R. Lewandowski
// 2025-12-05
// Solution to Advent of Code 2025 day x
package main

import (
	"fmt"
	"log"
	"os"
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

	return part_1, part_2
}
