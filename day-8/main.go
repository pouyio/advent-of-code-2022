package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func part1() {
	// Open the file for reading
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)
	grid := [][]string{}
	// Use the Scanner to read each line from the file
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, strings.Split(line, ""))
	}

	totalVisible := 0

	for iY, line := range grid {
		for iX := range line {
			// trees on the edges
			if iX == 0 || iY == 0 || iX == len(line)-1 || iY == len(grid)-1 {
				totalVisible++
				continue
			}

			// inner trees
			if someSmaller(grid, iY, iX) {
				totalVisible++
			}
		}

	}

	fmt.Println("Trees visible from outside the grid:", totalVisible)
}

func part2() {
	// Open the file for reading
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)
	grid := [][]string{}
	// Use the Scanner to read each line from the file
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, strings.Split(line, ""))
	}

	score := 0

	for iY, line := range grid {
		for iX := range line {
			// trees on the edges
			if iX == 0 || iY == 0 || iX == len(line)-1 || iY == len(grid)-1 {
				continue
			}

			// inner trees
			treeScore := calculateScore(grid, iY, iX)
			if treeScore > score {
				score = treeScore
			}
		}

	}

	fmt.Println("Highest scenic score:", score)
}

func calculateScore(grid [][]string, posY, posX int) int {
	tree := grid[posY][posX]
	top := 0
	bottom := 0
	left := 0
	right := 0
	// top
	for i := posY - 1; i >= 0; i-- {
		if grid[i][posX] < tree {
			top++
		}
		if grid[i][posX] >= tree {
			top++
			break
		}
	}
	// bottom
	for i := posY + 1; i < len(grid); i++ {
		if grid[i][posX] < tree {
			bottom++
		}
		if grid[i][posX] >= tree {
			bottom++
			break
		}
	}
	// right
	for i := posX + 1; i < len(grid[posY]); i++ {
		if grid[posY][i] < tree {
			right++
		}
		if grid[posY][i] >= tree {
			right++
			break
		}
	}
	// left
	for i := posX - 1; i >= 0; i-- {
		if grid[posY][i] < tree {
			left++
		}
		if grid[posY][i] >= tree {
			left++
			break
		}
	}
	return top * bottom * right * left
}

func someSmaller(grid [][]string, posY, posX int) bool {
	tree := grid[posY][posX]
	bigger := 0
	// top
	for i := posY - 1; i >= 0; i-- {
		if grid[i][posX] >= tree {
			bigger++
			break
		}
	}
	// bottom
	for i := posY + 1; i < len(grid); i++ {
		if grid[i][posX] >= tree {
			bigger++
			break
		}
	}
	// right
	for i := posX + 1; i < len(grid[posY]); i++ {
		if grid[posY][i] >= tree {
			bigger++
			break
		}
	}
	// left
	for i := posX - 1; i >= 0; i-- {
		if grid[posY][i] >= tree {
			bigger++
			break
		}
	}
	return bigger < 4
}

func main() {
	part1()
	part2()
}
