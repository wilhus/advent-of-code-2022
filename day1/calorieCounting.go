package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var mostCalories = []int{0, 0, 0}
	var currentElfCalories = 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			if currentElfCalories > mostCalories[0] {
				mostCalories[2] = mostCalories[1]
				mostCalories[1] = mostCalories[0]
				mostCalories[0] = currentElfCalories
			} else if currentElfCalories > mostCalories[1] {
				mostCalories[2] = mostCalories[1]
				mostCalories[1] = currentElfCalories
			} else if currentElfCalories > mostCalories[2] {
				mostCalories[2] = currentElfCalories
			}
			currentElfCalories = 0
		} else {
			int, err := strconv.Atoi(text)
			if err != nil {
				log.Fatal(err)
			}
			currentElfCalories += int
		}
	}

	var mostCaloriesTotal = 0
	for _, val := range mostCalories {
		mostCaloriesTotal += val
	}

	fmt.Println(mostCalories)
	fmt.Println(mostCaloriesTotal)
}
