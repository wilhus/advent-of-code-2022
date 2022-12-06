package main

import (
	"fmt"
	"os"
	"strings"
)

func part1() {
	file, _ := os.ReadFile("input.txt")

	ctr := 0
	chars := ""
	for _, char := range strings.Split(string(file), "") {
		ctr++
		if strings.Contains(chars, char) {
			char_index := strings.Index(chars, char)
			chars = chars[char_index+1:]
		}
		chars += char
		if len(chars) == 4 {
			break
		}
	}

	fmt.Println("Part 1: First marker after character", ctr)
}

func part2() {
	file, _ := os.ReadFile("input.txt")

	ctr := 0
	chars := ""
	for _, char := range strings.Split(string(file), "") {
		ctr++
		if strings.Contains(chars, char) {
			char_index := strings.Index(chars, char)
			chars = chars[char_index+1:]
		}
		chars += char
		if len(chars) == 14 {
			break
		}
	}
	fmt.Println("Part 2: First marker after character", ctr)
}

func main() {
	part1()
	part2()
}
