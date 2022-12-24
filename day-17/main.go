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
	rocks := [5][][2]int{
		{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
		{{1, 0}, {0, 1}, {1, 1}, {2, 1}, {1, 2}},
		{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}},
		{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
		{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
	}
	grid := map[[2]int]bool{}

	var movements []string
	// Use the Scanner to read each line from the file
	for scanner.Scan() {
		movements = strings.Split(scanner.Text(), "")
	}

	movementIndex := 0
	for i := 0; i < 2022; i++ {
		rockIndex := (i + len(rocks)) % len(rocks)
		movementIndex = (movementIndex + len(movements)) % len(movements)
		// place rock in starting point
		rockPosition := calculateStartingPosition(rocks[rockIndex], grid)
		grid, movementIndex = tetris(grid, rockPosition, movements, movementIndex)
	}

	height := 0
	for point := range grid {
		if point[1] > height {
			height = point[1]
		}
	}
	fmt.Println("Tower height:", height+1)
}

func tetris(grid map[[2]int]bool, rockPosition [][2]int, movements []string, moveIndex int) (map[[2]int]bool, int) {
	// move rock horizontally
	moveIndex = (moveIndex + len(movements)) % len(movements)
	rockPosition = moveRockHorizontally(rockPosition, movements[moveIndex], grid)
	moveIndex++
	// is it not settled?
	var moved bool
	moved, rockPosition = moveDown(rockPosition, grid)
	if moved {
		// yes: move horizontally and down again
		return tetris(grid, rockPosition, movements, moveIndex)
	} else {
		// no: add it to the grid and next iteration
		for _, point := range rockPosition {
			grid[point] = true
		}
		return grid, moveIndex
	}
}

func moveDown(rockPosition [][2]int, grid map[[2]int]bool) (bool, [][2]int) {
	targetRockPosition := append([][2]int{}, rockPosition...)
	for i := range targetRockPosition {
		targetRockPosition[i][1] -= 1
		if grid[targetRockPosition[i]] || targetRockPosition[i][1] == -1 {
			return false, rockPosition
		}
	}
	return true, targetRockPosition
}

func moveRockHorizontally(rockPosition [][2]int, s string, grid map[[2]int]bool) [][2]int {
	targetRockPosition := append([][2]int{}, rockPosition...)

	if s == "<" {
		// calculate leftist point in rock
		leftistPoint := 6
		for _, point := range targetRockPosition {
			if point[0] < leftistPoint {
				leftistPoint = point[0]
			}
		}
		// leftist point == 0?
		if leftistPoint == 0 {
			// yes: return original position
			return rockPosition
		}
		// no, move all points
		for i := range targetRockPosition {
			targetRockPosition[i][0] -= 1
			// is any point is already in the grid?
			if grid[targetRockPosition[i]] {
				// yes: return original position
				return rockPosition
			}
		}
		// no: return target position
		return targetRockPosition

	}

	// calculate rightest point in rock
	rightestPoint := 0
	for _, point := range targetRockPosition {
		if point[0] > rightestPoint {
			rightestPoint = point[0]
		}
	}
	// rightest point == 6?
	if rightestPoint == 6 {
		// yes: return original position
		return rockPosition
	}

	// no: move all points
	for i := range targetRockPosition {
		targetRockPosition[i][0] += 1
		// is any point is already in the grid?
		if grid[targetRockPosition[i]] {
			// yes: return original position
			return rockPosition
		}
	}
	// no: return target position
	return targetRockPosition

}

func calculateStartingPosition(rockShape [][2]int, grid map[[2]int]bool) [][2]int {
	newRockPosition := [][2]int{}

	// calculate highest point in grid
	highestPointGrid := 0
	for g := range grid {
		if g[1] > highestPointGrid {
			highestPointGrid = g[1]
		}
	}

	delta := 4 + highestPointGrid
	if len(grid) == 0 {
		delta -= 1
	}
	// add y diference to all points in rock and 2 for x axis
	for _, point := range rockShape {
		newRockPosition = append(newRockPosition, [2]int{point[0] + 2, point[1] + delta})
	}

	return newRockPosition
}

func main() {
	part1()
}
