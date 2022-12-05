package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sections := strings.Split(string(file), "\n\n")
	stacks := make(map[int][]string)
	// map of index to stack number
	indeces := make(map[int]int)
	lines := strings.Split(sections[0], "\n")

	// index of stack numbers
	stack_numbers := lines[len(lines)-1:]
	biggest_num := 0
	for _, stack_number := range stack_numbers {
		chars := strings.Split(stack_number, "")
		for idx, char := range chars {
			if char != " " {
				num, _ := strconv.Atoi(char)
				indeces[idx] = num
				if num > biggest_num {
					biggest_num = num
				}
			}
		}
	}

	// add crate to stack using index
	crates := lines[:len(lines)-1]
	for i := len(crates) - 1; i >= 0; i-- {
		crate_line := crates[i]
		chars := strings.Split(crate_line, "")
		for idx, char := range chars {
			if !strings.Contains(" []", char) {
				stacks[indeces[idx]] = append(stacks[indeces[idx]], char)
			}
		}
	}

	rearrangements := [][]int{}
	for _, section := range strings.Split(sections[1], "\n") {
		words := strings.Split(section, " ")
		numbers := []int{}
		for i := 1; i < len(words); i += 2 {
			num, _ := strconv.Atoi(words[i])
			numbers = append(numbers, num)
		}
		rearrangements = append(rearrangements, numbers)
	}

	for _, rearrangement := range rearrangements {
		amount := rearrangement[0]
		from := rearrangement[1]
		to := rearrangement[2]

		for i := 0; i < amount; i++ {
			n := len(stacks[from]) - 1

			// move to new stack
			stacks[to] = append(stacks[to], stacks[from][n])
			// remove from old stack
			stacks[from] = stacks[from][:n]
		}
	}

	message := []string{}
	for i := 1; i <= biggest_num; i++ {
		n := len(stacks[i]) - 1
		message = append(message, stacks[i][n])
	}

	fmt.Println("Message:", strings.Join(message, ""))
}

func part2() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sections := strings.Split(string(file), "\n\n")
	stacks := make(map[int][]string)
	// map of index to stack number
	indeces := make(map[int]int)
	lines := strings.Split(sections[0], "\n")

	// index of stack numbers
	stack_numbers := lines[len(lines)-1:]
	biggest_num := 0
	for _, stack_number := range stack_numbers {
		chars := strings.Split(stack_number, "")
		for idx, char := range chars {
			if char != " " {
				num, _ := strconv.Atoi(char)
				indeces[idx] = num
				if num > biggest_num {
					biggest_num = num
				}
			}
		}
	}

	// add crate to stack using index
	crates := lines[:len(lines)-1]
	for i := len(crates) - 1; i >= 0; i-- {
		crate_line := crates[i]
		chars := strings.Split(crate_line, "")
		for idx, char := range chars {
			if !strings.Contains(" []", char) {
				stacks[indeces[idx]] = append(stacks[indeces[idx]], char)
			}
		}
	}

	rearrangements := [][]int{}
	for _, section := range strings.Split(sections[1], "\n") {
		words := strings.Split(section, " ")
		numbers := []int{}
		for i := 1; i < len(words); i += 2 {
			num, _ := strconv.Atoi(words[i])
			numbers = append(numbers, num)
		}
		rearrangements = append(rearrangements, numbers)
	}

	for _, rearrangement := range rearrangements {
		amount := rearrangement[0]
		from := rearrangement[1]
		to := rearrangement[2]

		for i := amount; i > 0; i-- {
			n := len(stacks[from])
			// move to new stack
			stacks[to] = append(stacks[to], stacks[from][n-i])
			// remove from old stack
			stacks[from] = append(stacks[from][:n-i], stacks[from][n-i+1:]...)
		}
	}

	message := []string{}
	for i := 1; i <= biggest_num; i++ {
		n := len(stacks[i]) - 1
		message = append(message, stacks[i][n])
	}

	fmt.Println("Message:", strings.Join(message, ""))
}

func main() {
	part1()
	part2()
}
