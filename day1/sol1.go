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

	part1(string(content))

}

// Solves Part 1
func part1(content string) {
	dial := 50
	zeros := 0

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
			dial -= dist
			fmt.Println("L: ", dist)
		case "R":
			dial += dist
			fmt.Println("R: ", dist)
		default:
			log.Fatal("BAD INPUT")
		}

		// Wrap dial value at 0 and 99

		dial %= 100

		if dial > 99 {
			dial -= 100
		} else if dial < 0 {
			dial += 100
		}

		if dial == 0 {
			zeros += 1
		}

		fmt.Println("Dial: ", dial)
	}

	fmt.Println("***********************  RESULT  *************************")
	fmt.Println("VALUE: ", zeros)
}
