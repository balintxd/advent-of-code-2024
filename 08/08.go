package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	Char string
	X int
	Y int
}

func main() {
	lines, err := ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	width := len(lines[0])
	height := len(lines)

	nodes := []Node{}
	
	for y, line := range lines {
		for x, char := range line {
			nodes = append(nodes, Node{
				Char: string(char),
				X: x,
				Y: y,
			})
		}
	}

	/* First part */
	antinodes := map[[2]int]bool{}
	for _, node := range nodes {
		if node.Char == "."  {
			continue
		}
		for _, check := range nodes {
			if check == node || check.Char == "." {
				continue
			}
			if check.Char == node.Char {
				anode := Node{
					X: node.X + (node.X - check.X),
					Y: node.Y + (node.Y - check.Y),
				}
				if (anode.X < 0 || anode.X >= width || anode.Y < 0 || anode.Y >= height) {
					continue
				}
				antinodes[[2]int{anode.X, anode.Y}] = true
			}
		}
	}

	fmt.Println("There are",len(antinodes),"on the map for the first part")
	/* End of first part */

	/* Second part */
	antinodes2 := map[[2]int]bool{}
	for _, node := range nodes {
		if node.Char == "." {
			continue
		}
		for _, check := range nodes {
			if check == node || check.Char == "." {
				continue
			}
			if check.Char == node.Char {
				out := false
				iteration := 0
				antinodes2[[2]int{node.X, node.Y}] = true
				for !out {
					iteration += 1
					anode := Node{
						X: node.X + iteration * (node.X - check.X),
						Y: node.Y + iteration * (node.Y - check.Y),
					}
					if (anode.X < 0 || anode.X >= width || anode.Y < 0 || anode.Y >= height) {
						out = true
					} else {
						antinodes2[[2]int{anode.X, anode.Y}] = true
					}
				}
			}
		}
	}

	fmt.Println("There are",len(antinodes2),"on the map for the second part")
	/* End of second part */
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
