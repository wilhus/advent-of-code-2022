package main

import (
	"fmt"
	"os"
	"strings"
)

func getMarkerAtDistinct(length int) int {
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
		if len(chars) == length {
			break
		}
	}

	return ctr
}

func main() {
	fmt.Println("Part 1: First marker after character", getMarkerAtDistinct(4))
	fmt.Println("Part 2: First marker after character", getMarkerAtDistinct(14))
}
