package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func part1() {
	file, err := os.ReadFile("../data/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	score := 0
	// Iterate over each round
	for _, val := range strings.Split(string(file), "\n") {
		players := strings.Split(val, " ")
		opponent := players[0]
		me := players[1]

		if opponent == "A" && me == "X" || opponent == "B" && me == "Y" || opponent == "C" && me == "Z" {
			score += 3
		} else if opponent == "A" && me == "Y" {
			score += 6
		} else if opponent == "B" && me == "Z" {
			score += 6
		} else if opponent == "C" && me == "X" {
			score += 6
		}

		if me == "X" {
			score += 1
		} else if me == "Y" {
			score += 2
		} else if me == "Z" {
			score += 3
		}
	}

	fmt.Println("My score is:", score)
}

func part2() {
	file, err := os.ReadFile("../data/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	score := 0
	// Iterate over each round
	for _, val := range strings.Split(string(file), "\n") {
		players := strings.Split(val, " ")
		opponent := players[0]
		me := players[1]

		if opponent == "A" {
			if me == "X" {
				score += 3
			} else if me == "Y" {
				score += 4
			} else {
				score += 8
			}
		} else if opponent == "B" {
			if me == "X" {
				score += 1
			} else if me == "Y" {
				score += 5
			} else {
				score += 9
			}
		} else {
			if me == "X" {
				score += 2
			} else if me == "Y" {
				score += 6
			} else {
				score += 7
			}
		}
	}

	fmt.Println("My score is:", score)
}

func main() {
	part1()
	part2()
}
