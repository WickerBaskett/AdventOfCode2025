// sol.go
// Elliott R. Lewandowski
// 2025-12-17
// Solution to Advent of Code 2025 day 8
package main

import (
	uf "day8/unionfind"
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

	uf := uf.UnionFind{}

	uf.MakeSet(1)
	uf.MakeSet(2)
	uf.MakeSet(3)
	uf.MakeSet(4)
	uf.MakeSet(5)

	uf.PrintData()
	fmt.Println()

	uf.Union(1, 2)
	uf.Union(2, 5)
	uf.Union(3, 4)
	uf.Union(3, 1)
	uf.PrintData()
	uf.Find(2)
}

// Union Find ADT to check if two junctions boxes are in the same circuit
// Hash/Dictionary to store distances between any two junction boxes after we compute it once
// Slice of arrays of length 3 to store each point in space
func solve(content string) (int, int) {

	part_1 := 0
	part_2 := 0

	return part_1, part_2
}
