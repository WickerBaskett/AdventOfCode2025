// sol.go
// Elliott R. Lewandowski
// 2025-12-15
// Solution to Advent of Code 2025 day 6
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

	//////////////
	//  Part 1  //
	//////////////

	problems := make([][]string, 0)
	ops := make([]string, 0)

	for _, line := range strings.Split(content, "\n") {
		count := 0
		for _, token := range strings.Split(line, " ") {
			if token == "" {
				continue
			}

			if token == "+" || token == "*" {
				ops = append(ops, token)
				continue
			}

			if len(problems) <= count {
				problems = append(problems, make([]string, 0))
			}

			problems[count] = append(problems[count], token)
			count++
		}
	}

	if len(problems) != len(ops) {
		log.Fatal("Malformed Input: Number of problems does not match number of operations")
	}

	for i, prob := range problems {

		res := 0
		if ops[i] == "*" {
			res = 1
		}

		for value := range prob {

			out, err := strconv.Atoi(prob[value])
			if err != nil {
				log.Fatal(err)
			}

			if ops[i] == "+" {
				res += out
			} else {
				res *= out
			}
		}

		part_1 += res
		fmt.Println("Part 1: ", part_1)
	}

	//////////////
	//  Part 2  //
	//////////////

	problems = make([][]string, 0)
	ops = make([]string, 0)

	for _, line := range strings.Split(content, "\n") {
		count := 0
		for _, token_rune := range line {
			token := string(token_rune)

			if token == "" {
				continue
			}

			if token == "+" || token == "*" {
				ops = append(ops, token)
				continue
			}

			if len(problems) <= count {
				problems = append(problems, make([]string, 0))
			}

			if token == " " {
				count++
				continue
			}

			problems[count] = append(problems[count], token)
			count++
		}
	}

	parsed_problems := make([][]int, 0)
	count := 0
	for _, num := range problems {
		if len(num) == 0 {
			count++
			continue
		}

		if len(parsed_problems) <= count {
			parsed_problems = append(parsed_problems, make([]int, 0))
		}

		output := ""
		for _, char := range num {
			output += char
		}

		value, err := strconv.Atoi(output)
		if err != nil {
			log.Fatal(err)
		}

		parsed_problems[count] = append(parsed_problems[count], value)
	}

	for i, prob := range parsed_problems {

		res := 0
		if ops[i] == "*" {
			res = 1
		}

		for value := range prob {

			if ops[i] == "+" {
				res += prob[value]
			} else {
				res *= prob[value]
			}
		}

		part_2 += res
	}

	return part_1, part_2
}
