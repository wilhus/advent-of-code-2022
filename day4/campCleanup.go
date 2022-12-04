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

	ctr := 0
	for _, pair := range strings.Split(string(file), "\n") {
		elves := strings.Split(pair, ",")

		limits := [][]int{}
		for _, elf := range elves {
			string_limits := strings.Split(elf, "-")
			res := make([]int, len(string_limits))
			for idx := range res {
				res[idx], _ = strconv.Atoi(string_limits[idx])
			}
			limits = append(limits, res)
		}

		if limits[0][0] >= limits[1][0] && limits[0][1] <= limits[1][1] ||
			limits[0][0] <= limits[1][0] && limits[0][1] >= limits[1][1] {
			ctr++
		}
	}

	fmt.Println("Number of pairs:", ctr)
}

func part2() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	ctr := 0
	for _, pair := range strings.Split(string(file), "\n") {
		elves := strings.Split(pair, ",")

		limits := [][]int{}
		for _, elf := range elves {
			string_limits := strings.Split(elf, "-")
			res := make([]int, len(string_limits))
			for idx := range res {
				res[idx], _ = strconv.Atoi(string_limits[idx])
			}
			limits = append(limits, res)
		}

		for idx, limit := range limits {
			pls_break := false
			for _, section := range limit {
				if idx == 0 && section >= limits[1][0] && section <= limits[1][1] ||
					idx == 1 && section >= limits[0][0] && section <= limits[0][1] {
					ctr++
					pls_break = true
					break
				}
			}
			if pls_break {
				break
			}
		}
	}

	fmt.Println("Number of pairs:", ctr)
}

func main() {
	part1()
	part2()
}
