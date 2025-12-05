// tmpl.go
// Elliott R. Lewandowski
// 2025-12-05
// Template for solving Advent of Code challenges
package main

import (
	"log"
	"os"
)

func main() {

	///////////////////////
	//  Read Input File  //
	///////////////////////

	content, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	solve(string(content))

}

func solve(content string) {

}
