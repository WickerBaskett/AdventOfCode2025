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

	solve(string(content))

}

// Solves Part 1
func solve(content string) {
	dial := 50
	part_1 := 0
	part_2 := 0

	for _, line := range strings.Split(string(content), "\n") {
		if line == "" {
			continue
		}

		// Get distance to turn dial
		dist, converr := strconv.Atoi(line[1:])
		if converr != nil {
			log.Fatal(converr)
		}

		// Turn dial in appropriate direction
		switch string(line[0]) {
		case "L":
			new_pos := (dial - dist)
			if dial-(dist%100) < 0 && dial != 0 {
				part_2 += 1
			}
			dial = new_pos
		case "R":
			new_pos := (dial + dist)
			if dial+(dist%100) > 100 && dial != 0 {
				part_2 += 1
			}
			dial = new_pos
		default:
			log.Fatal("BAD INPUT")
		}

		// Counts full rotations past 0
		part_2 += dist / 100

		// Wrap value to range [0, 99]
		dial %= 100

		if dial > 99 {
			dial -= 100
		} else if dial < 0 {
			dial += 100
		}

		// Count exact 0's
		if dial == 0 {
			part_1 += 1
			part_2 += 1
		}
	}

	fmt.Println("***********************  RESULT  *************************")
	fmt.Println("PART 1: ", part_1, "\nPART 2: ", part_2)
}

// Part 2 should be: 5820
