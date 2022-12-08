package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Node struct {
	name      string
	node_type string
	size      int
	children  []*Node
	parent    *Node
}

type Tree struct {
	root *Node
}

func get_node(children []*Node, name string) *Node {
	for _, node := range children {
		if node.name == name {
			return node
		}
	}
	return nil
}

func recursive_print(node *Node, space int, total_size *int) {
	for _, child := range node.children {
		for i := 0; i < space; i++ {
			fmt.Print(" ")
		}
		if child.node_type == "dir" {
			fmt.Printf("- %s (%s)\n", child.name, child.node_type)
			recursive_print(child, space+2, total_size)
		} else {
			fmt.Printf("- %s (%s, size=%d)\n", child.name, child.node_type, child.size)
		}
		node.size += child.size
	}
	if node.size < 100000 {
		*total_size += node.size
	}
}

func find_space(node *Node, needed_space int, least_space *int) int {
	for _, child := range node.children {
		if child.node_type == "dir" {
			find_space(child, needed_space, least_space)
		}
	}
	if node.size > needed_space {
		if *least_space == 0 {
			*least_space = node.size
		} else if node.size < *least_space {
			*least_space = node.size
		}
	}
	return *least_space
}

func print(tree *Tree) {
	root := tree.root
	fmt.Printf("- %s (%s)\n", root.name, root.node_type)
	total_size := 0
	recursive_print(root, 2, &total_size)
	needed_space := 30000000 - (70000000 - root.size)
	least_space := 0
	fmt.Println("Part 1: Total size of dirs < 100000:", total_size)
	fmt.Println("Part 2: Deletable space:", find_space(root, needed_space, &least_space))
}

func main() {
	start := time.Now()
	file, _ := os.ReadFile("input.txt")

	tree := &Tree{}
	current_node := &Node{}

	reg_cd := regexp.MustCompile(`\$ cd .+`)
	reg_dir := regexp.MustCompile(`dir [a-z]+`)
	reg_file := regexp.MustCompile(`[0-9]+ .+`)
	for _, line := range strings.Split(string(file), "\n") {
		if reg_cd.MatchString(line) {
			dir := strings.Split(line, " ")[2]
			if tree.root == nil {
				tree.root = &Node{name: dir, node_type: "dir", size: 0, children: []*Node{}}
				current_node = tree.root
			} else if dir != ".." {
				current_node = get_node(current_node.children, dir)
			} else {
				current_node = current_node.parent
			}
		} else if reg_dir.MatchString(line) {
			current_node.children = append(current_node.children, &Node{name: strings.Split(line, " ")[1], node_type: "dir", size: 0, children: []*Node{}, parent: current_node})
		} else if reg_file.MatchString(line) {
			file_split := strings.Split(line, " ")
			num, _ := strconv.Atoi(file_split[0])
			current_node.children = append(current_node.children, &Node{name: file_split[1], node_type: "file", size: num, children: []*Node{}, parent: current_node})
		}
	}

	print(tree)

	timeElapsed := time.Since(start)
	log.Printf("Solutions took %dms", timeElapsed.Milliseconds())
}
