package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func make_grid() *[][]int {
	file, _ := os.ReadFile("input.txt")

	grid := make([][]int, 99)
	for row, rows := range strings.Split(string(file), "\n") {
		for _, col := range strings.Split(rows, "") {
			tree_height, _ := strconv.Atoi(col)
			grid[row] = append(grid[row], tree_height)
		}
	}

	return &grid
}

func get_cols(grid *[][]int, col int) []int {
	cols := []int{}
	for _, rows := range *grid {
		cols = append(cols, rows[col])
	}
	return cols
}

func visible_direction(tree_height int, dir []int) bool {
	for _, dir_height := range dir {
		if dir_height >= tree_height {
			return false
		}
	}
	return true
}

func is_visible(tree_height int, left []int, right []int, up []int, down []int) bool {
	if visible_direction(tree_height, left) ||
		visible_direction(tree_height, right) ||
		visible_direction(tree_height, up) ||
		visible_direction(tree_height, down) {
		return true
	} else {
		return false
	}
}

func reverse(dir []int) []int {
	new_dir := []int{}
	new_dir = append(new_dir, dir...)
	for i, j := 0, len(new_dir)-1; i < j; i, j = i+1, j-1 {
		new_dir[i], new_dir[j] = new_dir[j], new_dir[i]
	}
	return new_dir
}

func get_scenic_score(tree_height int, left []int, right []int, up []int, down []int) int {
	scenic_score := 1
	dirs := [][]int{reverse(left), right, reverse(up), down}
	for _, dir := range dirs {
		num_visible_trees := 0
		for _, dir_height := range dir {
			if dir_height < tree_height {
				num_visible_trees++
			} else {
				num_visible_trees++
				break
			}
		}
		scenic_score *= num_visible_trees
	}
	return scenic_score
}

func calc_trees(grid *[][]int) {
	sum_trees := 0
	highest_scenic := 0
	for row_idx, rows := range *grid {
		for col_idx, tree_height := range rows {
			cols := get_cols(grid, col_idx)
			left, right, up, down := rows[:col_idx], rows[col_idx+1:], cols[:row_idx], cols[row_idx+1:]
			if is_visible(tree_height, left, right, up, down) {
				sum_trees++
			}
			scenic_score := get_scenic_score(tree_height, left, right, up, down)
			if scenic_score > highest_scenic {
				highest_scenic = scenic_score
			}
		}
	}

	fmt.Println("Part 1 - Amount of visible trees:", sum_trees)
	fmt.Println("Part 2 - Highest scenic score:", highest_scenic)
}

func main() {
	start := time.Now()

	grid := make_grid()
	calc_trees(grid)

	timeElapsed := time.Since(start)
	log.Printf("Solutions took %dms", timeElapsed.Milliseconds())
}
