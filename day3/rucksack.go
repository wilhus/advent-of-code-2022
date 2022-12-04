package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func part1() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	priority := make(map[string]int)
	ctr := 1
	for item_type := 'a'; item_type <= 'z'; item_type++ {
		priority[string(item_type)] = ctr
		priority[strings.ToUpper(string(item_type))] = ctr + 26
		ctr++
	}

	sum := 0
	for _, rucksack := range strings.Split(string(file), "\n") {
		middle := len(rucksack) / 2
		compartments := []string{rucksack[:middle], rucksack[middle:]}

		for _, item := range compartments[0] {
			string_item := string(item)
			if strings.Contains(compartments[1], string_item) {
				sum += priority[string_item]
				break
			}
		}
	}

	fmt.Println("Sum is:", sum)
}

func part2() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	priority := make(map[string]int)
	ctr := 1
	for item_type := 'a'; item_type <= 'z'; item_type++ {
		priority[string(item_type)] = ctr
		priority[strings.ToUpper(string(item_type))] = ctr + 26
		ctr++
	}

	groups := [][]string{}
	group := []string{}
	rucksack_ctr := 1
	for _, rucksack := range strings.Split(string(file), "\n") {
		group = append(group, rucksack)

		if rucksack_ctr == 3 {
			groups = append(groups, group)
			group = []string{}
			rucksack_ctr = 1
		} else {
			rucksack_ctr++
		}
	}

	sum := 0
	for _, current_group := range groups {
		for _, item := range current_group[0] {
			string_item := string(item)
			if strings.Contains(current_group[1], string_item) &&
				strings.Contains(current_group[2], string_item) {
				sum += priority[string_item]
				break
			}
		}
	}

	fmt.Println("Sum is:", sum)
}

func main() {
	part1()
	part2()
}
