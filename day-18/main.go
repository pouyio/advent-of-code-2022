package main

import (
	"bufio"
	"fmt"
	"os"
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
	grid := map[[3]int][][3]int{}
	emptyFaces := 0
	for scanner.Scan() {
		var x, y, z int
		fmt.Sscanf(scanner.Text(), "%d,%d,%d", &x, &y, &z)
		grid[[3]int{x, y, z}] = make([][3]int, 0, 6)
	}

	// calculate adjacent faces for each cube and save them
	for i := range grid {
		faces := totalConnectedCubes(grid, i)
		emptyFaces += 6 - faces
	}

	fmt.Println("Lava droplet surface area:", emptyFaces)
}

func totalConnectedCubes(grid map[[3]int][][3]int, i [3]int) int {
	connectedFaces := 0
	facesToCheck := [][3]int{
		{i[0] - 1, i[1], i[2]},
		{i[0] + 1, i[1], i[2]},
		{i[0], i[1] - 1, i[2]},
		{i[0], i[1] + 1, i[2]},
		{i[0], i[1], i[2] - 1},
		{i[0], i[1], i[2] + 1},
	}

	for _, faceToCheck := range facesToCheck {
		if grid[faceToCheck] != nil {
			connectedFaces++
		}
	}
	return connectedFaces
}

func main() {
	part1()
}
